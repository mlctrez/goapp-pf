package alert

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type Alert struct {
	ID          string
	Title       string
	ActionClose *ActionCloseButton
	ActionLinks []*ActionLink
	AriaLabel   string
	Children    []app.UI
	ClassName   []string
	CustomIcon  app.UI
	Expandable  bool
	Inline      bool
	LiveRegion  bool
	Plain       bool
	// TODO: timeouts and tool tips
	//Timeout           time.Duration
	TitleHeadingLevel string
	ToggleAriaLabel   string
	//ToolTipPosition   string
	TruncateTitle int
	Variant       Variant
	VariantLabel  string
}

type ActionCloseButton struct {
	AriaLabel    string
	ClassName    []string
	OnClose      app.EventHandler
	VariantLabel string
}

type ActionLink struct {
	Children  string
	ClassName []string
	OnClick   app.EventHandler
}

func (a *Alert) UI() app.UI {
	return &alert{Alert: *a}
}

type alert struct {
	app.Compo
	Alert
	expanded bool
}

func (a *alert) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(a.ID, "expanded")).Value(&a.expanded)
}

func (a *alert) Render() app.UI {

	a.Variant = a.Variant.checkVariant()

	class := []string{"pf-c-alert", fmt.Sprintf("pf-m-%s", a.Variant)}
	if a.Inline {
		class = append(class, "pf-m-inline")
		if a.Plain {
			class = append(class, "pf-m-plain")
		}
	}
	if a.Expandable {
		class = append(class, "pf-m-expandable")
		if a.expanded {
			class = append(class, "pf-m-expanded")
		}
	}

	class = append(class, a.ClassName...)

	div := app.Div().ID(a.ID).Class(class...)
	if a.LiveRegion {
		div.Aria("live", "polite")
		div.Aria("atomic", "false")
	}
	if a.AriaLabel != "" {
		div.Aria("label", a.AriaLabel)
	} else {
		div.Aria("label", fmt.Sprintf("%s Alert", a.Variant.toTitle()))
	}
	var body []app.UI
	if a.Expandable {
		body = append(body, a.toggle())
	}

	body = append(body, a.icon(), a.title())
	if a.ActionClose != nil {
		body = append(body, a.action())
	}
	if !a.Expandable || (a.Expandable && a.expanded) {
		body = append(body, a.description())
	}
	if a.ActionLinks != nil {
		body = append(body, a.actionGroup())
	}

	return div.Body(body...)

}

func (a *alert) toggle() app.HTMLDiv {
	div := app.Div().Class("pf-c-alert__toggle").Aria("label", a.ToggleAriaLabel)
	button := app.Button()
	button.Class("pf-c-button pf-m-plain").Type("button").Aria("expanded", a.expanded)
	button.OnClick(func(ctx app.Context, e app.Event) {
		ctx.SetState(stateKey(a.ID, "expanded"), !a.expanded)
	})
	button.Body(
		app.Span().Class("pf-c-alert__toggle-icon").Body(fa.Icon("angle-right")),
	)

	return div.Body(button)
}

func (a *alert) icon() app.HTMLDiv {
	div := app.Div().Class("pf-c-alert__icon")
	if a.CustomIcon != nil {
		return div.Body(a.CustomIcon)
	}
	return div.Body(a.Variant.icon().Class("fa-fw"))
}

func (a *alert) title() app.UI {
	if a.VariantLabel == "" {
		a.VariantLabel = a.Variant.toTitle() + " alert:"
	}

	screenReader := app.Span().Class("pf-screen-reader").Text(a.VariantLabel)

	return a.titleElement(screenReader, app.Text(a.Title))
}

func (a *alert) titleElement(elem ...app.UI) app.UI {
	class := []string{"pf-c-alert__title"}
	styles := map[string]string{}
	if a.TruncateTitle > 0 {
		class = append(class, "pf-m-truncate")
		styles["--pf-c-alert__title--max-lines"] = fmt.Sprintf("%d", a.TruncateTitle)
	}

	switch a.TitleHeadingLevel {
	case "h1":
		return app.H1().Class(class...).Styles(styles).Body(elem...)
	case "h2":
		return app.H2().Class(class...).Styles(styles).Body(elem...)
	case "h3":
		return app.H3().Class(class...).Styles(styles).Body(elem...)
	case "h4":
		return app.H4().Class(class...).Styles(styles).Body(elem...)
	case "h5":
		return app.H5().Class(class...).Styles(styles).Body(elem...)
	case "h6":
		return app.H6().Class(class...).Styles(styles).Body(elem...)
	default:
		return app.H4().Class(class...).Styles(styles).Body(elem...)
	}
}

func (a *alert) action() app.HTMLDiv {
	div := app.Div().Class("pf-c-alert__action")

	ariaLabel := a.ActionClose.AriaLabel
	if ariaLabel == "" {
		variantLabel := a.ActionClose.VariantLabel
		if variantLabel == "" {
			variantLabel = string(a.Variant)
		}
		ariaLabel = fmt.Sprintf("Close %s alert: %s", variantLabel, a.Title)
	}

	button := app.Button().Class("pf-c-button pf-m-plain").Class(a.ActionClose.ClassName...)
	button.Type("button").Aria("label", ariaLabel).Body(fa.Icon("times")).OnClick(a.ActionClose.OnClose)

	return div.Body(button)
}

func (a *alert) description() app.HTMLDiv {
	div := app.Div().Class("pf-c-alert__description")
	div.Body(a.Children...)
	return div
}

func (a *alert) actionGroup() app.HTMLDiv {
	div := app.Div().Class("pf-c-alert__action-group")
	var body []app.UI

	for _, link := range a.ActionLinks {
		button := app.Button().Class("pf-c-button pf-m-link pf-m-inline")
		button.Class(link.ClassName...)
		button.Type("button").Text(link.Children).OnClick(link.OnClick)
		body = append(body, button)
	}

	return div.Body(body...)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Alert{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
