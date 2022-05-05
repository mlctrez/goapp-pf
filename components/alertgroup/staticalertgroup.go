package alertgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StaticAlertGroup struct {
	app.Compo
}

func (c *StaticAlertGroup) Render() app.UI {
	return app.Ul().
		Class("pf-c-alert-group").
		Body(
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-inline pf-m-success").
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
				),
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-inline pf-m-danger").
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
				),
			app.Li().
				Class("pf-c-alert-group__item").
				Body(
					app.Div().
						Class("pf-c-alert pf-m-inline pf-m-info").
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
							app.Div().
								Class("pf-c-alert__description").
								Body(
									app.P().
										Body(
											app.Text("Info alert description."), app.A().
												Href("#").
												Body(
													app.Text("This is a link."),
												),
										),
								),
						),
				),
		)
}
