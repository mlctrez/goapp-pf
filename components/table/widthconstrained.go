package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WidthConstrained struct {
	app.Compo
}

func (c *WidthConstrained) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a simple table example").
		ID("width-constrained-example").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-width-40").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Width 40"),
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
								Class("pf-m-fit-content").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Fit content th"),
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
								Class("").
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("Since this is a long string of text and the other cells contain short strings (narrower than their table header), we'll need to control width this table header's width. Let's set width to 40%."),
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
				Aria("role", "rowgroup").
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-truncate").
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("This string will truncate in table mode only. Since this is a long string of text and the other cells contain short strings (narrower than their table header), we'll need to control width this table header's width. Let's set width to 40%."),
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
