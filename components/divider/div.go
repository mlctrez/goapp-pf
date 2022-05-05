package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Div struct {
	app.Compo
}

func (c *Div) Render() app.UI {
	return app.Div().
		Class("pf-c-divider").
		Aria("role", "separator")
}
