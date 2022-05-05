package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FavoritesSortableExample struct {
	app.Compo
}

func (c *FavoritesSortableExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a sortable with favorites table example").
		ID("table-favorites-sortable").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-c-table__sort pf-m-selected pf-m-favorite").
								Aria("role", "columnheader").
								Aria("sort", "none").
								Scope("col").
								Body(
									app.Button().
										Class("pf-c-table__button").
										Aria("label", "Favorite").
										Body(
											app.Div().
												Class("pf-c-table__button-content").
												Body(
													app.Span().
														Class("pf-c-table__text").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-table__sort-indicator").
														Body(
															app.I().
																Class("fas fa-long-arrow-alt-down"),
														),
												),
										),
								),
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
								Class("pf-c-table__favorite pf-m-favorited").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Starred").
										ID("table-favorites-sortable-favorite-button1").
										Aria("labelledby", "table-favorites-sortable-node1 table-favorites-sortable-favorite-button1").
										Body(
											app.I().
												Class("fas fa-star").
												Aria("hidden", true),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Div().
										Body(
											app.Span().
												ID("table-favorites-sortable-node1").
												Body(
													app.Text("Repository 1"),
												),
											app.Text(". This is a long title that will wrap to multiple lines. This is a long title that will wrap to multiple lines. This is a long title that will wrap to multiple lines. This is a long title that will wrap to multiple lines."),
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
								Class("pf-c-table__favorite pf-m-favorited").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Starred").
										ID("table-favorites-sortable-favorite-button3").
										Aria("labelledby", "table-favorites-sortable-node3 table-favorites-sortable-favorite-button3").
										Body(
											app.I().
												Class("fas fa-star").
												Aria("hidden", true),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Span().
										ID("table-favorites-sortable-node3").
										Body(
											app.Text("Repository 3"),
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
								Class("pf-c-table__favorite").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Not starred").
										ID("table-favorites-sortable-favorite-button2").
										Aria("labelledby", "table-favorites-sortable-node2 table-favorites-sortable-favorite-button2").
										Body(
											app.I().
												Class("fas fa-star").
												Aria("hidden", true),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Span().
										ID("table-favorites-sortable-node2").
										Body(
											app.Text("Repository 2"),
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
								Class("pf-c-table__favorite").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Not starred").
										ID("table-favorites-sortable-favorite-button4").
										Aria("labelledby", "table-favorites-sortable-node4 table-favorites-sortable-favorite-button4").
										Body(
											app.I().
												Class("fas fa-star").
												Aria("hidden", true),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Span().
										ID("table-favorites-sortable-node4").
										Body(
											app.Text("Repository 4"),
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
								Class("pf-c-table__favorite").
								Aria("role", "cell").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Not starred").
										ID("table-favorites-sortable-favorite-button5").
										Aria("labelledby", "table-favorites-sortable-node5 table-favorites-sortable-favorite-button5").
										Body(
											app.I().
												Class("fas fa-star").
												Aria("hidden", true),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Span().
										ID("table-favorites-sortable-node5").
										Body(
											app.Text("Repository 5"),
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
