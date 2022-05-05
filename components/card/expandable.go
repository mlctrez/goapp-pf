package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expandable struct {
	app.Compo
}

func (c *Expandable) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-expandable-example").
		Body(
			app.Div().
				Class("pf-c-card__header").
				Body(
					app.Div().
						Class("pf-c-card__header-toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Details").
								ID("card-expandable-example-toggle").
								Aria("labelledby", "card-expandable-example-title card-expandable-example-toggle").
								Body(
									app.Span().
										Class("pf-c-card__header-toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-card__actions").
						Body(
							app.Div().
								Class("pf-c-dropdown").
								Body(
									app.Button().
										Class("pf-c-dropdown__toggle pf-m-plain").
										ID("card-expandable-example-dropdown-kebab-button").
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
										Aria("labelledby", "card-expandable-example-dropdown-kebab-button").
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
										ID("card-expandable-example-check").
										Name("card-expandable-example-check").
										Aria("label", "card-expandable-example checkbox"),
								),
						),
					app.Div().
						Class("pf-c-card__title").
						ID("card-expandable-example-title").
						Body(
							app.Text("Title"),
						),
				),
		)
}
