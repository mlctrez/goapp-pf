package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithOnlyActionsInHeadNoTitlefooter struct {
	app.Compo
}

func (c *WithOnlyActionsInHeadNoTitlefooter) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-action-example-3").
		Body(
			app.Div().
				Class("pf-c-card__header").
				Body(
					app.Div().
						Class("pf-c-card__actions").
						Body(
							app.Div().
								Class("pf-c-dropdown").
								Body(
									app.Button().
										Class("pf-c-dropdown__toggle pf-m-plain").
										ID("card-action-example-3-dropdown-kebab-button").
										Aria("expanded", "false").
										Type("button").
										Aria("label", "Actions").
										Body(
											app.I().
												Class("fas fa-ellipsis-v").
												Aria("hidden", true),
										),
									app.Ul().
										Class("pf-c-dropdown__menu pf-m-align-right").
										Aria("labelledby", "card-action-example-3-dropdown-kebab-button").
										Hidden(true).
										Body(
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Link"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-dropdown__menu-item").
														Type("button").
														Body(
															app.Text("Action"),
														),
												),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item pf-m-disabled").
														Href("#").
														Aria("disabled", true).
														TabIndex(-1).
														Body(
															app.Text("Disabled link"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-dropdown__menu-item").
														Type("button").
														Disabled(true).
														Body(
															app.Text("Disabled action"),
														),
												),
											app.Li().
												Class("pf-c-divider").
												Aria("role", "separator"),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Separated link"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-check pf-m-standalone").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("card-action-example-3-check").
										Name("card-action-example-3-check").
										Aria("label", "card-action-example-3 checkbox"),
								),
						),
				),
			app.Div().
				Class("pf-c-card__body").
				ID("card-action-example-3-check-label").
				Body(
					app.Text("This is the card body. There are only actions in the card head."),
				),
		)
}
