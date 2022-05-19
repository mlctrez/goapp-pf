package dropdown

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type DropDown struct {
	ID           string
	Toggle       *Toggle
	Alignments   map[string]string
	AutoFocus    bool
	Children     []app.UI
	ClassName    []string
	ContextProps any
	Direction    string
	Items        []Item
	FullHeight   bool
	Grouped      bool
	Open         bool
	Plain        bool
	Text         bool
	OnSelect     app.EventHandler
	Position     string
}

type Group struct{}
type Item struct {
	Href     string
	Children app.UI
}
type Toggle struct {
	ID       string
	Children []app.UI
	Icon     app.HTMLI // fa.Icon("caret-down")
}

func (t *Toggle) UI() app.HTMLButton {
	return app.Button().Class("pf-c-dropdown__toggle").Body(
		app.Span().Class("pf-c-dropdown__toggle-text").Body(t.Children...),
		app.Span().Class("pf-c-dropdown__toggle-icon").Body(t.Icon),
	)
}

func (d *DropDown) UI() app.UI {
	if d.ID == "" {
		d.ID = uuid.NewString()
	}
	return &dropDown{state: *d}
}

func (d *DropDown) UpdateState(ctx app.Context) {
	ctx.SetState(stateKey(d.ID, "state"), d)
}

func (d *DropDown) onInit() {
	if d.Toggle == nil {
		panic("no toggle provided")
	}
	if d.Toggle.ID == "" {
		d.Toggle.ID = d.ID + "-toggle"
	}
}

func (d *DropDown) menu() app.UI {
	ul := app.Ul().Class("pf-c-dropdown__menu")
	var ulBody []app.UI
	for _, item := range d.Items {
		li := app.Li().Role("menuitem").Body(
			app.A().Class("pf-c-dropdown__menu-item").Href(item.Href).Body(item.Children),
		)
		ulBody = append(ulBody, li)
	}
	ul.Body(ulBody...)
	return ul
}

func (d *DropDown) render() app.UI {

	div := app.Div().Class("pf-c-dropdown")
	var body []app.UI
	body = append(body, d.Toggle.UI().OnClick(func(ctx app.Context, e app.Event) {
		d.Open = !d.Open
		d.UpdateState(ctx)
	}))
	if d.Open {
		div.Class("pf-m-expanded")
		body = append(body, d.menu())
	}
	return div.Body(body...)
}

/*
<ul aria-labelledby="toggle-id" class="pf-c-dropdown__menu" role="menu">
<li role="menuitem">
	<a tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true"
		data-ouia-component-id="OUIA-Generated-DropdownItem-100" aria-disabled="false" class="pf-c-dropdown__menu-item">Link</a></li><li role="menuitem"><button tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true" data-ouia-component-id="OUIA-Generated-DropdownItem-101" aria-disabled="false" type="button" class="pf-c-dropdown__menu-item"> Action </button></li><li role="menuitem"><a tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true" data-ouia-component-id="OUIA-Generated-DropdownItem-102" aria-disabled="true" href="www.google.com" class="pf-m-disabled pf-c-dropdown__menu-item"> Disabled link </a></li><li role="menuitem"><button tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true" data-ouia-component-id="OUIA-Generated-DropdownItem-103" aria-disabled="true" type="button" class="pf-m-aria-disabled pf-c-dropdown__menu-item"> Disabled action </button></li><li role="separator"><div class="pf-c-divider" role="separator"></div></li><li role="menuitem"><a tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true" data-ouia-component-id="OUIA-Generated-DropdownItem-104" aria-disabled="false" class="pf-c-dropdown__menu-item">Separated link</a></li><li role="menuitem"><button tabindex="-1" data-ouia-component-type="PF4/DropdownItem" data-ouia-safe="true" data-ouia-component-id="OUIA-Generated-DropdownItem-105" aria-disabled="false" type="button" class="pf-c-dropdown__menu-item"> Separated action </button></li></ul>
*/

/*



Everything below here is boilerplate and should not need modification



*/

var _ app.Mounter = (*dropDown)(nil)
var _ app.Initializer = (*dropDown)(nil)

type dropDown struct {
	app.Compo
	state DropDown
}

func (d *dropDown) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(d.state.ID, "state")).Value(&d.state)
}

func (d *dropDown) OnInit() {
	d.state.onInit()
}

func (d *dropDown) Render() app.UI {
	return d.state.render()
}

func stateKey(id, name string) string {
	prefix := "github.com/mlctrez/goapp-pf/components/dropdown/DropDown"
	return fmt.Sprintf("%s/%s/%s", prefix, id, name)
}
