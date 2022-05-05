package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DrilldownWithBreadcrumbsLevel2 struct {
	app.Compo
}

func (c *DrilldownWithBreadcrumbsLevel2) Render() app.UI {
	return app.Div().
		Class("pf-c-menu pf-m-drilldown pf-m-drilled-in").
		Style("--pf-c-menu__content--Height", " 96px;").
		Body(
			app.Div().
				Class("pf-c-menu__breadcrumb").
				Body(
					app.Nav().
						Class("pf-c-breadcrumb").
						Aria("label", "breadcrumb").
						Body(
							app.Ol().
								Class("pf-c-breadcrumb__list").
								Body(
									app.Li().
										Class("pf-c-breadcrumb__item").
										Body(
											app.Span().
												Class("pf-c-breadcrumb__item-divider").
												Body(
													app.I().
														Class("fas fa-angle-right").
														Aria("hidden", true),
												),
											app.Button().
												Class("pf-c-breadcrumb__link").
												Type("button").
												Body(
													app.Text("All"),
												),
										),
									app.Li().
										Class("pf-c-breadcrumb__item").
										Body(
											app.Span().
												Class("pf-c-breadcrumb__item-divider").
												Body(
													app.I().
														Class("fas fa-angle-right").
														Aria("hidden", true),
												),
											app.H1().
												Class("pf-c-breadcrumb__heading").
												Body(
													app.Button().
														Class("pf-c-breadcrumb__link pf-m-current").
														Type("button").
														Aria("current", "page").
														Body(
															app.Text("Edit"),
														),
												),
										),
								),
						),
				),
			app.Hr().
				Class("pf-c-divider"),
			app.Div().
				Class("pf-c-menu__content").
				Body(
					app.Ul().
						Class("pf-c-menu__list").
						Aria("role", "menu").
						Body(
							app.Li().
								Class("pf-c-menu__list-item pf-m-current-path").
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
