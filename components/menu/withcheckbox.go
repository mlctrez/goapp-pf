package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithCheckbox struct {
	app.Compo
}

func (c *WithCheckbox) Render() app.UI {
	return app.Div().
		Class("pf-c-menu").
		Body(
			app.Div().
				Class("pf-c-menu__content").
				Body(
					app.Ul().
						Class("pf-c-menu__list").
						Aria("role", "menu").
						Body(
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "menuitem").
								Body(
									app.Label().
										Class("pf-c-menu__item").
										For("with-checkbox-example-check-1").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-check").
														Body(
															app.Span().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("with-checkbox-example-check-1").
																		Name("with-checkbox-example-check-1"),
																),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Checkbox 1"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "menuitem").
								Body(
									app.Label().
										Class("pf-c-menu__item").
										For("with-checkbox-example-check-2").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-check").
														Body(
															app.Span().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("with-checkbox-example-check-2").
																		Name("with-checkbox-example-check-2"),
																),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Checkbox 2"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item pf-m-disabled").
								Aria("role", "menuitem").
								Body(
									app.Label().
										Class("pf-c-menu__item").
										For("with-checkbox-example-check-3").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-check").
														Body(
															app.Span().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("with-checkbox-example-check-3").
																		Name("with-checkbox-example-check-3").
																		Disabled(true),
																),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Checkbox 3"),
														),
												),
										),
								),
						),
				),
		)
}
