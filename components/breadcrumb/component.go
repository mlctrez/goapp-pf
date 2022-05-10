package breadcrumb

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/key"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

type BreadCrumb struct {
	ID        string
	AriaLabel string
	Children  []ui.UI
	ClassName []string
}

type Item struct {
	Children  []app.UI
	ClassName []string
	//Component string
	Active bool
	//DropDown  bool
	Target string
	To     string
}

func (i *Item) UI() app.UI {
	var liBody []app.UI
	liBody = append(liBody, app.Span().Class("pf-c-breadcrumb__item-divider").Body(fa.Icon("angle-right")))
	if i.To != "" {
		anchor := app.A().Href(i.To).Class("pf-c-breadcrumb__link")
		if i.Active {
			anchor.Class("pf-m-current").Aria("current", "page")
		}
		if i.Target != "" {
			anchor.Target(i.Target)
		}
		liBody = append(liBody, anchor.Body(i.Children...))
	} else {
		liBody = append(liBody, i.Children...)
	}

	return app.Li().Class("pf-c-breadcrumb__item").Class(i.ClassName...).Body(liBody...)
}

type Heading struct {
	Children  []app.UI
	ClassName []string
	//Component string
	Active bool
	//DropDown  bool
	Target string
	To     string
}

func (h *Heading) UI() app.UI {
	var liBody []app.UI
	liBody = append(liBody, app.Span().Class("pf-c-breadcrumb__item-divider").Body(fa.Icon("angle-right")))

	heading := app.H1().Class("pf-c-breadcrumb__heading")
	liBody = append(liBody, heading)

	var headingBody []app.UI
	if h.To != "" {
		anchor := app.A().Href(h.To).Class("pf-c-breadcrumb__link")
		if h.Active {
			anchor.Class("pf-m-current").Aria("current", "page")
		}
		if h.Target != "" {
			anchor.Target(h.Target)
		}
		headingBody = append(headingBody, anchor.Body(h.Children...))
	} else {
		headingBody = append(headingBody, h.Children...)
	}

	heading.Body(headingBody...)

	return app.Li().Class("pf-c-breadcrumb__item").Class(h.ClassName...).Body(liBody...)
}

func (b *BreadCrumb) UI() app.UI {
	return &breadCrumb{BreadCrumb: *b}
}

type breadCrumb struct {
	app.Compo
	BreadCrumb
}

func (b *breadCrumb) OnMount(ctx app.Context) {
}

func (b *breadCrumb) Render() app.UI {
	nav := app.Nav().Class("pf-c-breadcrumb").Aria("label", b.AriaLabel)
	nav.Class(b.ClassName...)
	var items []app.UI
	for _, child := range b.Children {
		items = append(items, child.UI())
	}
	return nav.Body(app.Ol().Class("pf-c-breadcrumb__list").Body(items...))
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&BreadCrumb{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
