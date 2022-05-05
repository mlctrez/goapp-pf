package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableWithNestedTableExample struct {
	app.Compo
}

func (c *ExpandableWithNestedTableExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-expandable pf-m-grid-lg").
		Aria("role", "grid").
		Aria("label", "Expandable with nested table example").
		ID("table-expandable-nested-table").
		Body(
			app.THead().
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
										ID("table-expandable-nested-table-expandable-toggle-thead").
										Aria("label", "Expand all").
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
												Name("table-expandable-nested-table-checkrowthead").
												Aria("label", "Select all rows"),
										),
								),
							app.Th().
								Class("pf-m-width-30 pf-c-table__sort pf-m-selected").
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
															app.Text("Repositories"),
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
															app.Text("Branches"),
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
							app.Td(),
							app.Td(),
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
										Aria("labelledby", "table-expandable-nested-table-node1 table-expandable-nested-table-expandable-toggle1").
										ID("table-expandable-nested-table-expandable-toggle1").
										Aria("label", "Details").
										Aria("controls", "table-expandable-nested-table-content1").
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
												Name("table-expandable-nested-table-checkrow1").
												Aria("labelledby", "table-expandable-nested-table-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-expandable-nested-table-node1").
										Body(
											app.Text("Node 1"),
										),
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Link 1"),
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
												ID("table-expandable-nested-table-dropdown-kebab-1-button").
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
												Aria("labelledby", "table-expandable-nested-table-dropdown-kebab-1-button").
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
								Class("").
								Aria("role", "cell").
								ColSpan(7).
								ID("table-expandable-nested-table-content1").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Table().
												Class("pf-c-table pf-m-compact pf-m-grid-md").
												Aria("role", "grid").
												Aria("label", "This is a simple table").
												ID("table-expandable-nested-table-table-basic").
												Body(
													app.THead().
														Body(
															app.Tr().
																Aria("role", "row").
																Body(
																	app.Th().
																		Aria("role", "columnheader").
																		Scope("col").
																		Body(
																			app.Text("Repositories"),
																		),
																	app.Th().
																		Aria("role", "columnheader").
																		Scope("col").
																		Body(
																			app.Text("Branches"),
																		),
																	app.Th().
																		Aria("role", "columnheader").
																		Scope("col").
																		Body(
																			app.Text("Pull requests"),
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
																),
														),
													app.TBody().
														Aria("role", "rowgroup").
														Body(
															app.Tr().
																Aria("role", "row").
																Body(
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Repository name").
																		Body(
																			app.Text("Repository 1"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Branches").
																		Body(
																			app.Text("10"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Pull requests").
																		Body(
																			app.Text("25"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Workspaces").
																		Body(
																			app.Text("5"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Last commit").
																		Body(
																			app.Text("2 days ago"),
																		),
																),
															app.Tr().
																Aria("role", "row").
																Body(
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Repository name").
																		Body(
																			app.Text("Repository 2"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Branches").
																		Body(
																			app.Text("10"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Pull requests").
																		Body(
																			app.Text("25"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Workspaces").
																		Body(
																			app.Text("5"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Last commit").
																		Body(
																			app.Text("2 days ago"),
																		),
																),
															app.Tr().
																Aria("role", "row").
																Body(
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Repository name").
																		Body(
																			app.Text("Repository 3"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Branches").
																		Body(
																			app.Text("10"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Pull requests").
																		Body(
																			app.Text("25"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Workspaces").
																		Body(
																			app.Text("5"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Last commit").
																		Body(
																			app.Text("2 days ago"),
																		),
																),
															app.Tr().
																Aria("role", "row").
																Body(
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Repository name").
																		Body(
																			app.Text("Repository 4"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Branches").
																		Body(
																			app.Text("10"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Pull requests").
																		Body(
																			app.Text("25"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Workspaces").
																		Body(
																			app.Text("5"),
																		),
																	app.Td().
																		Aria("role", "cell").
																		DataSet("label", "Last commit").
																		Body(
																			app.Text("2 days ago"),
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
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-nested-table-node2 table-expandable-nested-table-expandable-toggle2").
										ID("table-expandable-nested-table-expandable-toggle2").
										Aria("label", "Details").
										Aria("controls", "table-expandable-nested-table-content2").
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
												Name("table-expandable-nested-table-checkrow2").
												Aria("labelledby", "table-expandable-nested-table-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-expandable-nested-table-node2").
										Body(
											app.Text("Node 2"),
										),
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Link 2"),
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
												ID("table-expandable-nested-table-dropdown-kebab-2-button").
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
												Aria("labelledby", "table-expandable-nested-table-dropdown-kebab-2-button").
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
								Class("").
								Aria("role", "cell").
								ColSpan(7).
								ID("table-expandable-nested-table-content2").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
										),
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
										Aria("labelledby", "table-expandable-nested-table-node3 table-expandable-nested-table-expandable-toggle3").
										ID("table-expandable-nested-table-expandable-toggle3").
										Aria("label", "Details").
										Aria("controls", "table-expandable-nested-table-content3").
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
												Name("table-expandable-nested-table-checkrow3").
												Aria("labelledby", "table-expandable-nested-table-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-expandable-nested-table-node3").
										Body(
											app.Text("Node 3"),
										),
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Link 3"),
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
												ID("table-expandable-nested-table-dropdown-kebab-3-button").
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
												Aria("labelledby", "table-expandable-nested-table-dropdown-kebab-3-button").
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
								Class("").
								Aria("role", "cell").
								ColSpan(7).
								ID("table-expandable-nested-table-content3").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
										),
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
										Aria("labelledby", "table-expandable-nested-table-node4 table-expandable-nested-table-expandable-toggle4").
										ID("table-expandable-nested-table-expandable-toggle4").
										Aria("label", "Details").
										Aria("controls", "table-expandable-nested-table-content4").
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
												Name("table-expandable-nested-table-checkrow4").
												Aria("labelledby", "table-expandable-nested-table-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-expandable-nested-table-node4").
										Body(
											app.Text("Node 4"),
										),
									app.A().
										Href("#").
										Body(
											app.Text("siemur/test-space"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Link 4"),
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
												ID("table-expandable-nested-table-dropdown-kebab-4-button").
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
												Aria("labelledby", "table-expandable-nested-table-dropdown-kebab-4-button").
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
								ID("table-expandable-nested-table-content4").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable row content has no padding."),
										),
								),
						),
				),
		)
}
