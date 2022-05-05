package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableHoverableAndSelectedExample struct {
	app.Compo
}

func (c *ExpandableHoverableAndSelectedExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-expandable pf-m-grid-lg").
		Aria("role", "grid").
		Aria("label", "Expandable and hoverable table example").
		ID("table-expandable-hoverable").
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
												Name("table-expandable-hoverable-checkrowthead").
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
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node1 table-expandable-hoverable-expandable-toggle1").
										ID("table-expandable-hoverable-expandable-toggle1").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content1").
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
												Name("table-expandable-hoverable-checkrow1").
												Aria("labelledby", "table-expandable-hoverable-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node1").
												Body(
													app.Text("Hoverable"),
												),
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
												ID("table-expandable-hoverable-dropdown-kebab-1-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-1-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content1").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node2 table-expandable-hoverable-expandable-toggle2").
										ID("table-expandable-hoverable-expandable-toggle2").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content2").
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
												Name("table-expandable-hoverable-checkrow2").
												Aria("labelledby", "table-expandable-hoverable-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node2").
												Body(
													app.I().
														Body(
															app.Text("Selected and not expanded"),
														),
												),
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
												ID("table-expandable-hoverable-dropdown-kebab-2-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-2-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content2").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node3 table-expandable-hoverable-expandable-toggle3").
										ID("table-expandable-hoverable-expandable-toggle3").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content3").
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
												Name("table-expandable-hoverable-checkrow3").
												Aria("labelledby", "table-expandable-hoverable-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node3").
												Body(
													app.Text("Hoverable"),
												),
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
												ID("table-expandable-hoverable-dropdown-kebab-3-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-3-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content3").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node4 table-expandable-hoverable-expandable-toggle4").
										ID("table-expandable-hoverable-expandable-toggle4").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content4").
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
												Name("table-expandable-hoverable-checkrow4").
												Aria("labelledby", "table-expandable-hoverable-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node4").
												Body(
													app.Text("Hoverable"),
												),
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
												ID("table-expandable-hoverable-dropdown-kebab-4-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-4-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content4").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node5 table-expandable-hoverable-expandable-toggle5").
										ID("table-expandable-hoverable-expandable-toggle5").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content5").
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
												Name("table-expandable-hoverable-checkrow5").
												Aria("labelledby", "table-expandable-hoverable-node5"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node5").
												Body(
													app.I().
														Body(
															app.Text("Selected and not expanded"),
														),
												),
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
											app.Text("Link 5"),
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
												ID("table-expandable-hoverable-dropdown-kebab-5-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-5-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content5").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node6 table-expandable-hoverable-expandable-toggle6").
										ID("table-expandable-hoverable-expandable-toggle6").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content6").
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
												Name("table-expandable-hoverable-checkrow6").
												Aria("labelledby", "table-expandable-hoverable-node6"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node6").
												Body(
													app.I().
														Body(
															app.Text("Selected and not expanded"),
														),
												),
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
											app.Text("Link 6"),
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
												ID("table-expandable-hoverable-dropdown-kebab-6-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-6-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content6").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node7 table-expandable-hoverable-expandable-toggle7").
										ID("table-expandable-hoverable-expandable-toggle7").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content7").
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
												Name("table-expandable-hoverable-checkrow7").
												Aria("labelledby", "table-expandable-hoverable-node7"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node7").
												Body(
													app.I().
														Body(
															app.Text("Selected and not expanded"),
														),
												),
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
											app.Text("Link 7"),
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
												ID("table-expandable-hoverable-dropdown-kebab-7-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-7-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content7").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node8 table-expandable-hoverable-expandable-toggle8").
										ID("table-expandable-hoverable-expandable-toggle8").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content8").
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
												Name("table-expandable-hoverable-checkrow8").
												Aria("labelledby", "table-expandable-hoverable-node8"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node8").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 8"),
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
												ID("table-expandable-hoverable-dropdown-kebab-8-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-8-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content8").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node9 table-expandable-hoverable-expandable-toggle9").
										ID("table-expandable-hoverable-expandable-toggle9").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content9").
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
												Name("table-expandable-hoverable-checkrow9").
												Aria("labelledby", "table-expandable-hoverable-node9"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node9").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 9"),
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
												ID("table-expandable-hoverable-dropdown-kebab-9-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-9-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content9").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node10 table-expandable-hoverable-expandable-toggle10").
										ID("table-expandable-hoverable-expandable-toggle10").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content10").
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
												Name("table-expandable-hoverable-checkrow10").
												Aria("labelledby", "table-expandable-hoverable-node10"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node10").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 10"),
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
												ID("table-expandable-hoverable-dropdown-kebab-10-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-10-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content10").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected pf-m-expanded").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain pf-m-expanded").
										Aria("labelledby", "table-expandable-hoverable-node11 table-expandable-hoverable-expandable-toggle11").
										ID("table-expandable-hoverable-expandable-toggle11").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content11").
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
												Name("table-expandable-hoverable-checkrow11").
												Aria("labelledby", "table-expandable-hoverable-node11"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node11").
												Body(
													app.B().
														Body(
															app.Text("Expanded and selected"),
														),
												),
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
											app.Text("Link 11"),
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
												ID("table-expandable-hoverable-dropdown-kebab-11-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-11-button").
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
								ID("table-expandable-hoverable-content11").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node12 table-expandable-hoverable-expandable-toggle12").
										ID("table-expandable-hoverable-expandable-toggle12").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content12").
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
												Name("table-expandable-hoverable-checkrow12").
												Aria("labelledby", "table-expandable-hoverable-node12"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node12").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 12"),
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
												ID("table-expandable-hoverable-dropdown-kebab-12-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-12-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content12").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected pf-m-expanded").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain pf-m-expanded").
										Aria("labelledby", "table-expandable-hoverable-node13 table-expandable-hoverable-expandable-toggle13").
										ID("table-expandable-hoverable-expandable-toggle13").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content13").
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
												Name("table-expandable-hoverable-checkrow13").
												Aria("labelledby", "table-expandable-hoverable-node13"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node13").
												Body(
													app.B().
														Body(
															app.Text("Expanded and selected"),
														),
												),
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
											app.Text("Link 13"),
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
												ID("table-expandable-hoverable-dropdown-kebab-13-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-13-button").
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
								ID("table-expandable-hoverable-content13").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-selected pf-m-expanded").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain pf-m-expanded").
										Aria("labelledby", "table-expandable-hoverable-node15 table-expandable-hoverable-expandable-toggle15").
										ID("table-expandable-hoverable-expandable-toggle15").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content15").
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
												Name("table-expandable-hoverable-checkrow15").
												Aria("labelledby", "table-expandable-hoverable-node15"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node15").
												Body(
													app.B().
														Body(
															app.Text("Expanded and selected"),
														),
												),
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
											app.Text("Link 15"),
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
												ID("table-expandable-hoverable-dropdown-kebab-15-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-15-button").
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
								ID("table-expandable-hoverable-content15").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-expanded").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain pf-m-expanded").
										Aria("labelledby", "table-expandable-hoverable-node14 table-expandable-hoverable-expandable-toggle14").
										ID("table-expandable-hoverable-expandable-toggle14").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content14").
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
												Name("table-expandable-hoverable-checkrow14").
												Aria("labelledby", "table-expandable-hoverable-node14"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node14").
												Body(
													app.Text("Expanded and not selected"),
												),
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
											app.Text("Link 14"),
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
												ID("table-expandable-hoverable-dropdown-kebab-14-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-14-button").
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
								ID("table-expandable-hoverable-content14").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node16 table-expandable-hoverable-expandable-toggle16").
										ID("table-expandable-hoverable-expandable-toggle16").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content16").
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
												Name("table-expandable-hoverable-checkrow16").
												Aria("labelledby", "table-expandable-hoverable-node16"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node16").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 16"),
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
												ID("table-expandable-hoverable-dropdown-kebab-16-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-16-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content16").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable pf-m-expanded").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain pf-m-expanded").
										Aria("labelledby", "table-expandable-hoverable-node17 table-expandable-hoverable-expandable-toggle17").
										ID("table-expandable-hoverable-expandable-toggle17").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content17").
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
												Name("table-expandable-hoverable-checkrow17").
												Aria("labelledby", "table-expandable-hoverable-node17"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node17").
												Body(
													app.Text("Expanded and not selected"),
												),
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
											app.Text("Link 17"),
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
												ID("table-expandable-hoverable-dropdown-kebab-17-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-17-button").
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
								ID("table-expandable-hoverable-content17").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node18 table-expandable-hoverable-expandable-toggle18").
										ID("table-expandable-hoverable-expandable-toggle18").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content18").
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
												Name("table-expandable-hoverable-checkrow18").
												Aria("labelledby", "table-expandable-hoverable-node18"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node18").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 18"),
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
												ID("table-expandable-hoverable-dropdown-kebab-18-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-18-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content18").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
			app.TBody().
				Class("pf-m-hoverable").
				Aria("role", "rowgroup").
				TabIndex(0).
				Body(
					app.Tr().
						Class("pf-m-expanded").
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-c-table__toggle").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Aria("labelledby", "table-expandable-hoverable-node19 table-expandable-hoverable-expandable-toggle19").
										ID("table-expandable-hoverable-expandable-toggle19").
										Aria("label", "Details").
										Aria("controls", "table-expandable-hoverable-content19").
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
												Name("table-expandable-hoverable-checkrow19").
												Aria("labelledby", "table-expandable-hoverable-node19"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Div().
												ID("table-expandable-hoverable-node19").
												Body(
													app.Text("Hoverable"),
												),
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
											app.Text("Link 19"),
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
												ID("table-expandable-hoverable-dropdown-kebab-19-button").
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
												Aria("labelledby", "table-expandable-hoverable-dropdown-kebab-19-button").
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
								ColSpan(4).
								ID("table-expandable-hoverable-content19").
								Body(
									app.Div().
										Class("pf-c-table__expandable-row-content").
										Body(
											app.Text("Expandable content"),
										),
								),
							app.Td(),
						),
				),
		)
}
