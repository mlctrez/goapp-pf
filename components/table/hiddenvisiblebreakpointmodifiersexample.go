package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HiddenvisibleBreakpointModifiersExample struct {
	app.Compo
}

func (c *HiddenvisibleBreakpointModifiersExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-lg").
		Aria("role", "grid").
		Aria("label", "Table with hidden and visible modifiers example").
		ID("table-hidden-visible").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-hidden pf-m-visible-on-md pf-m-hidden-on-lg").
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
								Class("pf-m-hidden-on-md pf-m-visible-on-lg").
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
								Class("pf-m-hidden pf-m-visible-on-sm").
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
								Class("pf-m-hidden pf-m-visible-on-md pf-m-hidden-on-lg").
								Aria("role", "cell").
								DataSet("label", "Repository name").
								Body(
									app.Text("Visible only on md breakpoint"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Branches").
								Body(
									app.Text("10"),
								),
							app.Td().
								Class("pf-m-hidden-on-md pf-m-visible-on-lg").
								Aria("role", "cell").
								DataSet("label", "Pull requests").
								Body(
									app.Text("Hidden only on md breakpoint"),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Workspaces").
								Body(
									app.Text("5"),
								),
							app.Td().
								Class("pf-m-hidden pf-m-visible-on-sm").
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Text("Hidden on xs breakpoint"),
								),
						),
					app.Tr().
						Aria("role", "row").
						Body(
							app.Td().
								Class("pf-m-hidden pf-m-visible-on-md pf-m-hidden-on-lg").
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
								Class("pf-m-hidden-on-md pf-m-visible-on-lg").
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
								Class("pf-m-hidden pf-m-visible-on-sm").
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
								Class("pf-m-hidden pf-m-visible-on-md pf-m-hidden-on-lg").
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
								Class("pf-m-hidden-on-md pf-m-visible-on-lg").
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
								Class("pf-m-hidden pf-m-visible-on-sm").
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
								Class("pf-m-hidden pf-m-visible-on-md pf-m-hidden-on-lg").
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
								Class("pf-m-hidden-on-md pf-m-visible-on-lg").
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
								Class("pf-m-hidden pf-m-visible-on-sm").
								Aria("role", "cell").
								DataSet("label", "Last commit").
								Body(
									app.Text("2 days ago"),
								),
						),
				),
		)
}
