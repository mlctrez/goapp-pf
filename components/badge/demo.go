package badge

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

func Demo() app.UI {
	return app.Div().Body(
		(&Badge{ID: "badgeOne", Read: true, Children: ui.S(app.Text("7"))}).UI(),
		(&Badge{ID: "badgeTwo", Read: false, Children: ui.S(app.Text("999+"))}).UI(),
	)

}
