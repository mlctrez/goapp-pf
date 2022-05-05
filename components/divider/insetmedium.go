package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InsetMedium struct {
	app.Compo
}

func (c *InsetMedium) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-inset-md").
		Aria("role", "separator")
}
