package panel

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NoBody struct {
	app.Compo
}

func (c *NoBody) Render() app.UI {
	return app.Div().
		Class("pf-c-panel").
		Body(
			app.Div().
				Class("pf-c-panel__main").
				Body(
					app.Text("Main content"),
				),
		)
}
