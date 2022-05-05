package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StickyColumnsAndHeader struct {
	app.Compo
}

func (c *StickyColumnsAndHeader) Render() app.UI {
	return app.Div().
		Class("pf-c-scroll-outer-wrapper").
		Body(
			app.Div().
				Class("pf-c-scroll-inner-wrapper").
				Body(
					app.Table().
						Class("pf-c-table pf-m-sticky-header").
						Aria("role", "grid").
						Aria("label", "This is a scrollable table").
						ID("sticky-header-columns-example").
						Body(
							app.THead().
								Body(
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sort pf-c-table__sticky-column").
												Aria("role", "columnheader").
												Aria("sort", "none").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Button().
														Class("pf-c-table__button").
														Body(
															app.Div().
																Class("pf-c-table__button-content").
																Body(
																	app.Span().
																		Class("pf-c-table__text").
																		Body(
																			app.Text("Fact"),
																		),
																	app.Span().
																		Class("pf-c-table__sort-indicator").
																		Body(
																			app.I().
																				Class("fas fa-arrows-alt-v"),
																		),
																),
														),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sort pf-c-table__sticky-column").
												Aria("role", "columnheader").
												Aria("sort", "none").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Button().
														Class("pf-c-table__button").
														Body(
															app.Div().
																Class("pf-c-table__button-content").
																Body(
																	app.Span().
																		Class("pf-c-table__text").
																		Body(
																			app.Text("State"),
																		),
																	app.Span().
																		Class("pf-c-table__sort-indicator").
																		Body(
																			app.I().
																				Class("fas fa-arrows-alt-v"),
																		),
																),
														),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 3"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 4"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 5"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 6"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 7"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 8"),
												),
											app.Th().
												Class("pf-m-truncate").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Body(
													app.Text("Header 9"),
												),
										),
								),
							app.TBody().
								Aria("role", "rowgroup").
								Body(
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 1"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 1"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 1-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 2"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 2"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 2-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 3"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 3-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 4"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 4-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 5"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 5-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 6"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 6-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 7"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 7-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 8"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 8-9"),
												),
										),
									app.Tr().
										Aria("role", "row").
										Body(
											app.Th().
												Class("pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("--pf-c-table__sticky-column--MinWidth", " 100px;").
												Body(
													app.Text("Fact 9"),
												),
											app.Th().
												Class("pf-m-border-right pf-m-truncate pf-c-table__sticky-column").
												Aria("role", "columnheader").
												DataSet("label", "Example th").
												Scope("col").
												Style("", "").
												Body(
													app.Text("State 9"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-3"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-4"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-5"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-6"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-7"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-8"),
												),
											app.Td().
												Class("pf-m-nowrap").
												Aria("role", "cell").
												DataSet("label", "Example td").
												Body(
													app.Text("Test cell 9-9"),
												),
										),
								),
						),
				),
		)
}
