package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NestedColumnHeadersAndExpandableRows struct {
	app.Compo
}

func (c *NestedColumnHeadersAndExpandableRows) Render() app.UI {
	return app.Div().
		Class("pf-c-scroll-inner-wrapper").
		Body(
			app.Table().
				Class("pf-c-table pf-m-expandable").
				Aria("role", "grid").
				Aria("label", "This is a nested column header table example").
				ID("nested-columns-expandable-example").
				Body(
					app.THead().
						Class("pf-m-nested-column-header").
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Td().
										Rowspan(2),
									app.Td().
										Class("pf-c-table__check").
										Aria("role", "cell").
										Rowspan(2).
										Body(
											app.Label().
												Body(
													app.Input().
														Type("checkbox").
														Name("nested-columns-expandable-example-checkrow").
														Aria("label", "Select all rows"),
												),
										),
									app.Th().
										Class("pf-m-border-right pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
										Rowspan(2).
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
																	app.Text("Team"),
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
										Class("pf-m-border-right").
										Aria("role", "columnheader").
										Scope("col").
										ColSpan(3).
										Body(
											app.Text("Members"),
										),
									app.Th().
										Aria("role", "columnheader").
										Scope("col").
										Rowspan(2).
										Body(
											app.Text("Contact"),
										),
									app.Td().
										Rowspan(2),
								),
							app.Tr().
								Class("pf-m-first-cell-offset-reset").
								Aria("role", "row").
								Body(
									app.Th().
										Class("pf-c-table__subhead").
										Aria("role", "columnheader").
										Scope("col").
										Body(
											app.Text("Design lead"),
										),
									app.Th().
										Class("pf-c-table__subhead").
										Aria("role", "columnheader").
										Scope("col").
										Body(
											app.Text("Interaction design"),
										),
									app.Th().
										Class("pf-c-table__subhead pf-m-border-right").
										Aria("role", "columnheader").
										Scope("col").
										Body(
											app.Text("Visual designers"),
										),
								),
						),
					app.TBody().
						Class("pf-m-expanded").
						Aria("role", "rowgroup").
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Td().
										Class("pf-c-table__toggle").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain pf-m-expanded").
												Aria("labelledby", "nested-columns-expandable-example-node1 nested-columns-expandable-example-expandable-toggle1").
												ID("nested-columns-expandable-example-expandable-toggle1").
												Aria("label", "Details").
												Aria("controls", "nested-columns-expandable-example-content1").
												Aria("expanded", true).
												Body(
													app.Div().
														Class("pf-c-table__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-down").
																Aria("hidden", true),
														),
												),
										),
									app.Td().
										Class("pf-c-table__check").
										Aria("role", "cell").
										Body(
											app.Label().
												Body(
													app.Input().
														Type("checkbox").
														Name("nested-columns-expandable-example-checkrow1").
														Aria("labelledby", "nested-columns-expandable-example-node1"),
												),
										),
									app.Th().
										Class("").
										Aria("role", "columnheader").
										DataSet("label", "Developer program").
										ID("nested-columns-expandable-example-node1").
										Body(
											app.Text("Developer program"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Branches").
										Body(
											app.Text("Stacey Logan"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Pull requests").
										Body(
											app.Text("Mark Shakshober"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Workspaces").
										Body(
											app.Text("Kaliq Ray"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Last commit").
										Body(
											app.Button().
												Class("pf-c-button pf-m-inline pf-m-link").
												Type("button").
												Body(
													app.Text("Message us!"),
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
														ID("nested-columns-expandable-example-dropdown-kebab-1-button").
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
														Aria("labelledby", "nested-columns-expandable-example-dropdown-kebab-1-button").
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
									app.Td(),
									app.Td(),
									app.Td().
										Class("").
										Aria("role", "cell").
										ColSpan(5).
										ID("nested-columns-expandable-example-content1").
										Body(
											app.Div().
												Class("pf-c-table__expandable-row-content").
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
												),
										),
									app.Td(),
								),
						),
					app.TBody().
						Aria("role", "rowgroup").
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Td().
										Class("pf-c-table__toggle").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Aria("labelledby", "nested-columns-expandable-example-node2 nested-columns-expandable-example-expandable-toggle2").
												ID("nested-columns-expandable-example-expandable-toggle2").
												Aria("label", "Details").
												Aria("controls", "nested-columns-expandable-example-content2").
												Body(
													app.Div().
														Class("pf-c-table__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-down").
																Aria("hidden", true),
														),
												),
										),
									app.Td().
										Class("pf-c-table__check").
										Aria("role", "cell").
										Body(
											app.Label().
												Body(
													app.Input().
														Type("checkbox").
														Name("nested-columns-expandable-example-checkrow2").
														Aria("labelledby", "nested-columns-expandable-example-node2"),
												),
										),
									app.Th().
										Class("").
										Aria("role", "columnheader").
										DataSet("label", "Developer program").
										ID("nested-columns-expandable-example-node2").
										Body(
											app.Text("Developer program"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Branches").
										Body(
											app.Text("Stacey Logan"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Pull requests").
										Body(
											app.Text("Mark Shakshober"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Workspaces").
										Body(
											app.Text("Kaliq Ray"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Last commit").
										Body(
											app.Button().
												Class("pf-c-button pf-m-inline pf-m-link").
												Type("button").
												Body(
													app.Text("Message us!"),
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
														ID("nested-columns-expandable-example-dropdown-kebab-2-button").
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
														Aria("labelledby", "nested-columns-expandable-example-dropdown-kebab-2-button").
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
									app.Td(),
									app.Td(),
									app.Td().
										Class("").
										Aria("role", "cell").
										ColSpan(5).
										ID("nested-columns-expandable-example-content2").
										Body(
											app.Div().
												Class("pf-c-table__expandable-row-content").
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
												),
										),
									app.Td(),
								),
						),
					app.TBody().
						Aria("role", "rowgroup").
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Td().
										Class("pf-c-table__toggle").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Aria("labelledby", "nested-columns-expandable-example-node3 nested-columns-expandable-example-expandable-toggle3").
												ID("nested-columns-expandable-example-expandable-toggle3").
												Aria("label", "Details").
												Aria("controls", "nested-columns-expandable-example-content3").
												Body(
													app.Div().
														Class("pf-c-table__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-down").
																Aria("hidden", true),
														),
												),
										),
									app.Td().
										Class("pf-c-table__check").
										Aria("role", "cell").
										Body(
											app.Label().
												Body(
													app.Input().
														Type("checkbox").
														Name("nested-columns-expandable-example-checkrow3").
														Aria("labelledby", "nested-columns-expandable-example-node3"),
												),
										),
									app.Th().
										Class("").
										Aria("role", "columnheader").
										DataSet("label", "Developer program").
										ID("nested-columns-expandable-example-node3").
										Body(
											app.Text("Developer program"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Branches").
										Body(
											app.Text("Stacey Logan"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Pull requests").
										Body(
											app.Text("Mark Shakshober"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Workspaces").
										Body(
											app.Text("Kaliq Ray"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Last commit").
										Body(
											app.Button().
												Class("pf-c-button pf-m-inline pf-m-link").
												Type("button").
												Body(
													app.Text("Message us!"),
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
														ID("nested-columns-expandable-example-dropdown-kebab-3-button").
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
														Aria("labelledby", "nested-columns-expandable-example-dropdown-kebab-3-button").
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
									app.Td(),
									app.Td(),
									app.Td().
										Class("").
										Aria("role", "cell").
										ColSpan(5).
										ID("nested-columns-expandable-example-content3").
										Body(
											app.Div().
												Class("pf-c-table__expandable-row-content").
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
												),
										),
									app.Td(),
								),
						),
				),
		)
}
