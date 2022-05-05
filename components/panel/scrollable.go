package panel

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Scrollable struct {
	app.Compo
}

func (c *Scrollable) Render() app.UI {
	return app.Div().
		Class("pf-c-panel pf-m-scrollable").
		Body(
			app.Div().
				Class("pf-c-panel__main").
				Body(
					app.Div().
						Class("pf-c-panel__main-body").
						Body(
							app.Text("Main content"), app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
							app.Br(),
							app.Br(),
							app.Text("Main content"),
						),
				),
		)
}
