package panel

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Footer struct {
	app.Compo
}

func (c *Footer) Render() app.UI {
	return app.Div().
		Class("pf-c-panel").
		Body(
			app.Div().
				Class("pf-c-panel__main").
				Body(
					app.Div().
						Class("pf-c-panel__main-body").
						Body(
							app.Text("Main content"),
						),
				),
			app.Div().
				Class("pf-c-panel__footer").
				Body(
					app.Text("Footer content"),
				),
		)
}
