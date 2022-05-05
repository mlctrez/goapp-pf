package duallistselector

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TreeViewWithChosenAndDisabledOptions struct {
	app.Compo
}

func (c *TreeViewWithChosenAndDisabledOptions) Render() app.UI {
	return app.Div().
		Class("pf-c-dual-list-selector").
		Body(
			app.Div().
				Class("pf-c-dual-list-selector__pane pf-m-available").
				Body(
					app.Div().
						Class("pf-c-dual-list-selector__header").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__title").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__title-text").
										Body(
											app.Text("Available options"),
										),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__tools").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__tools-filter").
								Body(
									app.Div().
										Class("pf-c-search-input").
										Body(
											app.Div().
												Class("pf-c-search-input__bar").
												Body(
													app.Span().
														Class("pf-c-search-input__text").
														Body(
															app.Span().
																Class("pf-c-search-input__icon").
																Body(
																	app.I().
																		Class("fas fa-search fa-fw").
																		Aria("hidden", true),
																),
															app.Input().
																Class("pf-c-search-input__text-input").
																Type("text").
																Placeholder("false").
																Aria("label", "Filter options"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__tools-actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Sort").
										Body(
											app.I().
												Class("fas fa-sort-amount-down").
												Aria("hidden", true),
										),
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("dropdown-kebab-tree-options-available-button").
												Aria("expanded", "false").
												Type("button").
												Aria("label", "Actions").
												Body(
													app.I().
														Class("fas fa-ellipsis-v").
														Aria("hidden", true),
												),
											app.Ul().
												Class("pf-c-dropdown__menu").
												Aria("labelledby", "dropdown-kebab-tree-options-available-button").
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
					app.Div().
						Class("pf-c-dual-list-selector__status").
						Body(
							app.Span().
								Class("pf-c-dual-list-selector__status-text").
								ID("tree-options-available-status-text").
								Body(
									app.Text("0 of 10 items selected"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__menu").
						Body(
							app.Ul().
								Class("pf-c-dual-list-selector__list").
								Aria("role", "tree").
								Aria("labelledby", "tree-options-available-status-text").
								Aria("multiselectable", true).
								Aria("activedescendant", true).
								Body(
									app.Li().
										Class("pf-c-dual-list-selector__list-item pf-m-expandable pf-m-expanded").
										Aria("expanded", true).
										Aria("role", "treeitem").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__list-item-row pf-m-check").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__item").
														TabIndex(0).
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item-main").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item-toggle").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-11").
																						Aria("label", "Dual list selector item check"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-text").
																		Body(
																			app.Text("Colors"),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item-count").
																Body(
																	app.Span().
																		Class("pf-c-badge pf-m-read").
																		Body(
																			app.Text("6"),
																		),
																),
														),
												),
											app.Ul().
												Class("pf-c-dual-list-selector__list").
												Aria("role", "group").
												Aria("labelledby", "-status-text").
												Body(
													app.Li().
														Class("pf-c-dual-list-selector__list-item").
														Aria("role", "treeitem").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__list-item-row pf-m-check pf-m-selected").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item").
																		TabIndex(0).
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-dual-list-selector__item-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-12").
																										Checked(true).
																										Aria("label", "Dual list selector item check checked"),
																								),
																						),
																					app.Span().
																						Class("pf-c-dual-list-selector__item-text").
																						Body(
																							app.Text("Orange"),
																						),
																				),
																			app.Span().
																				Class("pf-c-dual-list-selector__item-count").
																				Body(
																					app.Span().
																						Class("pf-c-badge pf-m-read"),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-dual-list-selector__list-item").
														Aria("role", "treeitem").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__list-item-row pf-m-check").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item").
																		TabIndex(0).
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-dual-list-selector__item-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-13").
																										Aria("label", "Dual list selector item check"),
																								),
																						),
																					app.Span().
																						Class("pf-c-dual-list-selector__item-text").
																						Body(
																							app.Text("Yellow"),
																						),
																				),
																			app.Span().
																				Class("pf-c-dual-list-selector__item-count").
																				Body(
																					app.Span().
																						Class("pf-c-badge pf-m-read"),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-dual-list-selector__list-item pf-m-expandable pf-m-expanded pf-m-disabled").
														Aria("expanded", true).
														Aria("role", "treeitem").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__list-item-row pf-m-check").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-main").
																				Body(
																					app.Div().
																						Class("pf-c-dual-list-selector__item-toggle").
																						Body(
																							app.Span().
																								Class("pf-c-dual-list-selector__item-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-right").
																										Aria("hidden", true),
																								),
																						),
																					app.Span().
																						Class("pf-c-dual-list-selector__item-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										Disabled(true).
																										ID("check-14").
																										Aria("label", "Dual list selector item check"),
																								),
																						),
																					app.Span().
																						Class("pf-c-dual-list-selector__item-text").
																						Body(
																							app.Text("Green (disabled)"),
																						),
																				),
																			app.Span().
																				Class("pf-c-dual-list-selector__item-count").
																				Body(
																					app.Span().
																						Class("pf-c-badge pf-m-read"),
																				),
																		),
																),
															app.Ul().
																Class("pf-c-dual-list-selector__list").
																Aria("role", "group").
																Aria("labelledby", "-status-text").
																Body(
																	app.Li().
																		Class("pf-c-dual-list-selector__list-item").
																		Aria("role", "treeitem").
																		Body(
																			app.Div().
																				Class("pf-c-dual-list-selector__list-item-row pf-m-check").
																				Body(
																					app.Div().
																						Class("pf-c-dual-list-selector__item").
																						Body(
																							app.Span().
																								Class("pf-c-dual-list-selector__item-main").
																								Body(
																									app.Span().
																										Class("pf-c-dual-list-selector__item-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														Disabled(true).
																														ID("check-15").
																														Aria("label", "Dual list selector item check"),
																												),
																										),
																									app.Span().
																										Class("pf-c-dual-list-selector__item-text").
																										Body(
																											app.Text("Light green"),
																										),
																								),
																							app.Span().
																								Class("pf-c-dual-list-selector__item-count").
																								Body(
																									app.Span().
																										Class("pf-c-badge pf-m-read"),
																								),
																						),
																				),
																		),
																	app.Li().
																		Class("pf-c-dual-list-selector__list-item").
																		Aria("role", "treeitem").
																		Body(
																			app.Div().
																				Class("pf-c-dual-list-selector__list-item-row pf-m-check").
																				Body(
																					app.Div().
																						Class("pf-c-dual-list-selector__item").
																						Body(
																							app.Span().
																								Class("pf-c-dual-list-selector__item-main").
																								Body(
																									app.Span().
																										Class("pf-c-dual-list-selector__item-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														Disabled(true).
																														ID("check-16").
																														Aria("label", "Dual list selector item check"),
																												),
																										),
																									app.Span().
																										Class("pf-c-dual-list-selector__item-text").
																										Body(
																											app.Text("Medium green"),
																										),
																								),
																							app.Span().
																								Class("pf-c-dual-list-selector__item-count").
																								Body(
																									app.Span().
																										Class("pf-c-badge pf-m-read"),
																								),
																						),
																				),
																		),
																	app.Li().
																		Class("pf-c-dual-list-selector__list-item").
																		Aria("role", "treeitem").
																		Body(
																			app.Div().
																				Class("pf-c-dual-list-selector__list-item-row pf-m-check").
																				Body(
																					app.Div().
																						Class("pf-c-dual-list-selector__item").
																						Body(
																							app.Span().
																								Class("pf-c-dual-list-selector__item-main").
																								Body(
																									app.Span().
																										Class("pf-c-dual-list-selector__item-check").
																										Body(
																											app.Div().
																												Class("pf-c-check pf-m-standalone").
																												Body(
																													app.Input().
																														Class("pf-c-check__input").
																														Type("checkbox").
																														Disabled(true).
																														ID("check-17").
																														Aria("label", "Dual list selector item check"),
																												),
																										),
																									app.Span().
																										Class("pf-c-dual-list-selector__item-text").
																										Body(
																											app.Text("Dark green"),
																										),
																								),
																							app.Span().
																								Class("pf-c-dual-list-selector__item-count").
																								Body(
																									app.Span().
																										Class("pf-c-badge pf-m-read"),
																								),
																						),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-dual-list-selector__list-item pf-m-expandable").
										Aria("expanded", true).
										Aria("role", "treeitem").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__list-item-row pf-m-check").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__item").
														TabIndex(0).
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item-main").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item-toggle").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-18").
																						Aria("label", "Dual list selector item check"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-text").
																		Body(
																			app.Text("Type something"),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item-count").
																Body(
																	app.Span().
																		Class("pf-c-badge pf-m-read"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-dual-list-selector__list-item").
										Aria("role", "treeitem").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__list-item-row pf-m-check").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__item").
														TabIndex(0).
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item-main").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-19").
																						Aria("label", "Dual list selector item check"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-text").
																		Body(
																			app.Text("Type something"),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item-count").
																Body(
																	app.Span().
																		Class("pf-c-badge pf-m-read"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-dual-list-selector__list-item pf-m-expandable").
										Aria("expanded", true).
										Aria("role", "treeitem").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__list-item-row pf-m-check").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__item").
														TabIndex(0).
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item-main").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item-toggle").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-20").
																						Aria("label", "Dual list selector item check"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-text").
																		Body(
																			app.Text("Type something"),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item-count").
																Body(
																	app.Span().
																		Class("pf-c-badge pf-m-read"),
																),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-dual-list-selector__controls").
				Body(
					app.Div().
						Class("pf-c-dual-list-selector__controls-item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Disabled(true).
								Aria("label", "Add selected").
								Body(
									app.I().
										Class("fas fa-fw fa-angle-right"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__controls-item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Add all").
								Body(
									app.I().
										Class("fas fa-fw fa-angle-double-right"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__controls-item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Disabled(true).
								Aria("label", "Remove all").
								Body(
									app.I().
										Class("fas fa-fw fa-angle-double-left"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__controls-item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Disabled(true).
								Aria("label", "Remove selected").
								Body(
									app.I().
										Class("fas fa-fw fa-angle-left"),
								),
						),
				),
			app.Div().
				Class("pf-c-dual-list-selector__pane pf-m-chosen").
				Body(
					app.Div().
						Class("pf-c-dual-list-selector__header").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__title").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__title-text").
										Body(
											app.Text("Chosen options"),
										),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__tools").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__tools-filter").
								Body(
									app.Div().
										Class("pf-c-search-input").
										Body(
											app.Div().
												Class("pf-c-search-input__bar").
												Body(
													app.Span().
														Class("pf-c-search-input__text").
														Body(
															app.Span().
																Class("pf-c-search-input__icon").
																Body(
																	app.I().
																		Class("fas fa-search fa-fw").
																		Aria("hidden", true),
																),
															app.Input().
																Class("pf-c-search-input__text-input").
																Type("text").
																Placeholder("false").
																Aria("label", "Filter options"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__tools-actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Sort").
										Body(
											app.I().
												Class("fas fa-sort-amount-down").
												Aria("hidden", true),
										),
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("dropdown-kebab-tree-options-chosen-button").
												Aria("expanded", "false").
												Type("button").
												Aria("label", "Actions").
												Body(
													app.I().
														Class("fas fa-ellipsis-v").
														Aria("hidden", true),
												),
											app.Ul().
												Class("pf-c-dropdown__menu").
												Aria("labelledby", "dropdown-kebab-tree-options-chosen-button").
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
					app.Div().
						Class("pf-c-dual-list-selector__status").
						Body(
							app.Span().
								Class("pf-c-dual-list-selector__status-text").
								ID("tree-options-chosen-status-text").
								Body(
									app.Text("0 of 0 items selected"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__menu").
						Body(
							app.Ul().
								Class("pf-c-dual-list-selector__list").
								Aria("role", "tree").
								Aria("labelledby", "tree-options-chosen-status-text").
								Aria("multiselectable", true).
								Aria("activedescendant", true).
								Body(
									app.Li().
										Class("pf-c-dual-list-selector__list-item pf-m-expandable pf-m-expanded").
										Aria("expanded", true).
										Aria("role", "treeitem").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__list-item-row pf-m-check").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__item").
														TabIndex(0).
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item-main").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item-toggle").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-check").
																		Body(
																			app.Div().
																				Class("pf-c-check pf-m-standalone").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("check-21").
																						Aria("label", "Dual list selector item check"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-text").
																		Body(
																			app.Text("Colors"),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item-count").
																Body(
																	app.Span().
																		Class("pf-c-badge pf-m-read"),
																),
														),
												),
											app.Ul().
												Class("pf-c-dual-list-selector__list").
												Aria("role", "group").
												Aria("labelledby", "-status-text").
												Body(
													app.Li().
														Class("pf-c-dual-list-selector__list-item").
														Aria("role", "treeitem").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__list-item-row pf-m-check pf-m-selected").
																Body(
																	app.Div().
																		Class("pf-c-dual-list-selector__item").
																		TabIndex(0).
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-main").
																				Body(
																					app.Span().
																						Class("pf-c-dual-list-selector__item-check").
																						Body(
																							app.Div().
																								Class("pf-c-check pf-m-standalone").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("check-22").
																										Checked(true).
																										Aria("label", "Dual list selector item check checked"),
																								),
																						),
																					app.Span().
																						Class("pf-c-dual-list-selector__item-text").
																						Body(
																							app.Text("Orange"),
																						),
																				),
																			app.Span().
																				Class("pf-c-dual-list-selector__item-count").
																				Body(
																					app.Span().
																						Class("pf-c-badge pf-m-read"),
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
