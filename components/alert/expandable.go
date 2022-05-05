package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expandable struct {
	app.Compo
}

func (c *Expandable) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-alert pf-m-expandable pf-m-success").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("expanded", "false").
								ID("alert-expandable-example-1-toggle").
								Aria("label", "Details").
								Aria("labelledby", "alert-expandable-example-1-title alert-expandable-example-1-toggle").
								Body(
									app.Span().
										Class("pf-c-alert__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						ID("alert-expandable-example-1-title").
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
						Hidden(true).
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
				Class("pf-c-alert pf-m-expandable pf-m-expanded pf-m-success").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("expanded", true).
								ID("alert-expandable-example-2-toggle").
								Aria("label", "Details").
								Aria("labelledby", "alert-expandable-example-2-title alert-expandable-example-2-toggle").
								Body(
									app.Span().
										Class("pf-c-alert__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						ID("alert-expandable-example-2-title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Success alert:"),
								),
							app.Text("Success alert title (expanded)"),
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
				Class("pf-c-alert pf-m-expandable pf-m-success pf-m-inline").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("expanded", "false").
								ID("alert-expandable-example-3-toggle").
								Aria("label", "Details").
								Aria("labelledby", "alert-expandable-example-3-title alert-expandable-example-3-toggle").
								Body(
									app.Span().
										Class("pf-c-alert__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						ID("alert-expandable-example-3-title").
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
						Hidden(true).
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
				Class("pf-c-alert pf-m-expandable pf-m-expanded pf-m-success pf-m-inline").
				Aria("label", "Success alert").
				Body(
					app.Div().
						Class("pf-c-alert__toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("expanded", true).
								ID("alert-expandable-example-4-toggle").
								Aria("label", "Details").
								Aria("labelledby", "alert-expandable-example-4-title alert-expandable-example-4-toggle").
								Body(
									app.Span().
										Class("pf-c-alert__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title").
						ID("alert-expandable-example-4-title").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Success alert:"),
								),
							app.Text("Success alert title (expanded)"),
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
		)
}
