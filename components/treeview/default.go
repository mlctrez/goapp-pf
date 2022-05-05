package treeview

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Default struct {
	app.Compo
}

func (c *Default) Render() app.UI {
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
									app.Button().
										Class("pf-c-tree-view__node").
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-toggle").
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
														Class("pf-c-tree-view__node-text").
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
													app.Button().
														Class("pf-c-tree-view__node").
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-toggle").
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
																		Class("pf-c-tree-view__node-text").
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
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-text").
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
																	app.Button().
																		Class("pf-c-tree-view__node pf-m-current").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-text").
																						Body(
																							app.Text("Current"),
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
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node-toggle").
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
																						Class("pf-c-tree-view__node-text").
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
													app.Button().
														Class("pf-c-tree-view__node").
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-toggle").
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
																		Class("pf-c-tree-view__node-text").
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
														TabIndex(-1).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Span().
																						Class("pf-c-tree-view__node-text").
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
														Aria("expanded", "false").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node-toggle").
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
																						Class("pf-c-tree-view__node-text").
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
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node-toggle").
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
																						Class("pf-c-tree-view__node-text").
																						Body(
																							app.Text("Loader"),
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
																					app.Button().
																						Class("pf-c-tree-view__node").
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Div().
																										Class("pf-c-tree-view__node-toggle").
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
																										Class("pf-c-tree-view__node-text").
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
																					app.Button().
																						Class("pf-c-tree-view__node").
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Span().
																										Class("pf-c-tree-view__node-text").
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
																					app.Button().
																						Class("pf-c-tree-view__node").
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Span().
																										Class("pf-c-tree-view__node-text").
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
									app.Button().
										Class("pf-c-tree-view__node").
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-toggle").
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
														Class("pf-c-tree-view__node-text").
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
									app.Button().
										Class("pf-c-tree-view__node").
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-toggle").
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
														Class("pf-c-tree-view__node-text").
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
									app.Button().
										Class("pf-c-tree-view__node").
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-toggle").
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
														Class("pf-c-tree-view__node-text").
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
