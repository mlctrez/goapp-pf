package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StripedTableExample struct {
	app.Compo
}

func (c *StripedTableExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md pf-m-striped").
		Aria("role", "grid").
		Aria("label", "This is a striped table example").
		ID("table-striped").
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
									app.Text("Repository 3"),
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
									app.Text("Repository 4"),
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
