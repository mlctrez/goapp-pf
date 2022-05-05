package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineVariations struct {
	app.Compo
}

func (c *InlineVariations) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-alert pf-m-success pf-m-inline").
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
					app.Div().
						Class("pf-c-alert__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Close success alert: Success alert title").
								Body(
									app.I().
										Class("fas fa-times").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-alert__description").
						Body(
							app.P().
								Body(
									app.Text("Success alert description. This should tell the user more information about the alert."),
								),
						),
					app.Div().
						Class("pf-c-alert__action-group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("View details"),
								),
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("Ignore"),
								),
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
					app.Div().
						Class("pf-c-alert__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Close success alert: Success alert title").
								Body(
									app.I().
										Class("fas fa-times").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-alert__description").
						Body(
							app.P().
								Body(
									app.Text("Success alert description. This should tell the user more information about the alert."), app.A().
										Href("#").
										Body(
											app.Text("This is a link."),
										),
								),
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
					app.Div().
						Class("pf-c-alert__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Close success alert: Success alert title").
								Body(
									app.I().
										Class("fas fa-times").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-alert__action-group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("View details"),
								),
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("Ignore"),
								),
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
		)
}
