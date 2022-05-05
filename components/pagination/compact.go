package pagination

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Div().
		Class("pf-c-pagination pf-m-compact").
		Body(
			app.Div().
				Class("pf-c-pagination__total-items").
				Body(
					app.B().
						Body(
							app.Text("1 - 10"),
						),
					app.Text("of"),
					app.B().
						Body(
							app.Text("36"),
						),
				),
			app.Div().
				Class("pf-c-options-menu").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
						Type("button").
						ID("pagination-options-menu-compact-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.B().
										Body(
											app.Text("1 - 10"),
										),
									app.Text("of"),
									app.B().
										Body(
											app.Text("36"),
										),
								),
							app.Div().
								Class("pf-c-options-menu__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "pagination-options-menu-compact-example-toggle").
						Hidden(true).
						Body(
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("5 per page"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("10 per page"), app.Div().
												Class("pf-c-options-menu__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("20 per page"),
										),
								),
						),
				),
			app.Nav().
				Class("pf-c-pagination__nav").
				Aria("label", "Pagination").
				Body(
					app.Div().
						Class("pf-c-pagination__nav-control pf-m-prev").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Disabled(true).
								Aria("label", "Go to previous page").
								Body(
									app.I().
										Class("fas fa-angle-left").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-pagination__nav-control pf-m-next").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Go to next page").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
		)
}
