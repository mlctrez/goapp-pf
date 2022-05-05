package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WidthModifiersExamples struct {
	app.Compo
}

func (c *WidthModifiersExamples) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a width modifier expandable").
		ID("table-width-modifiers").
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
												Name("table-width-modifiers-checkrowthead").
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
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Workspaces"),
								),
							app.Th().
								Class("pf-m-fit-content").
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
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-width-modifiers-checkrow1").
												Aria("labelledby", "table-width-modifiers-node1"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-width-modifiers-node1").
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
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-width-modifiers-checkrow2").
												Aria("labelledby", "table-width-modifiers-node2"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-width-modifiers-node2").
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
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-width-modifiers-checkrow3").
												Aria("labelledby", "table-width-modifiers-node3"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-width-modifiers-node3").
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
								Class("pf-c-table__check").
								Aria("role", "cell").
								Body(
									app.Label().
										Body(
											app.Input().
												Type("checkbox").
												Name("table-width-modifiers-checkrow4").
												Aria("labelledby", "table-width-modifiers-node4"),
										),
								),
							app.Th().
								Aria("role", "columnheader").
								DataSet("label", "Data label name").
								Body(
									app.Div().
										ID("table-width-modifiers-node4").
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
		)
}
