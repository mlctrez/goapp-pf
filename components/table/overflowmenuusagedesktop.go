package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type OverflowMenuUsageDesktop struct {
	app.Compo
}

func (c *OverflowMenuUsageDesktop) Render() app.UI {
	return app.Table().
		Class("pf-c-table").
		Aria("label", "This is a simple table example").
		ID("table-with-expanded-overflow-menu").
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
							app.Td(),
						),
				),
			app.TBody().
				Body(
					app.Tr().
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
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("table-with-expanded-overflow-menu-overflow-menu-1").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__content").
												Body(
													app.Div().
														Class("pf-c-overflow-menu__group pf-m-button-group").
														Body(
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-primary").
																		Type("button").
																		Body(
																			app.Text("Start"),
																		),
																),
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-secondary").
																		Type("button").
																		Body(
																			app.Text("Stop"),
																		),
																),
														),
												),
										),
								),
						),
					app.Tr().
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
								Class("pf-c-table__action").
								Aria("role", "cell").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("table-with-expanded-overflow-menu-overflow-menu-2").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__content").
												Body(
													app.Div().
														Class("pf-c-overflow-menu__group pf-m-button-group").
														Body(
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-primary").
																		Type("button").
																		Body(
																			app.Text("Start"),
																		),
																),
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-secondary").
																		Type("button").
																		Body(
																			app.Text("Stop"),
																		),
																),
														),
												),
										),
								),
						),
				),
		)
}
