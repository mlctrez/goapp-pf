package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TableWithLongStringsInTheContent struct {
	app.Compo
}

func (c *TableWithLongStringsInTheContent) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tooltip pf-m-top").
				Aria("role", "tooltip").
				Body(
					app.Div().
						Class("pf-c-tooltip__arrow"),
					app.Div().
						Class("pf-c-tooltip__content").
						ID("tooltip-top-content").
						Body(
							app.Text("Pull Requests"),
						),
				),
			app.Table().
				Class("pf-c-table").
				Aria("label", "This is a simple table example").
				ID("th-truncation-example").
				Body(
					app.THead().
						Body(
							app.Tr().
								Body(
									app.Th().
										Scope("col").
										Body(
											app.Text("Repositories"),
										),
									app.Th().
										Scope("col").
										Body(
											app.Text("Branches"),
										),
									app.Th().
										Scope("col").
										Body(
											app.Text("Pull requests"),
										),
									app.Th().
										Scope("col").
										Body(
											app.Text("Workspaces"),
										),
									app.Th().
										Scope("col").
										Body(
											app.Text("Last commit"),
										),
								),
						),
					app.TBody().
						Body(
							app.Tr().
								Body(
									app.Td().
										Class("").
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Text("Long lines of text will shrink adjacent column widths."),
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
								Body(
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Repository name").
										Body(
											app.Text("This example is not responsive. Adjacent"), app.Code().
												Body(
													app.Text("tbody"),
												),
											app.Text("cells will shrink as a result of this text being a longer string and adjacent text being shorter in length. Truncation can be overridden in"),
											app.Code().
												Body(
													app.Text("th"),
												),
											app.Text("cells with the addition of"),
											app.Code().
												Body(
													app.Text(".pf-m-wrap"),
												),
											app.Text(","),
											app.Code().
												Body(
													app.Text(".pf-m-nowrap"),
												),
											app.Text("or"),
											app.Code().
												Body(
													app.Text(".pf-m-fit-content"),
												),
											app.Text("."),
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
		)
}
