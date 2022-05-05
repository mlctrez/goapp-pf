package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TreeTableWithCheckboxesAndIcons struct {
	app.Compo
}

func (c *TreeTableWithCheckboxesAndIcons) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-tree-view pf-m-tree-view-checkboxes pf-m-tree-view-grid-lg").
		Aria("role", "treegrid").
		Aria("label", "This is a simple tree table, with checkboxes and icons example").
		ID("tree-table-with-checkboxes-icons-example").
		Body(
			app.THead().
				Body(
					app.Tr().
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-header-cell").
								Scope("col").
								Body(
									app.Text("Name"),
								),
							app.Th().
								Scope("col").
								Body(
									app.Text("Count"),
								),
							app.Th().
								Scope("col").
								Body(
									app.Text("Size"),
								),
							app.Th().
								Scope("col").
								Body(
									app.Text("Data Stores"),
								),
						),
				),
			app.TBody().
				Body(
					app.Tr().
						Class("").
						Aria("level", "1").
						Aria("setsize", "1").
						Aria("posinset", "1").
						TabIndex(0).
						Aria("expanded", true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-1 tree-table-with-checkboxes-icons-example-expandable-toggle-1").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-1").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-1").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-1"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-1").
														Body(
															app.Text("Level 1 all folders"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-1--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-1-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-1-actions-button").
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
						Class("pf-m-tree-view-details-expanded").
						Aria("level", "2").
						Aria("setsize", "5").
						Aria("posinset", "1").
						TabIndex(0).
						Aria("expanded", "false").
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-2 tree-table-with-checkboxes-icons-example-expandable-toggle-2").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-2").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-2").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-2"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-2").
														Body(
															app.Text("Level 2 node"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("expanded", true).
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-2--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-2-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-2-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "1").
						TabIndex(-1).
						Hidden(true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-3").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-3"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-3").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-3--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-3-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-3-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "2").
						TabIndex(-1).
						Hidden(true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-4").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-4"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-4").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-4--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-4-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-4-actions-button").
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
						Class("").
						Aria("level", "2").
						Aria("setsize", "5").
						Aria("posinset", "2").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-5").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-5"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-5").
														Body(
															app.Text("Level 2 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-5--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-5-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-5-actions-button").
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
						Class("").
						Aria("level", "2").
						Aria("setsize", "5").
						Aria("posinset", "3").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-6").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-6"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-6").
														Body(
															app.Text("Level 2 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-6--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-6-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-6-actions-button").
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
						Class("").
						Aria("level", "2").
						Aria("setsize", "5").
						Aria("posinset", "4").
						TabIndex(0).
						Aria("expanded", true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-7 tree-table-with-checkboxes-icons-example-expandable-toggle-7").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-7").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-7").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-7"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-7").
														Body(
															app.Text("Level 2 node"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-7--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-7-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-7-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "3").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-8").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-8"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-8").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-8--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-8-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-8-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "1").
						TabIndex(0).
						Aria("expanded", true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-9 tree-table-with-checkboxes-icons-example-expandable-toggle-9").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-9").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-9").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-9"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-9").
														Body(
															app.Text("Level 3 node"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-9--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-9-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-9-actions-button").
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
						Class("pf-m-tree-view-details-expanded").
						Aria("level", "4").
						Aria("setsize", "1").
						Aria("posinset", "1").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-10").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-10"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-10").
														Body(
															app.Text("Level 4 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("expanded", true).
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-10--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-10-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-10-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "2").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-11").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-11"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-11").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-11--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-11-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-11-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "5").
						TabIndex(0).
						Aria("expanded", true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-expanded").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-12 tree-table-with-checkboxes-icons-example-expandable-toggle-12").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-12").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-12").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-12"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-12").
														Body(
															app.Text("Level 3 node"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-12--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-12-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-12-actions-button").
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
						Class("").
						Aria("level", "4").
						Aria("setsize", "2").
						Aria("posinset", "1").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-13").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-13"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-13").
														Body(
															app.Text("Level 4 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-13--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-13-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-13-actions-button").
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
						Class("").
						Aria("level", "4").
						Aria("setsize", "2").
						Aria("posinset", "2").
						TabIndex(-1).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-14").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-14"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-14").
														Body(
															app.Text("Level 4 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-14--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-14-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-14-actions-button").
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
						Class("").
						Aria("level", "2").
						Aria("setsize", "5").
						Aria("posinset", "6").
						TabIndex(0).
						Aria("expanded", "false").
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-15 tree-table-with-checkboxes-icons-example-expandable-toggle-15").
														ID("tree-table-with-checkboxes-icons-example-expandable-toggle-15").
														Aria("label", "Details").
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
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-15").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-15"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-folder").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-15").
														Body(
															app.Text("Level 2 node"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-15--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-15-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-15-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "1").
						TabIndex(-1).
						Hidden(true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-16").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-16"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-16").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-16--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-16-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-16-actions-button").
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
						Class("").
						Aria("level", "3").
						Aria("setsize", "2").
						Aria("posinset", "2").
						TabIndex(-1).
						Hidden(true).
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__tree-view-title-cell").
								Body(
									app.Div().
										Class("pf-c-table__tree-view-main").
										Body(
											app.Span().
												Class("pf-c-table__check").
												Body(
													app.Label().
														Body(
															app.Input().
																Type("checkbox").
																Name("tree-table-with-checkboxes-icons-example-checkrow-17").
																Aria("labelledby", "tree-table-with-checkboxes-icons-example-node-17"),
														),
												),
											app.Div().
												Class("pf-c-table__tree-view-text").
												Body(
													app.Span().
														Class("pf-c-table__tree-view-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-leaf").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__text").
														ID("tree-table-with-checkboxes-icons-example-node-17").
														Body(
															app.Text("Level 3 leaf"),
														),
												),
											app.Span().
												Class("pf-c-table__tree-view-details-toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "tree-table-with-checkboxes-icons-example-17--tree-table--details-toggle").
														Body(
															app.Span().
																Class("pf-c-table__details-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-h").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Migration assessment").
								Body(
									app.Text("10"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Size of VM").
								Body(
									app.Text("25"),
								),
							app.Td().
								Aria("role", "gridcell").
								DataSet("label", "Number of Data Stores").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-c-table__action").
								Aria("role", "gridcell").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("tree-table-with-checkboxes-icons-example-17-actions-button").
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
												Aria("labelledby", "tree-table-with-checkboxes-icons-example-17-actions-button").
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
