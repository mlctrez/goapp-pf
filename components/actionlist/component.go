package actionlist

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

type ActionList struct {
	ID        string
	IconList  bool
	ClassName []string
	Children  []*Group
}

type Group struct {
	Children  []*Item
	ClassName []string
}

func (g *Group) UI() app.UI {
	div := app.Div().Class("pf-c-action-list__group").Class(g.ClassName...)
	var children []ui.UI
	for _, child := range g.Children {
		children = append(children, child)
	}
	return div.Body(ui.ToAppUI(children...)...)
}

type Item struct {
	Child     app.UI
	ClassName []string
}

func (i *Item) UI() app.UI {
	div := app.Div().Class("pf-c-action-list__item").Class(i.ClassName...)
	return div.Body(i.Child)
}

func (a *ActionList) UI() app.UI {
	return &actionList{ActionList: *a}
}

type actionList struct {
	app.Compo
	ActionList
}

func (a *actionList) Render() app.UI {
	div := app.Div().Class("pf-c-action-list")
	if a.IconList {
		div.Class("pf-m-icons")
	}
	var children []app.UI
	if len(a.Children) == 1 {
		for _, child := range a.Children[0].Children {
			children = append(children, child.UI())
		}
	} else {
		for _, child := range a.Children {
			children = append(children, child.UI())
		}
	}
	return div.Body(children...)
}
