package dropdown

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

var demoDropDown = &DropDown{
	Toggle: &Toggle{
		Children: ui.S(app.Text("dropdown 1")),
		Icon:     fa.Icon("caret-down"),
	},
	Items: []Item{
		{Href: "#", Children: app.Text("item one")},
		{Href: "#", Children: app.Text("item two")},
		{Href: "#", Children: app.Text("item three")},
	},
}

func Demo() app.UI {
	return demoDropDown.UI()
}
