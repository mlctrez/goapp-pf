package panel

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Raised struct {
	app.Compo
}

func (c *Raised) Render() app.UI {
	return app.Div().
		Class("pf-c-panel pf-m-raised").
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
		)
}
