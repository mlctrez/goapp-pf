package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Types struct {
	app.Compo
}

func (c *Types) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-alert").
				Aria("label", "Default alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-bell").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Default alert:"),
								),
							app.Text("Default alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-info").
				Aria("label", "Information alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-info-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Info alert:"),
								),
							app.Text("Info alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-success").
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
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-warning").
				Aria("label", "Warning alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-exclamation-triangle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Warning alert:"),
								),
							app.Text("Warning alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-danger").
				Aria("label", "Danger alert").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-exclamation-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Danger alert:"),
								),
							app.Text("Danger alert title"),
						),
				),
		)
}
