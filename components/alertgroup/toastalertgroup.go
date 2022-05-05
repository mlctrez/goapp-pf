package alertgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ToastAlertGroup struct {
	app.Compo
}

func (c *ToastAlertGroup) Render() app.UI {
	return app.Ul().
		Class("pf-c-alert-group pf-m-toast").
		Body(
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-success").
						Aria("label", "Success toast alert").
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
								ID("alert_one_title").
								Body(
									app.Span().
										Class("pf-screen-reader").
										Body(
											app.Text("Success alert:"),
										),
									app.Text("Success toast alert title"),
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
				),
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-danger").
						Aria("label", "Danger toast alert").
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
								ID("alert_two_title").
								Body(
									app.Span().
										Class("pf-screen-reader").
										Body(
											app.Text("Danger alert:"),
										),
									app.Text("Danger toast alert title"),
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
				),
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-info").
						Aria("label", "Information toast alert").
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
								ID("alert_three_title").
								Body(
									app.Span().
										Class("pf-screen-reader").
										Body(
											app.Text("Info alert:"),
										),
									app.Text("Info toast alert title"),
								),
							app.Div().
								Class("pf-c-alert__description").
								Body(
									app.P().
										Body(
											app.Text("Info toast alert description. From the settings tab, click"), app.A().
												Href("#").
												Body(
													app.Text("View logs"),
												),
											app.Text("to review the details."),
										),
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
				),
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Button().
						Class("pf-c-alert-group__overflow-button").
						Body(
							app.Text("View 3 more notifications"),
						),
				),
		)
}
