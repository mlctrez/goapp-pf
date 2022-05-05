package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithActions struct {
	app.Compo
}

func (c *WithActions) Render() app.UI {
	return app.Div().
		Class("pf-c-menu").
		Body(
			app.Div().
				Class("pf-c-menu__content").
				Body(
					app.Section().
						Class("pf-c-menu__group").
						Body(
							app.H1().
								Class("pf-c-menu__group-title").
								Body(
									app.Text("Actions"),
								),
							app.Ul().
								Class("pf-c-menu__list").
								Aria("role", "menu").
								Body(
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item pf-m-selected").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 1"),
																),
															app.Span().
																Class("pf-c-menu__item-select-icon").
																Body(
																	app.I().
																		Class("fas fa-check").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action").
												Type("button").
												Aria("label", "Actions").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-ellipsis-v").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item pf-m-select").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 2"),
																),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action").
												Type("button").
												Aria("label", "Alert").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-bell").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item pf-m-disabled").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item").
												Type("button").
												Disabled(true).
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 3 disabled"),
																),
															app.Span().
																Class("pf-c-menu__item-select-icon").
																Body(
																	app.I().
																		Class("fas fa-check").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action").
												Type("button").
												Disabled(true).
												Aria("label", "Copy").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-clipboard").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item pf-m-selected").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 4"),
																),
															app.Span().
																Class("pf-c-menu__item-select-icon").
																Body(
																	app.I().
																		Class("fas fa-check").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action").
												Type("button").
												Aria("label", "Expand").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-bars").
																Aria("hidden", true),
														),
												),
										),
								),
						),
				),
		)
}
