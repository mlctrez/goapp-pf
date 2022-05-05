package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NavWithDrilldownMenuLevelTwo struct {
	app.Compo
}

func (c *NavWithDrilldownMenuLevelTwo) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav").
		Aria("label", "Drilldown menu example").
		Body(
			app.Div().
				Class("pf-c-menu pf-m-drilldown pf-m-drilled-in").
				Body(
					app.Div().
						Class("pf-c-menu__content").
						Style("--pf-c-menu__content--Height", " 228px;").
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
																	app.Text("Pause rollout"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item pf-m-current").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Current link"),
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
																	app.Text("Add storage"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item pf-m-current-path").
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
																		Class("pf-c-menu__list-item pf-m-drill-up").
																		Aria("role", "none").
																		Body(
																			app.Button().
																				Class("pf-c-menu__item").
																				Type("button").
																				Aria("role", "menuitem").
																				TabIndex(0).
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-main").
																						Body(
																							app.Span().
																								Class("pf-c-menu__item-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-left"),
																								),
																							app.Span().
																								Class("pf-c-menu__item-text").
																								Body(
																									app.Text("Edit"),
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
																										Class("pf-c-menu__list-item pf-m-drill-up").
																										Aria("role", "none").
																										Body(
																											app.Button().
																												Class("pf-c-menu__item").
																												Type("button").
																												Aria("role", "menuitem").
																												TabIndex(0).
																												Body(
																													app.Span().
																														Class("pf-c-menu__item-main").
																														Body(
																															app.Span().
																																Class("pf-c-menu__item-toggle-icon").
																																Body(
																																	app.I().
																																		Class("fas fa-angle-left"),
																																),
																															app.Span().
																																Class("pf-c-menu__item-text").
																																Body(
																																	app.Text("Deployment"),
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
																																	app.Text("Routes"),
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
																																	app.Text("Nodes"),
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
																												Aria("expanded", "false").
																												Body(
																													app.Span().
																														Class("pf-c-menu__item-main").
																														Body(
																															app.Span().
																																Class("pf-c-menu__item-text").
																																Body(
																																	app.Text("Advanced settings"),
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
																																		Class("pf-c-menu__list-item pf-m-drill-up").
																																		Aria("role", "none").
																																		Body(
																																			app.Button().
																																				Class("pf-c-menu__item").
																																				Type("button").
																																				Aria("role", "menuitem").
																																				TabIndex(0).
																																				Body(
																																					app.Span().
																																						Class("pf-c-menu__item-main").
																																						Body(
																																							app.Span().
																																								Class("pf-c-menu__item-toggle-icon").
																																								Body(
																																									app.I().
																																										Class("fas fa-angle-left"),
																																								),
																																							app.Span().
																																								Class("pf-c-menu__item-text").
																																								Body(
																																									app.Text("Advanced settings"),
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
																																									app.Text("Reports"),
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
																																									app.Text("Policies"),
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
																																									app.Text("Systems"),
																																								),
																																						),
																																				),
																																		),
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
																				Aria("expanded", "false").
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-main").
																						Body(
																							app.Span().
																								Class("pf-c-menu__item-text").
																								Body(
																									app.Text("RBAC"),
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
																										Class("pf-c-menu__list-item pf-m-drill-up").
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
																																Class("pf-c-menu__item-toggle-icon").
																																Body(
																																	app.I().
																																		Class("fas fa-angle-left"),
																																),
																															app.Span().
																																Class("pf-c-menu__item-text").
																																Body(
																																	app.Text("RBAC"),
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
																																	app.Text("Reports"),
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
																																	app.Text("Policies"),
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
																																	app.Text("Systems"),
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
																									app.Text("Thing with a longer label"),
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
												Aria("expanded", "false").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Configuration"),
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
																		Class("pf-c-menu__list-item pf-m-drill-up").
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
																								Class("pf-c-menu__item-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-left"),
																								),
																							app.Span().
																								Class("pf-c-menu__item-text").
																								Body(
																									app.Text("Configuration"),
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
																									app.Text("Profile"),
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
																										Class("pf-c-menu__list-item pf-m-drill-up").
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
																																Class("pf-c-menu__item-toggle-icon").
																																Body(
																																	app.I().
																																		Class("fas fa-angle-left"),
																																),
																															app.Span().
																																Class("pf-c-menu__item-text").
																																Body(
																																	app.Text("Profile"),
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
																																	app.Text("Avatar"),
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
																																	app.Text("Name"),
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
																																	app.Text("Title"),
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
																									app.Text("Time zone"),
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
																									app.Text("Account settings"),
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
																									app.Text("Thing with a longer label"),
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
																									app.Text("Edit access settings"),
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
																										Class("pf-c-menu__list-item pf-m-drill-up").
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
																																Class("pf-c-menu__item-toggle-icon").
																																Body(
																																	app.I().
																																		Class("fas fa-angle-left"),
																																),
																															app.Span().
																																Class("pf-c-menu__item-text").
																																Body(
																																	app.Text("Edit access settings"),
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
																																	app.Text("Global access"),
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
																																	app.Text("Account access"),
																																),
																														),
																												),
																										),
																								),
																						),
																				),
																		),
																),
														),
												),
										),
								),
						),
				),
		)
}
