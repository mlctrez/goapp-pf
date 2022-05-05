package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TextControlExample struct {
	app.Compo
}

func (c *TextControlExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-lg").
		Aria("role", "grid").
		Aria("label", "This is a simple table example").
		ID("modifiers-without-text-wrapper-example").
		Body(
			app.THead().
				Body(
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-width-20").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Truncate (width 20%)"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Break word"),
								),
							app.Th().
								Class("pf-m-wrap").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Wrapping table header text. This"), app.Code().
										Body(
											app.Text("th"),
										),
									app.Text("text will wrap instead of truncate."),
								),
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Fit content"),
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
								Class("pf-m-truncate").
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Text("This text will truncate instead of wrap in table layout and wrap gracefully in grid layout."),
								),
							app.Td().
								Class("pf-m-break-word").
								Aria("role", "cell").
								DataSet("label", "Break word").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Wrapping").
								Body(
									app.P().
										Body(
											app.Text("By default,"), app.Code().
												Body(
													app.Text("thead"),
												),
											app.Text("cells will truncate and"),
											app.Code().
												Body(
													app.Text("tbody"),
												),
											app.Text("cells will wrap. Use"),
											app.Code().
												Body(
													app.Text(".pf-m-wrap"),
												),
											app.Text("on a"),
											app.Code().
												Body(
													app.Text("th"),
												),
											app.Text("to change its behavior."),
										),
								),
							app.Td().
								Class("").
								Aria("role", "cell").
								DataSet("label", "Fit content").
								Body(
									app.Text("This cell's content will adjust itself to the parent th width. This modifier only affects table layouts."),
								),
							app.Td().
								Class("pf-m-nowrap").
								Aria("role", "cell").
								DataSet("label", "No wrap").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("No wrap"),
										),
								),
						),
				),
		)
}
