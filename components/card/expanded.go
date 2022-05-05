package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expanded struct {
	app.Compo
}

func (c *Expanded) Render() app.UI {
	return app.Div().
		Class("pf-c-card pf-m-expanded").
		ID("card-expanded-example").
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
								ID("card-expanded-example-toggle").
								Aria("labelledby", "card-expanded-example-title card-expanded-example-toggle").
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
										ID("card-expanded-example-dropdown-kebab-button").
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
										Aria("labelledby", "card-expanded-example-dropdown-kebab-button").
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
										ID("card-expanded-example-check").
										Name("card-expanded-example-check").
										Aria("label", "card-expanded-example checkbox"),
								),
						),
					app.Div().
						Class("pf-c-card__title").
						ID("card-expanded-example-title").
						Body(
							app.Text("Title"),
						),
				),
			app.Div().
				Class("pf-c-card__expandable-content").
				Body(
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
				),
		)
}
