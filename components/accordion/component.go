package accordion

import (
	"fmt"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type Accordion struct {
	ID               string
	AriaLabel        string
	AsDefinitionList bool
	DisplaySize      string
	HeadingLevel     string
	Bordered         bool
	MultipleExpand   bool
	Fixed            bool
	Items            []*Item
}

func (a *Accordion) UI() app.UI {
	return &accordion{Accordion: *a}
}

type Item struct {
	Toggle  Toggle
	Content Content
}

type Content struct {
	ID       string
	Children []app.UI
}

type Toggle struct {
	ID       string
	Expanded bool
	Children []app.UI
}

var _ app.Mounter = (*accordion)(nil)

type accordion struct {
	app.Compo
	Accordion
	sync.Once
}

func (a *accordion) OnMount(ctx app.Context) {
	ctx.Handle(stateKey(a.ID, "toggleAction"), func(context app.Context, action app.Action) {
		toggleId := action.Tags.Get("id")

		for _, item := range a.Items {
			if item.Toggle.ID == toggleId {
				item.Toggle.Expanded = !item.Toggle.Expanded
			} else {
				if !a.MultipleExpand {
					item.Toggle.Expanded = false
				}
			}

		}

		a.Update()
	})
}

func (a *accordion) Render() app.UI {

	var body []app.UI
	for i, item := range a.Items {

		if item.Toggle.ID == "" {
			item.Toggle.ID = fmt.Sprintf("%s-toggle-%d", a.ID, i+1)
		}

		expanded := item.Toggle.Expanded

		toggleButton := a.toggleButton(item, expanded)

		var contentBodies []app.UI
		for _, child := range item.Content.Children {
			cb := app.Div().Class("pf-c-accordion__expanded-content-body").Body(child)
			contentBodies = append(contentBodies, cb)
		}

		body = append(body,
			a.wrapToggleButton(toggleButton),
			a.wrapContentBodies(i+1, item, expanded, contentBodies...),
		)
	}

	classes := []string{"pf-c-accordion"}
	if a.Bordered {
		classes = append(classes, "pf-m-bordered")
	}
	if a.DisplaySize == "large" {
		classes = append(classes, "pf-m-display-lg")
	}

	if a.AsDefinitionList {
		return app.Dl().ID(a.ID).Class(classes...).
			Aria("label", a.AriaLabel).Body(body...)
	} else {
		return app.Div().ID(a.ID).Class(classes...).
			Aria("label", a.AriaLabel).Body(body...)
	}

}

func (a *accordion) wrapContentBodies(index int, item *Item, expanded bool, bodies ...app.UI) app.UI {
	classes := []string{"pf-c-accordion__expanded-content"}
	if expanded {
		classes = append(classes, "pf-m-expanded")
	}
	if a.Fixed {
		classes = append(classes, "pf-m-fixed")
	}
	id := item.Content.ID
	if id == "" {
		id = fmt.Sprintf("%s-expanded-content-%d", a.ID, index)
	}
	if a.AsDefinitionList {
		return app.Dd().ID(id).Class(classes...).Hidden(!expanded).Body(bodies...)
	}
	return app.Div().ID(id).Class(classes...).Hidden(!expanded).Body(bodies...)
}

func (a *accordion) wrapToggleButton(button app.HTMLButton) app.UI {
	if a.AsDefinitionList {
		return app.Dt().Body(button)
	}
	// 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6'
	switch a.HeadingLevel {
	case "h1":
		return app.H1().Body(button)
	case "h2":
		return app.H2().Body(button)
	case "h3":
		return app.H3().Body(button)
	case "h4":
		return app.H4().Body(button)
	case "h5":
		return app.H5().Body(button)
	case "h6":
		return app.H6().Body(button)
	default:
		return app.H3().Body(button)
	}
}

func (a *accordion) toggleButton(item *Item, expanded bool) app.HTMLButton {
	classes := []string{"pf-c-accordion__toggle"}
	if expanded {
		classes = append(classes, "pf-m-expanded")
	}
	toggleId := item.Toggle.ID

	click := func(ctx app.Context, e app.Event) {
		ctx.NewAction(stateKey(a.ID, "toggleAction"), app.T("id", toggleId))
	}

	b := app.Button().Class(classes...)
	b.Type("button").ID(toggleId).Aria("expanded", expanded).
		OnClick(click).Body(toggleText(item), toggleIcon())

	return b
}

func toggleText(item *Item) app.HTMLSpan {
	return app.Span().Class("pf-c-accordion__toggle-text").Body(item.Toggle.Children...)
}

func toggleIcon() app.HTMLSpan {
	return app.Span().Class("pf-c-accordion__toggle-icon").Body(
		app.I().Class("fas fa-angle-right").Aria("hidden", true),
	)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Accordion{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
