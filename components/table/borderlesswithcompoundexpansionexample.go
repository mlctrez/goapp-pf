package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BorderlessWithCompoundExpansionExample struct {
	app.Compo
}

func (c *BorderlessWithCompoundExpansionExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-expandable pf-m-grid-md pf-m-no-border-rows").
		Aria("role", "grid").
		Aria("label", "Compound expandable table example").
		ID("borderless-compound-expansion-table").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__sort").
								Aria("role", "columnheader").
								Aria("sort", "none").
								Scope("col").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Div().
												Class("pf-c-table__button-content").
												Body(
													app.Span().
														Class("pf-c-table__text").
														Body(
															app.Text("Repositories"),
														),
													app.Span().
														Class("pf-c-table__sort-indicator").
														Body(
															app.I().
																Class("fas fa-arrows-alt-v"),
														),
												),
										),
								),
							app.Th().
								Class("pf-c-table__sort pf-m-selected").
								Aria("role", "columnheader").
								Aria("sort", "ascending").
								Scope("col").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Div().
												Class("pf-c-table__button-content").
												Body(
													app.Span().
														Class("pf-c-table__text").
														Body(
															app.Text("Branches"),
														),
													app.Span().
														Class("pf-c-table__sort-indicator").
														Body(
															app.I().
																Class("fas fa-long-arrow-alt-up"),
														),
												),
										),
								),
							app.Th().
								Class("pf-c-table__sort").
								Aria("role", "columnheader").
								Aria("sort", "none").
								Scope("col").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Div().
												Class("pf-c-table__button-content").
												Body(
													app.Span().
														Class("pf-c-table__text").
														Body(
															app.Text("Pull requests"),
														),
													app.Span().
														Class("pf-c-table__sort-indicator").
														Body(
															app.I().
																Class("fas fa-arrows-alt-v"),
														),
												),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Workspaces"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Last commit"),
								),
							app.Td(),
							app.Td(),
						),
				),
			app.TBody().
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Class("pf-c-table__control-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Repositories").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code-branch").
														Aria("hidden", true),
													app.Text("10"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code").
														Aria("hidden", true),
													app.Text("234"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-cube").
														Aria("hidden", true),
													app.Text("4"),
												),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Workspaces").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Span().
										Body(
											app.Text("20 minutes"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Open in Github"),
										),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("borderless-compound-expansion-table-dropdown-kebab-1-button").
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
												Aria("labelledby", "borderless-compound-expansion-table-dropdown-kebab-1-button").
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
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-1-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-1--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-2-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-2--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-3-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-3--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__control-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Repositories").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code-branch").
														Aria("hidden", true),
													app.Text("10"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code").
														Aria("hidden", true),
													app.Text("234"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-cube").
														Aria("hidden", true),
													app.Text("4"),
												),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Workspaces").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Span().
										Body(
											app.Text("20 minutes"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Open in Github"),
										),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("borderless-compound-expansion-table-dropdown-kebab-2-button").
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
												Aria("labelledby", "borderless-compound-expansion-table-dropdown-kebab-2-button").
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
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-4-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-4--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-5-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-5--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-6-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-6--dropdown-kebab-nested-tr5-button").
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
								),
						),
				),
			app.TBody().
				Class("pf-m-expanded").
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Class("pf-c-table__control-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Repositories").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code-branch").
														Aria("hidden", true),
													app.Text("2"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle pf-m-expanded").
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code").
														Aria("hidden", true),
													app.Text("82"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-cube").
														Aria("hidden", true),
													app.Text("1"),
												),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Workspaces").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Span().
										Body(
											app.Text("1 day ago"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Open in Github"),
										),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("borderless-compound-expansion-table-dropdown-kebab-3-button").
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
												Aria("labelledby", "borderless-compound-expansion-table-dropdown-kebab-3-button").
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
					app.Tr().
						Class("pf-c-table__expandable-row pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-7-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-7--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-8-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-8--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-9-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-9--dropdown-kebab-nested-tr5-button").
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
								),
						),
				),
			app.TBody().
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Class("pf-c-table__control-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Repositories").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code-branch").
														Aria("hidden", true),
													app.Text("4"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-code").
														Aria("hidden", true),
													app.Text("4"),
												),
										),
								),
							app.Td().
								Class("pf-c-table__compound-expansion-toggle").
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Body(
											app.Span().
												Class("pf-c-table__text").
												Body(
													app.I().
														Class("fas fa-cube").
														Aria("hidden", true),
													app.Text("1"),
												),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Workspaces").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Span().
										Body(
											app.Text("2 days ago"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Open in Github"),
										),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("borderless-compound-expansion-table-dropdown-kebab-4-button").
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
												Aria("labelledby", "borderless-compound-expansion-table-dropdown-kebab-4-button").
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
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-10-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-10--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-11-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-11--dropdown-kebab-nested-tr5-button").
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
								),
						),
					app.Tr().
						Class("pf-c-table__expandable-row").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-no-padding").
								Aria("role", "cell").
								ColSpan(7).
								Body(
									app.Table().
										Class("pf-c-table pf-m-compact pf-m-no-border-rows").
										Aria("role", "grid").
										ID("borderless-compound-expansion-table-nested-table-12-").
										Aria("label", "Nested table").
										Body(
											app.THead().
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Class("pf-c-table__sort").
																Aria("role", "columnheader").
																Aria("sort", "none").
																Scope("col").
																Body(
																	app.Button().
																		Class("pf-c-table__button").
																		Body(
																			app.Div().
																				Class("pf-c-table__button-content").
																				Body(
																					app.Span().
																						Class("pf-c-table__text").
																						Body(
																							app.Text("Description"),
																						),
																					app.Span().
																						Class("pf-c-table__sort-indicator").
																						Body(
																							app.I().
																								Class("fas fa-arrows-alt-v"),
																						),
																				),
																		),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Date"),
																),
															app.Th().
																Aria("role", "columnheader").
																Scope("col").
																Body(
																	app.Text("Status"),
																),
															app.Td().
																Aria("role", "cell"),
														),
												),
											app.TBody().
												Aria("role", "rowgroup").
												Body(
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item one"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr1-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr1-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item two"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Warning"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr2-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr2-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item three"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr3-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr3-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item four"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr4-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr4-button").
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
													app.Tr().
														Aria("role", "row").
														Body(
															app.Th().
																Aria("role", "columnheader").
																DataSet("label", "Description").
																Body(
																	app.Text("Item five"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Date").
																Body(
																	app.Text("May 9, 2018"),
																),
															app.Td().
																Aria("role", "cell").
																DataSet("label", "Status").
																Body(
																	app.Text("Active"),
																),
															app.Td().
																Class("pf-c-table__action").
																Aria("role", "cell").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__toggle pf-m-plain").
																				ID("borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr5-button").
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
																				Aria("labelledby", "borderless-compound-expansion-table-nested-table-12--dropdown-kebab-nested-tr5-button").
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
								),
						),
				),
		)
}
