package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NestedColumnHeaders struct {
	app.Compo
}

func (c *NestedColumnHeaders) Render() app.UI {
	return app.Div().
		Class("pf-c-scroll-inner-wrapper").
		Body(
			app.Table().
				Class("pf-c-table").
				Aria("role", "grid").
				Aria("label", "This is a nested column header table example").
				ID("table-nested-column-headers-example").
				Body(
					app.THead().
						Class("pf-m-nested-column-header").
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Th().
										Class("pf-m-border-right").
										Aria("role", "columnheader").
										Scope("col").
										ColSpan(3).
										Body(
											app.Text("Pods"),
										),
									app.Th().
										Class("pf-m-border-right").
										Aria("role", "columnheader").
										Scope("col").
										ColSpan(2).
										Body(
											app.Text("Ports"),
										),
									app.Th().
										Class("pf-m-border-right pf-m-fit-content pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
										Rowspan(2).
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
																	app.Text("Protocol"),
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
										Class("pf-m-border-right pf-m-fit-content pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
										Rowspan(2).
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
																	app.Text("Flow rate"),
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
										Class("pf-m-border-right pf-m-fit-content pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
										Rowspan(2).
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
																	app.Text("Traffic"),
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
										Class("pf-m-fit-content pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
										Rowspan(2).
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
																	app.Text("Packets"),
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
								),
							app.Tr().
								Aria("role", "row").
								Body(
									app.Th().
										Class("pf-c-table__subhead pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
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
																	app.Text("Source"),
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
										Class("pf-c-table__subhead pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
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
																	app.Text("Destination"),
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
										Class("pf-c-table__subhead pf-m-fit-content pf-m-border-right pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
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
																	app.Text("Date & Time"),
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
										Class("pf-c-table__subhead pf-m-fit-content pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
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
																	app.Text("Source"),
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
										Class("pf-c-table__subhead pf-m-fit-content pf-m-border-right pf-c-table__sort").
										Aria("role", "columnheader").
										Aria("sort", "none").
										Scope("col").
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
																	app.Text("Destination"),
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
										DataSet("label", "Source").
										Body(
											app.Div().
												Class("pf-l-flex pf-m-nowrap").
												Body(
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.Span().
																Class("pf-c-label pf-m-cyan").
																Body(
																	app.Span().
																		Class("pf-c-label__content").
																		Body(
																			app.Text("P"),
																		),
																),
														),
													app.Div().
														Class("pf-l-flex__item pf-m-flex-1").
														Body(
															app.Span().
																Class("pf-c-table__text pf-m-truncate").
																Body(
																	app.A().
																		Href("#").
																		Body(
																			app.Text("api-pod-source-name"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Destination").
										Body(
											app.Div().
												Class("pf-l-flex pf-m-nowrap").
												Body(
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.Span().
																Class("pf-c-label pf-m-cyan").
																Body(
																	app.Span().
																		Class("pf-c-label__content").
																		Body(
																			app.Text("P"),
																		),
																),
														),
													app.Div().
														Class("pf-l-flex__item pf-m-flex-1").
														Body(
															app.Span().
																Class("pf-c-table__text pf-m-truncate").
																Body(
																	app.A().
																		Href("#").
																		Body(
																			app.Text("api-pod-destination-name"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Date & time").
										Body(
											app.Div().
												Class("pf-l-stack").
												Body(
													app.Span().
														Body(
															app.Text("June 22, 2021"),
														),
													app.Span().
														Class("pf-u-color-200").
														Body(
															app.Text("3:58:24 PM"),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Source").
										Body(
											app.Div().
												Class("pf-l-stack").
												Body(
													app.Span().
														Body(
															app.Text("443"),
														),
													app.Span().
														Class("pf-u-color-200").
														Body(
															app.Text("(HTTPS)"),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Destination").
										Body(
											app.Div().
												Class("pf-l-stack").
												Body(
													app.Span().
														Body(
															app.Text("24"),
														),
													app.Span().
														Class("pf-u-color-200").
														Body(
															app.Text("(smtp)"),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Protocol").
										Body(
											app.Text("TCP"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Flow rate").
										Body(
											app.Text("1.9 Kbps"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Traffic").
										Body(
											app.Text("2.1 KB"),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Packets").
										Body(
											app.Text("3"),
										),
								),
						),
				),
		)
}
