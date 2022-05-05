package treeview

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithActionItem struct {
	app.Compo
}

func (c *WithActionItem) Render() app.UI {
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
									app.Div().
										Class("pf-c-tree-view__action").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Actions").
												Body(
													app.I().
														Class("fas fa-ellipsis-v").
														Aria("hidden", true),
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
													app.Div().
														Class("pf-c-tree-view__action").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Copy").
																Body(
																	app.I().
																		Class("fas fa-clipboard").
																		Aria("hidden", true),
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
													app.Div().
														Class("pf-c-tree-view__action").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Action").
																Body(
																	app.I().
																		Class("fas fa-bars").
																		Aria("hidden", true),
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
									app.Div().
										Class("pf-c-tree-view__action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("dropdown-kebab-button").
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
														Aria("labelledby", "dropdown-kebab-button").
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
										),
								),
						),
				),
		)
}
