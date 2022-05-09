package alertgroup

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/alert"
	"github.com/mlctrez/goapp-pf/internal/key"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

type AlertGroup struct {
	ID            string
	Children      []*alert.Alert
	ClassName     []string
	LiveRegion    bool
	Toast         bool
	OverflowClick app.EventHandler
}

func (a *AlertGroup) UI() app.UI {
	return &alertGroup{AlertGroup: *a}
}

func (a *AlertGroup) Add(ctx app.Context, alert *alert.Alert) {
	if alert.ID == "" {
		alert.ID = uuid.NewString()
	}
	ctx.NewActionWithValue(stateKey(a.ID, "add"), alert)
}

type alertGroup struct {
	app.Compo
	AlertGroup
	children []*alert.Alert
}

func (a *alertGroup) OnMount(ctx app.Context) {
	ctx.Handle(stateKey(a.ID, "add"), a.addAlert)
	ctx.Handle(stateKey(a.ID, "remove"), a.removeAlert)
}

func (a *alertGroup) Render() app.UI {
	source := a.Children
	if a.Toast {
		source = a.children
	}
	ul := app.Ul().ID(a.ID)
	if len(source) == 0 {
		ul.Style("display", "hidden")
		return ul
	}

	ul.Class("pf-c-alert-group")
	if a.Toast {
		ul.Class("pf-m-toast")
	}
	ul.Class(a.ClassName...)
	if a.LiveRegion {
		ul.Aria("live", "polite")
		ul.Aria("atomic", "false")
	}

	var items []app.UI

	overflowItems := 0
	doSwitch := len(source)%2 == 0
	for i, child := range source {
		if i > 3 {
			overflowItems++
			continue
		}
		alert := child.UI()
		if doSwitch {
			alert = ui.Wrap(alert)
		}
		li := app.Li().Class("pf-c-alert-group__item").Body(alert)
		items = append(items, li)
	}
	if overflowItems > 0 {
		button := app.Button().Class("pf-c-alert-group__overflow-button")
		button.Text(fmt.Sprintf("View %d more alerts", overflowItems))
		button.OnClick(a.OverflowClick)
		items = append(items, button)
	}

	return ul.Body(items...)
}

func (a *alertGroup) addAlert(_ app.Context, action app.Action) {

	anAlert := action.Value.(*alert.Alert)
	anAlert.ActionClose = &alert.ActionCloseButton{
		OnClose: func(ctx app.Context, e app.Event) {
			ctx.NewAction(stateKey(a.ID, "remove"), app.T("id", anAlert.ID))
		},
	}
	a.children = append(a.children, anAlert)
}

func (a *alertGroup) removeAlert(_ app.Context, action app.Action) {
	var newChildren []*alert.Alert

	id := action.Tags.Get("id")
	for _, child := range a.children {
		if id == child.ID {
			child.ActionClose = nil
			continue
		}
		newChildren = append(newChildren, child)
	}
	a.children = newChildren
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&AlertGroup{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
