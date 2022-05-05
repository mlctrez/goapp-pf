package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StripedMultipleTbodyExample struct {
	app.Compo
}

func (c *StripedMultipleTbodyExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a striped tbody example").
		ID("table-striped-tbody").
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
				Class("pf-m-striped").
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("tbody 1 - Repository 1"), app.Br(),
									app.Small().
										Body(
											app.Text("(odd rows striped)"),
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
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("tbody 1 - Repository 2"),
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
									app.Text("tbody 1 - Repository 3"), app.Br(),
									app.Small().
										Body(
											app.Text("(odd rows striped)"),
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
			app.TBody().
				Class("pf-m-striped-even").
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("tbody 2 - Repository 4"),
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
									app.Text("tbody 2 - Repository 5"), app.Br(),
									app.Small().
										Body(
											app.Text("(even rows striped using `.pf-m-striped-even`)"),
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
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("tbody 2 - Repository 6"),
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
									app.Text("tbody 2 - Repository 7"), app.Br(),
									app.Small().
										Body(
											app.Text("(even rows striped using `.pf-m-striped-even`)"),
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
