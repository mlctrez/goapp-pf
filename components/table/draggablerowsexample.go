package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DraggableRowsExample struct {
	app.Compo
}

func (c *DraggableRowsExample) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				ID("table-draggable-rows-help").
				Body(
					app.Text("Activate the reorder button and use the arrow keys to reorder the list or use your mouse to drag/reorder. Press escape to cancel the reordering."),
				),
			app.Table().
				Class("pf-c-table pf-m-grid-md").
				Aria("role", "grid").
				Aria("label", "This is a table with draggable rows example").
				ID("table-draggable-rows").
				Body(
					app.Caption().
						Body(
							app.Text("This is the table caption"),
						),
					app.THead().
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Td(),
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
										Class("pf-c-table__draggable").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Disabled(true).
												Aria("pressed", "false").
												Aria("label", "Reorder").
												Aria("describedby", "table-draggable-rows-help").
												ID("table-draggable-rows-row-1-draggable-button").
												Aria("labelledby", "table-draggable-rows-row-1-draggable-button table-draggable-rows-row-1-node").
												Body(
													app.I().
														Class("fas fa-grip-vertical").
														Aria("hidden", true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Span().
												ID("table-draggable-rows-row-1-node").
												Body(
													app.Text("Draggable icon disabled"),
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
										Class("pf-c-table__draggable").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("pressed", "false").
												Aria("label", "Reorder").
												Aria("describedby", "table-draggable-rows-help").
												ID("table-draggable-rows-row-2-draggable-button").
												Aria("labelledby", "table-draggable-rows-row-2-draggable-button table-draggable-rows-row-2-node").
												Body(
													app.I().
														Class("fas fa-grip-vertical").
														Aria("hidden", true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Span().
												ID("table-draggable-rows-row-2-node").
												Body(
													app.Text("Table cell"),
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
								Class("pf-m-ghost-row").
								Aria("role", "row").
								Body(
									app.Td().
										Class("pf-c-table__draggable").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Disabled(true).
												Aria("pressed", "false").
												Aria("label", "Reorder").
												Aria("describedby", "table-draggable-rows-help").
												ID("table-draggable-rows-row-3-draggable-button").
												Aria("labelledby", "table-draggable-rows-row-3-draggable-button table-draggable-rows-row-3-node").
												Body(
													app.I().
														Class("fas fa-grip-vertical").
														Aria("hidden", true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Span().
												ID("table-draggable-rows-row-3-node").
												Body(
													app.Text("Ghost row"),
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
										Class("pf-c-table__draggable").
										Aria("role", "cell").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("pressed", "false").
												Aria("label", "Reorder").
												Aria("describedby", "table-draggable-rows-help").
												ID("table-draggable-rows-row-4-draggable-button").
												Aria("labelledby", "table-draggable-rows-row-4-draggable-button table-draggable-rows-row-4-node").
												Body(
													app.I().
														Class("fas fa-grip-vertical").
														Aria("hidden", true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Span().
												ID("table-draggable-rows-row-4-node").
												Body(
													app.Text("Table cell"),
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
				),
			app.Div().
				Class("pf-screen-reader").
				Aria("live", "assertive").
				Body(
					app.Text("This is the aria-live section that provides real-time feedback to the user."),
				),
		)
}
