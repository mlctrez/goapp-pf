package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithFlyoutMenuLeft struct {
	app.Compo
}

func (c *WithFlyoutMenuLeft) Render() app.UI {
	return app.Div().
		Class("pf-c-menu pf-m-flyout").
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
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Start rollout"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Pause rollouts"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Aria("expanded", "false").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Add storage"),
														),
													app.Span().
														Class("pf-c-menu__item-toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right"),
														),
												),
										),
									app.Div().
										Class("pf-c-menu").
										Hidden(true).
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
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Application grouping"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Count"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Labels"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Annotations"),
																						),
																				),
																		),
																),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Aria("expanded", true).
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Edit"),
														),
													app.Span().
														Class("pf-c-menu__item-toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right"),
														),
												),
											app.Span().
												Class("pf-c-menu__item-description").
												Body(
													app.Text("Description"),
												),
										),
									app.Div().
										Class("pf-c-menu pf-m-left").
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
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Application grouping"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Count"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Labels"),
																						),
																				),
																		),
																),
															app.Li().
																Class("pf-c-menu__list-item").
																Aria("role", "none").
																Body(
																	app.Button().
																		Class("pf-c-menu__item").
																		Type("button").
																		Aria("role", "menuitem").
																		Body(
																			app.Span().
																				Class("pf-c-menu__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-text").
																						Body(
																							app.Text("Annotations"),
																						),
																				),
																		),
																),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Delete deployment config"),
														),
												),
										),
								),
						),
				),
		)
}
