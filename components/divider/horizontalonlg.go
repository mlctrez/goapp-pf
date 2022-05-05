package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontalOnLg struct {
	app.Compo
}

func (c *HorizontalOnLg) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-horizontal-on-lg pf-m-vertical").
		Aria("role", "separator")
}
