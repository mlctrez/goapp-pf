package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CompactExample struct {
	app.Compo
}

func (c *CompactExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-compact pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a compact table example").
		ID("table-compact").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-compact-checkrowthead").
												Aria("label", "Select all rows"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Contributor"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Position"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Location"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Last seen"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Numbers"),
								),
							app.Th().
								Class("pf-c-table__icon").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Icons"),
								),
							app.Td(),
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
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-compact-checkrow1").
												Aria("labelledby", "table-compact-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-node1").
										Body(
											app.Text("Sam Jones"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Position").
								Body(
									app.Text("CSS guru"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Location").
								Body(
									app.Text("Not too sure"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last seen").
								Body(
									app.Text("May 9, 2018"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Numbers").
								Body(
									app.Text("0556"),
								),
							app.Td().
								Class("pf-c-table__icon").
								Aria("role", "cell").
								DataSet("label", "Icon").
								Body(
									app.I().
										Class("fas fa-check"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Action link"),
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
												ID("table-compact-dropdown-kebab-1-button").
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
												Aria("labelledby", "table-compact-dropdown-kebab-1-button").
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
							app.Td().
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-compact-checkrow2").
												Aria("labelledby", "table-compact-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-node2").
										Body(
											app.Text("Amy Wilson"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Position").
								Body(
									app.Text("Visual design"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Location").
								Body(
									app.Text("Raleigh"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last seen").
								Body(
									app.Text("May 9, 2018"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Numbers").
								Body(
									app.Text("9492"),
								),
							app.Td().
								Class("pf-c-table__icon").
								Aria("role", "cell").
								DataSet("label", "Icon").
								Body(
									app.I().
										Class("fas fa-check"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Action link"),
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
												ID("table-compact-dropdown-kebab-2-button").
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
												Aria("labelledby", "table-compact-dropdown-kebab-2-button").
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
							app.Td().
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-compact-checkrow3").
												Aria("labelledby", "table-compact-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-node3").
										Body(
											app.Text("Steve Wilson"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Position").
								Body(
									app.Text("Visual design lead"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Location").
								Body(
									app.Text("Westford"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last seen").
								Body(
									app.Text("May 9, 2018"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Numbers").
								Body(
									app.Text("9929"),
								),
							app.Td().
								Class("pf-c-table__icon").
								Aria("role", "cell").
								DataSet("label", "Icon").
								Body(
									app.I().
										Class("fas fa-check"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Action link"),
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
												ID("table-compact-dropdown-kebab-3-button").
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
												Aria("labelledby", "table-compact-dropdown-kebab-3-button").
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
							app.Td().
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-compact-checkrow4").
												Aria("labelledby", "table-compact-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-compact-node4").
										Body(
											app.Text("Emma Jackson"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Position").
								Body(
									app.Text("Interaction design"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Location").
								Body(
									app.Text("Westford"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Workspaces").
								Body(
									app.Text("May 9, 2018"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Text("2217"),
								),
							app.Td().
								Class("pf-c-table__icon").
								Aria("role", "cell").
								DataSet("label", "Icon").
								Body(
									app.I().
										Class("fas fa-check"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Action").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("Action link"),
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
												ID("table-compact-dropdown-kebab-4-button").
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
												Aria("labelledby", "table-compact-dropdown-kebab-4-button").
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
		)
}
