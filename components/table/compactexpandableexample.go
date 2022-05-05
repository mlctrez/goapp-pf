package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CompactExpandableExample struct {
	app.Compo
}

func (c *CompactExpandableExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-expandable pf-m-compact pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "Compact expandable table example").
		ID("table-compact-expandable").
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
										ID("table-compact-expandable-expandable-toggle-thead").
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
												Name("table-compact-expandable-checkrowthead").
												Aria("label", "Select all rows"),
										),
								),
							app.Th().
								Class("pf-m-width-30").
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
										Aria("labelledby", "table-compact-expandable-node1 table-compact-expandable-expandable-toggle1").
										ID("table-compact-expandable-expandable-toggle1").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content1").
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
												Name("table-compact-expandable-checkrow1").
												Aria("labelledby", "table-compact-expandable-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node1").
										Body(
											app.Text("Node 1"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-1-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-1-button").
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
								Aria("role", "cell").
								ColSpan(4).
								ID("table-compact-expandable-content1").
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
										Aria("labelledby", "table-compact-expandable-node2 table-compact-expandable-expandable-toggle2").
										ID("table-compact-expandable-expandable-toggle2").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content2").
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
												Name("table-compact-expandable-checkrow2").
												Aria("labelledby", "table-compact-expandable-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node2").
										Body(
											app.Text("Node 2"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-2-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-2-button").
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
								ID("table-compact-expandable-content2").
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
										Aria("labelledby", "table-compact-expandable-node3 table-compact-expandable-expandable-toggle3").
										ID("table-compact-expandable-expandable-toggle3").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content3").
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
												Name("table-compact-expandable-checkrow3").
												Aria("labelledby", "table-compact-expandable-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node3").
										Body(
											app.Text("Node 3"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-3-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-3-button").
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
								Aria("role", "cell").
								ColSpan(7).
								ID("table-compact-expandable-content3").
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
										Aria("labelledby", "table-compact-expandable-node4 table-compact-expandable-expandable-toggle4").
										ID("table-compact-expandable-expandable-toggle4").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content4").
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
												Name("table-compact-expandable-checkrow4").
												Aria("labelledby", "table-compact-expandable-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node4").
										Body(
											app.Text("Node 4"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-4-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-4-button").
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
								ID("table-compact-expandable-content4").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("This content has no padding."),
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
										Aria("labelledby", "table-compact-expandable-node5 table-compact-expandable-expandable-toggle5").
										ID("table-compact-expandable-expandable-toggle5").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content5").
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
												Name("table-compact-expandable-checkrow5").
												Aria("labelledby", "table-compact-expandable-node5"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node5").
										Body(
											app.Text("Node 5"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-5-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-5-button").
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
								Aria("role", "cell").
								ColSpan(7).
								ID("table-compact-expandable-content5").
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
										Aria("labelledby", "table-compact-expandable-node6 table-compact-expandable-expandable-toggle6").
										ID("table-compact-expandable-expandable-toggle6").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content6").
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
												Name("table-compact-expandable-checkrow6").
												Aria("labelledby", "table-compact-expandable-node6"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node6").
										Body(
											app.Text("Node 6"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-6-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-6-button").
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
						ID("table-compact-expandable-content6").
						Body(
							app.Td(),
							app.Td(),
							app.Td().
								Aria("role", "cell").
								ColSpan(2).
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."),
										),
								),
							app.Td().
								Aria("role", "cell").
								ColSpan(2).
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
										Aria("labelledby", "table-compact-expandable-node7 table-compact-expandable-expandable-toggle7").
										ID("table-compact-expandable-expandable-toggle7").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content7").
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
												Name("table-compact-expandable-checkrow7").
												Aria("labelledby", "table-compact-expandable-node7"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node7").
										Body(
											app.Text("Node 7"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-7-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-7-button").
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
								Aria("role", "cell").
								ColSpan(7).
								ID("table-compact-expandable-content7").
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
										Aria("labelledby", "table-compact-expandable-node8 table-compact-expandable-expandable-toggle8").
										ID("table-compact-expandable-expandable-toggle8").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content8").
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
												Name("table-compact-expandable-checkrow8").
												Aria("labelledby", "table-compact-expandable-node8"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node8").
										Body(
											app.Text("Node 8"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-8-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-8-button").
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
						ID("table-compact-expandable-content8").
						Body(
							app.Td().
								Aria("role", "cell").
								ColSpan(4).
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."),
										),
								),
							app.Td().
								Aria("role", "cell").
								ColSpan(3).
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
										Aria("labelledby", "table-compact-expandable-node9 table-compact-expandable-expandable-toggle9").
										ID("table-compact-expandable-expandable-toggle9").
										Aria("label", "Details").
										Aria("controls", "table-compact-expandable-content9").
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
												Name("table-compact-expandable-checkrow9").
												Aria("labelledby", "table-compact-expandable-node9"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-expandable-node9").
										Body(
											app.Text("Node 9"),
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
											app.Text("Link"),
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
												ID("table-compact-expandable-dropdown-kebab-9-button").
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
												Aria("labelledby", "table-compact-expandable-dropdown-kebab-9-button").
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
								Aria("role", "cell").
								ColSpan(7).
								ID("table-compact-expandable-content9").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
										),
								),
						),
				),
		)
}
