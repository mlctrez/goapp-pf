package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticalOnLg struct {
	app.Compo
}

func (c *VerticalOnLg) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-horizontal pf-m-vertical-on-lg").
		Aria("role", "separator")
}
