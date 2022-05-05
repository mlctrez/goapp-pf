package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineTypes struct {
	app.Compo
}

func (c *InlineTypes) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-alert pf-m-inline").
				Aria("label", "Inline default alert").
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
									app.Text("Default inline alert:"),
								),
							app.Text("Default inline alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-info pf-m-inline").
				Aria("label", "Inline information alert").
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
							app.Text("Info inline alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-success pf-m-inline").
				Aria("label", "Inline success alert").
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
							app.Text("Success inline alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-warning pf-m-inline").
				Aria("label", "Inline warning alert").
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
							app.Text("Warning inline alert title"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-danger pf-m-inline").
				Aria("label", "Inline danger alert").
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
							app.Text("Danger inline alert title"),
						),
				),
		)
}
