package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type OverflowMenuUsageOverflowMenuCollapsed struct {
	app.Compo
}

func (c *OverflowMenuUsageOverflowMenuCollapsed) Render() app.UI {
	return app.Table().
		Class("pf-c-table").
		Aria("role", "grid").
		Aria("label", "This is a simple table example").
		ID("table-with-overflow-menu-collapsed").
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
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("table-with-overflow-menu-collapsed-overflow-menu-1").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__control").
												Body(
													app.Div().
														Class("pf-c-dropdown pf-m-expanded").
														Body(
															app.Button().
																Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																Type("button").
																ID("table-with-overflow-menu-collapsed-overflow-menu-1-dropdown-toggle").
																Aria("label", "Generic options").
																Aria("expanded", true).
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu pf-m-align-right").
																Aria("labelledby", "table-with-overflow-menu-collapsed-overflow-menu-1-dropdown-toggle").
																Body(
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Start"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Stop"),
																				),
																		),
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
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("table-with-overflow-menu-collapsed-overflow-menu-2").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__control").
												Body(
													app.Div().
														Class("pf-c-dropdown").
														Body(
															app.Button().
																Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																Type("button").
																ID("table-with-overflow-menu-collapsed-overflow-menu-2-dropdown-toggle").
																Aria("label", "Generic options").
																Aria("expanded", true).
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu pf-m-align-right").
																Aria("labelledby", "table-with-overflow-menu-collapsed-overflow-menu-2-dropdown-toggle").
																Hidden(true).
																Body(
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Start"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Stop"),
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
