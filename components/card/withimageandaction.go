package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithImageAndAction struct {
	app.Compo
}

func (c *WithImageAndAction) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-action-example-1").
		Body(
			app.Div().
				Class("pf-c-card__header").
				Body(
					app.Div().
						Class("pf-c-card__header-main").
						Body(
							app.Img().
								Src("/assets/images/pf_logo.svg").
								Width(300).
								Alt("Logo"),
						),
					app.Div().
						Class("pf-c-card__actions").
						Body(
							app.Div().
								Class("pf-c-dropdown").
								Body(
									app.Button().
										Class("pf-c-dropdown__toggle pf-m-plain").
										ID("card-action-example-1-dropdown-kebab-button").
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
										Aria("labelledby", "card-action-example-1-dropdown-kebab-button").
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
										ID("card-action-example-1-check").
										Name("card-action-example-1-check").
										Aria("label", "card-action-example-1 checkbox"),
								),
						),
				),
			app.Div().
				Class("pf-c-card__title").
				ID("card-action-example-1-check-label").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
				),
			app.Div().
				Class("pf-c-card__footer").
				Body(
					app.Text("Footer"),
				),
		)
}
