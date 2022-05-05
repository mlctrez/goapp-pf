package treeview

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithCheckboxes struct {
	app.Compo
}

func (c *WithCheckboxes) Render() app.UI {
	return app.Div().
		Class("pf-c-tree-view").
		Body(
			app.Ul().
				Class("pf-c-tree-view__list").
				Aria("role", "tree").
				Body(
					app.Li().
						Class("pf-c-tree-view__list-item pf-m-expanded").
						Aria("role", "treeitem").
						Aria("expanded", true).
						TabIndex(0).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Button().
														Class("pf-c-tree-view__node-toggle").
														Aria("label", "Toggle").
														Aria("expanded", true).
														Body(
															app.Span().
																Class("pf-c-tree-view__node-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-tree-view__node-check").
														Body(
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("check-1").
																		Aria("label", "Tree view node check"),
																),
														),
													app.Label().
														Class("pf-c-tree-view__node-text").
														For("check-1").
														Body(
															app.Text("Application launcher"),
														),
												),
										),
								),
							app.Ul().
								Class("pf-c-tree-view__list").
								Aria("role", "group").
								Body(
									app.Li().
										Class("pf-c-tree-view__list-item pf-m-expanded").
										Aria("role", "treeitem").
										Aria("expanded", true).
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Button().
																		Class("pf-c-tree-view__node-toggle").
																		Aria("label", "Toggle").
																		Aria("expanded", true).
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-tree-view__node-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-2").
																						Checked(true).
																						Aria("label", "Tree view node check checked"),
																				),
																		),
																	app.Label().
																		Class("pf-c-tree-view__node-text").
																		For("check-2").
																		Body(
																			app.Text("Application 1"),
																		),
																),
														),
												),
											app.Ul().
												Class("pf-c-tree-view__list").
												Aria("role", "group").
												Body(
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														TabIndex(-1).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-3").
																										Checked(true).
																										Aria("label", "Tree view node check checked"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-3").
																						Body(
																							app.Text("Settings"),
																						),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														TabIndex(-1).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-4").
																										Checked(true).
																										Aria("label", "Tree view node check checked"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-4").
																						Body(
																							app.Text("Loader"),
																						),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														Aria("expanded", "false").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Button().
																						Class("pf-c-tree-view__node-toggle").
																						Aria("label", "Toggle").
																						Aria("expanded", "false").
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-right").
																										Aria("hidden", true),
																								),
																						),
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-5").
																										Checked(true).
																										Aria("label", "Tree view node check checked"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-5").
																						Body(
																							app.Text("Loader"),
																						),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item pf-m-expanded").
										Aria("role", "treeitem").
										Aria("expanded", true).
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Button().
																		Class("pf-c-tree-view__node-toggle").
																		Aria("label", "Toggle").
																		Aria("expanded", true).
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-tree-view__node-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-6").
																						Aria("label", "Tree view node check"),
																				),
																		),
																	app.Label().
																		Class("pf-c-tree-view__node-text").
																		For("check-6").
																		Body(
																			app.Text("Application 2"),
																		),
																),
														),
												),
											app.Ul().
												Class("pf-c-tree-view__list").
												Aria("role", "group").
												Body(
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														Aria("expanded", "false").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Button().
																						Class("pf-c-tree-view__node-toggle").
																						Aria("label", "Toggle").
																						Aria("expanded", "false").
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-right").
																										Aria("hidden", true),
																								),
																						),
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-7").
																										Aria("label", "Tree view node check"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-7").
																						Body(
																							app.Text("Settings"),
																						),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														TabIndex(-1).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-8").
																										Aria("label", "Tree view node check"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-8").
																						Body(
																							app.Text("Settings"),
																						),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-tree-view__list-item pf-m-expanded").
														Aria("role", "treeitem").
														Aria("expanded", true).
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Button().
																						Class("pf-c-tree-view__node-toggle").
																						Aria("label", "Toggle").
																						Aria("expanded", true).
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-right").
																										Aria("hidden", true),
																								),
																						),
																					app.Span().
																						Class("pf-c-tree-view__node-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-9").
																										Aria("label", "Tree view node check"),
																								),
																						),
																					app.Label().
																						Class("pf-c-tree-view__node-text").
																						For("check-9").
																						Body(
																							app.Text("Current"),
																						),
																				),
																		),
																),
															app.Ul().
																Class("pf-c-tree-view__list").
																Aria("role", "group").
																Body(
																	app.Li().
																		Class("pf-c-tree-view__list-item").
																		Aria("role", "treeitem").
																		Aria("expanded", "false").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__content").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node").
																						TabIndex(0).
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Button().
																										Class("pf-c-tree-view__node-toggle").
																										Aria("label", "Toggle").
																										Aria("expanded", "false").
																										Body(
																											app.Span().
																												Class("pf-c-tree-view__node-toggle-icon").
																												Body(
																													app.I().
																														Class("fas fa-angle-right").
																														Aria("hidden", true),
																												),
																										),
																									app.Span().
																										Class("pf-c-tree-view__node-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														ID("check-10").
																														Aria("label", "Tree view node check"),
																												),
																										),
																									app.Label().
																										Class("pf-c-tree-view__node-text").
																										For("check-10").
																										Body(
																											app.Text("Loader app 1"),
																										),
																								),
																						),
																				),
																		),
																	app.Li().
																		Class("pf-c-tree-view__list-item").
																		Aria("role", "treeitem").
																		TabIndex(-1).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__content").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node").
																						TabIndex(0).
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Span().
																										Class("pf-c-tree-view__node-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														ID("check-11").
																														Checked(true).
																														Aria("label", "Tree view node check checked"),
																												),
																										),
																									app.Label().
																										Class("pf-c-tree-view__node-text").
																										For("check-11").
																										Body(
																											app.Text("Loader app 2"),
																										),
																								),
																						),
																				),
																		),
																	app.Li().
																		Class("pf-c-tree-view__list-item").
																		Aria("role", "treeitem").
																		TabIndex(-1).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__content").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node").
																						TabIndex(0).
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Span().
																										Class("pf-c-tree-view__node-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														ID("check-12").
																														Aria("label", "Tree view node check"),
																												),
																										),
																									app.Label().
																										Class("pf-c-tree-view__node-text").
																										For("check-12").
																										Body(
																											app.Text("Loader app 3"),
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
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						Aria("expanded", "false").
						TabIndex(0).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Button().
														Class("pf-c-tree-view__node-toggle").
														Aria("label", "Toggle").
														Aria("expanded", "false").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-tree-view__node-check").
														Body(
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("check-13").
																		Aria("label", "Tree view node check"),
																),
														),
													app.Label().
														Class("pf-c-tree-view__node-text").
														For("check-13").
														Body(
															app.Text("Cost management"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						Aria("expanded", "false").
						TabIndex(0).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Button().
														Class("pf-c-tree-view__node-toggle").
														Aria("label", "Toggle").
														Aria("expanded", "false").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-tree-view__node-check").
														Body(
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("check-14").
																		Aria("label", "Tree view node check"),
																),
														),
													app.Label().
														Class("pf-c-tree-view__node-text").
														For("check-14").
														Body(
															app.Text("Sources"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						Aria("expanded", "false").
						TabIndex(0).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Button().
														Class("pf-c-tree-view__node-toggle").
														Aria("label", "Toggle").
														Aria("expanded", "false").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
													app.Span().
														Class("pf-c-tree-view__node-check").
														Body(
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("check-15").
																		Checked(true).
																		Aria("label", "Tree view node check checked"),
																),
														),
													app.Label().
														Class("pf-c-tree-view__node-text").
														For("check-15").
														Body(
															app.Text("This is a really really really long folder name that overflows from the width of the container."),
														),
												),
										),
								),
						),
				),
		)
}
