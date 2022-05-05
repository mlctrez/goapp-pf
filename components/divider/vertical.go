package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Vertical struct {
	app.Compo
}

func (c *Vertical) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-vertical pf-m-inset-md").
		Aria("role", "separator")
}
