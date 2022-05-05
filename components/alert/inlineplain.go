package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlinePlain struct {
	app.Compo
}

func (c *InlinePlain) Render() app.UI {
	return app.Div().
		Class("pf-c-alert pf-m-success pf-m-plain pf-m-inline").
		Aria("label", "Success alert").
		Body(
			app.Div().
				Class("pf-c-alert__icon").
				Body(
					app.I().
						Class("fas fa-fw fa-check-circle").
						Aria("hidden", true),
				),
			app.P().
				Class("pf-c-alert__title").
				Body(
					app.Span().
						Class("pf-screen-reader").
						Body(
							app.Text("Success alert:"),
						),
					app.Text("Success alert title"),
				),
		)
}
