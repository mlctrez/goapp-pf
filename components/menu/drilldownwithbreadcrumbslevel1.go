package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DrilldownWithBreadcrumbsLevel1 struct {
	app.Compo
}

func (c *DrilldownWithBreadcrumbsLevel1) Render() app.UI {
	return app.Div().
		Class("pf-c-menu pf-m-drilldown").
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
										Aria("expanded", "false").
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
										),
									app.Div().
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
																							app.Text("Deployment"),
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
																										Aria("expanded", "false").
																										Body(
																											app.Span().
																												Class("pf-c-menu__item-main").
																												Body(
																													app.Span().
																														Class("pf-c-menu__item-text").
																														Body(
																															app.Text("Labels"),
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
																																							app.Text("Label 1"),
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
																																							app.Text("Label 2"),
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
																																							app.Text("Label 3"),
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
																																							app.Text("Label 4"),
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
																																							app.Text("Label 5"),
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
																															app.Text("URLs"),
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
																															app.Text("APIs"),
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
																							app.Text("Config"),
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
															app.Text("Start rollout"),
														),
												),
										),
								),
						),
				),
		)
}
