package breadcrumb

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

func Demo() app.UI {

	b := &BreadCrumb{
		Children: []ui.UI{
			&Item{To: "#first", Children: ui.S(app.Text("First"))},
			&Item{To: "#second", Children: ui.S(app.Text("Second"))},
			&Heading{Children: ui.S(app.Text("Heading"))},
		},
		AriaLabel: "A breadcrumb",
	}

	return b.UI()
}
