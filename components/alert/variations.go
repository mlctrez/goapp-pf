package alert

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Variations struct {
	app.Compo
}

func (c *Variations) Render() app.UI {
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
				Class("pf-c-alert pf-m-success").
				Aria("label", "Success alert with title truncation").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title pf-m-truncate").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Success alert:"),
								),
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur pellentesque neque cursus enim fringilla tincidunt. Proin lobortis aliquam dictum. Nam vel ullamcorper nulla, nec blandit dolor. Vivamus pellentesque neque justo, nec accumsan nulla rhoncus id. Suspendisse mollis, tortor quis faucibus volutpat, sem leo fringilla turpis, ac lacinia augue metus in nulla. Cras vestibulum lacinia orci. Pellentesque sodales consequat interdum. Sed porttitor tincidunt metus nec iaculis. Pellentesque non commodo justo. Morbi feugiat rhoncus neque, vitae facilisis diam aliquam nec. Sed dapibus vitae quam at tristique. Nunc vel commodo mi. Mauris et rhoncus leo."),
						),
					app.Div().
						Class("pf-c-alert__description").
						Body(
							app.P().
								Body(
									app.Text("This example uses \".pf-m-truncate\" to limit the title to a single line and truncate any overflow text with ellipses."),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-alert pf-m-success").
				Aria("label", "Success alert with title truncation at 2 lines").
				Body(
					app.Div().
						Class("pf-c-alert__icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
					app.P().
						Class("pf-c-alert__title pf-m-truncate").
						Style("--pf-c-alert__title--max-lines", " 2").
						Body(
							app.Span().
								Class("pf-screen-reader").
								Body(
									app.Text("Success alert:"),
								),
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur pellentesque neque cursus enim fringilla tincidunt. Proin lobortis aliquam dictum. Nam vel ullamcorper nulla, nec blandit dolor. Vivamus pellentesque neque justo, nec accumsan nulla rhoncus id. Suspendisse mollis, tortor quis faucibus volutpat, sem leo fringilla turpis, ac lacinia augue metus in nulla. Cras vestibulum lacinia orci. Pellentesque sodales consequat interdum. Sed porttitor tincidunt metus nec iaculis. Pellentesque non commodo justo. Morbi feugiat rhoncus neque, vitae facilisis diam aliquam nec. Sed dapibus vitae quam at tristique. Nunc vel commodo mi. Mauris et rhoncus leo."),
						),
					app.Div().
						Class("pf-c-alert__description").
						Body(
							app.P().
								Body(
									app.Text("This example uses \".pf-m-truncate\" and sets \"--pf-c-alert__title--max-lines: 2\" to limit title to two lines and truncate any overflow text with ellipses."),
								),
						),
				),
		)
}
