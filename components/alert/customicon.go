package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CustomIcon struct {
	app.Compo
}

func (c *CustomIcon) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-alert pf-m-success").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-cog").
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
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-success pf-m-inline").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-cog").
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
				),
		)
}
