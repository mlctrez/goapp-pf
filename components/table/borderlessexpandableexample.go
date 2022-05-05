package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BorderlessExpandableExample struct {
	app.Compo
}

func (c *BorderlessExpandableExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-expandable pf-m-grid-lg pf-m-no-border-rows").
		Aria("role", "grid").
		Aria("label", "Expandable table example").
		ID("borderless-table-expandable").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td(),
							app.Td().
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("borderless-table-expandable-checkrowthead").
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
										Aria("labelledby", "borderless-table-expandable-node1 borderless-table-expandable-expandable-toggle1").
										ID("borderless-table-expandable-expandable-toggle1").
										Aria("label", "Details").
										Aria("controls", "borderless-table-expandable-content1").
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
												Name("borderless-table-expandable-checkrow1").
												Aria("labelledby", "borderless-table-expandable-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("borderless-table-expandable-node1").
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
												ID("borderless-table-expandable-dropdown-kebab-1-button").
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
												Aria("labelledby", "borderless-table-expandable-dropdown-kebab-1-button").
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
								ColSpan(4).
								ID("borderless-table-expandable-content1").
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
										Aria("labelledby", "borderless-table-expandable-node2 borderless-table-expandable-expandable-toggle2").
										ID("borderless-table-expandable-expandable-toggle2").
										Aria("label", "Details").
										Aria("controls", "borderless-table-expandable-content2").
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
												Name("borderless-table-expandable-checkrow2").
												Aria("labelledby", "borderless-table-expandable-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("borderless-table-expandable-node2").
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
												ID("borderless-table-expandable-dropdown-kebab-2-button").
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
												Aria("labelledby", "borderless-table-expandable-dropdown-kebab-2-button").
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
								ID("borderless-table-expandable-content2").
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
										Aria("labelledby", "borderless-table-expandable-node3 borderless-table-expandable-expandable-toggle3").
										ID("borderless-table-expandable-expandable-toggle3").
										Aria("label", "Details").
										Aria("controls", "borderless-table-expandable-content3").
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
												Name("borderless-table-expandable-checkrow3").
												Aria("labelledby", "borderless-table-expandable-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("borderless-table-expandable-node3").
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
												ID("borderless-table-expandable-dropdown-kebab-3-button").
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
												Aria("labelledby", "borderless-table-expandable-dropdown-kebab-3-button").
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
								ID("borderless-table-expandable-content3").
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
										Aria("labelledby", "borderless-table-expandable-node4 borderless-table-expandable-expandable-toggle4").
										ID("borderless-table-expandable-expandable-toggle4").
										Aria("label", "Details").
										Aria("controls", "borderless-table-expandable-content4").
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
												Name("borderless-table-expandable-checkrow4").
												Aria("labelledby", "borderless-table-expandable-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("borderless-table-expandable-node4").
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
												ID("borderless-table-expandable-dropdown-kebab-4-button").
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
												Aria("labelledby", "borderless-table-expandable-dropdown-kebab-4-button").
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
								ID("borderless-table-expandable-content4").
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
