package table

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TextControlUsingTheTableTextElementExample struct {
	app.Compo
}

func (c *TextControlUsingTheTableTextElementExample) Render() app.UI {
	return app.Table().
		Class("pf-c-table pf-m-grid-md").
		Aria("role", "grid").
		Aria("label", "This is a simple table example").
		ID("table-text-element-example").
		Body(
			app.Caption().
				Body(
					app.Text("This table contains"), app.Code().
						Body(
							app.Text(".pf-c-table__text"),
						),
					app.Text("examples. The"),
					app.Code().
						Body(
							app.Text(".pf-c-table__text"),
						),
					app.Text("element can be using alone or in a nested configuration."),
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
									app.Text("Selector/element"),
								),
							app.Th().
								Aria("role", "columnheader").
								Scope("col").
								Body(
									app.Text("Result"),
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
								Class("pf-m-fit-content").
								Aria("role", "rowheader").
								DataSet("label", "Element").
								Scope("row").
								Body(
									app.Div().
										Class("pf-c-table__text").
										Body(
											app.B().
												Body(
													app.Code().
														Body(
															app.Text("th.pf-m-truncate"),
														),
												),
										),
								),
							app.Td().
								Class("pf-m-truncate").
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Span().
										Class("pf-c-table__text").
										Body(
											app.Text("This table cell contains a single"), app.Code().
												Body(
													app.Text("`.pf-c-table__text`"),
												),
											app.Text("wrapper with the parent table cell applying"),
											app.Code().
												Body(
													app.Text("`.pf-m-truncate`"),
												),
											app.Text(". The child"),
											app.Code().
												Body(
													app.Text("`.pf-c-table__text`"),
												),
											app.Text("element will inherit the modifier settings and apply to the grid layout."),
										),
								),
						),
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "rowheader").
								DataSet("label", "Element").
								Scope("row").
								Body(
									app.Div().
										Class("pf-c-table__text").
										Body(
											app.B().
												Body(
													app.Code().
														Body(
															app.Text(".pf-l-stack"),
														),
												),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Div().
										Class("pf-l-stack pf-m-gutter").
										Body(
											app.Div().
												Class("pf-l-stack__item").
												Body(
													app.Div().
														Class("pf-c-table__text").
														Body(
															app.Text("Because"), app.Code().
																Body(
																	app.Text(".pf-m-grid"),
																),
															app.Text("applies a grid layout to"),
															app.Code().
																Body(
																	app.Text(".pf-c-table"),
																),
															app.Text(", child elements will stack in the grid layout. To prevent this, wrap multiple elements with a div or use a PatternFly layout."),
														),
												),
											app.Div().
												Class("pf-l-stack__item").
												Body(
													app.P().
														Class("pf-c-table__text").
														Body(
															app.Text("The"), app.B().
																Body(
																	app.Code().
																		Body(
																			app.Text(".pf-c-table__text"),
																		),
																	app.Text("element"),
																),
															app.Text("can additionally be nested, like in this example. The next"),
															app.Code().
																Body(
																	app.Text(".pf-c-table__text"),
																),
															app.Text("element has a very long url whose width needs be constrained."),
														),
												),
											app.Div().
												Class("pf-l-stack__item").
												Body(
													app.A().
														Class("pf-c-table__text pf-m-truncate").
														Href("#").
														Body(
															app.Text("http://truncatemodifierappliedtoaverylongurlthatwillforcethetabletobreakluckilywehavethepfctabletextelement.com"),
														),
												),
											app.Div().
												Class("pf-l-stack__item").
												Body(
													app.P().
														Class("pf-c-table__text").
														Body(
															app.Text("This"), app.B().
																Body(
																	app.Code().
																		Body(
																			app.Text(".pf-c-table__text"),
																		),
																	app.Text("element"),
																),
															app.Text("applies its own built in grid layout"),
															app.B().
																Body(
																	app.Code().
																		Body(
																			app.Text(".pf-m-stack"),
																		),
																),
															app.Text("as well as a gutter"),
															app.B().
																Body(
																	app.Code().
																		Body(
																			app.Text(".pf-m-gutter"),
																		),
																),
															app.Text("."),
														),
												),
										),
								),
						),
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "rowheader").
								DataSet("label", "Element").
								Scope("row").
								Body(
									app.Div().
										Class("pf-c-table__text").
										Body(
											app.B().
												Body(
													app.Code().
														Body(
															app.Text(".pf-l-flex"),
														),
												),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Div().
										Class("pf-l-flex pf-m-column pf-m-row-on-xl").
										Body(
											app.Div().
												Class("pf-l-flex__item pf-m-flex-1").
												Body(
													app.P().
														Class("pf-c-table__text").
														Body(
															app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt."),
														),
												),
											app.Div().
												Class("pf-l-flex__item pf-m-flex-1").
												Body(
													app.A().
														Class("pf-c-table__text pf-m-break-word").
														Href("#").
														Body(
															app.Text("http://breakwordmodifierappliedtoaverylongurlthatwillforcethetabletobreakluckilywehavethepfctabletextelement.com"),
														),
												),
										),
								),
						),
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "rowheader").
								DataSet("label", "Element").
								Scope("row").
								Body(
									app.Div().
										Class("pf-c-table__text").
										Body(
											app.B().
												Body(
													app.Code().
														Body(
															app.Text(".pf-l-flex"),
														),
												),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Div().
										Class("pf-l-flex pf-m-column").
										Body(
											app.Div().
												Class("pf-l-flex").
												Body(
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.I().
																Class("fas fa-code-branch").
																Aria("hidden", true),
															app.Text("5"),
														),
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.I().
																Class("fas fa-code").
																Aria("hidden", true),
															app.Text("9"),
														),
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.I().
																Class("fas fa-cube").
																Aria("hidden", true),
															app.Text("2"),
														),
													app.Div().
														Class("pf-l-flex__item").
														Body(
															app.I().
																Class("fas fa-check-circle").
																Aria("hidden", true),
															app.Text("11"),
														),
												),
											app.Div().
												Class("pf-l-flex__item").
												Body(
													app.P().
														Class("pf-c-table__text").
														Body(
															app.Text("This is paragraph that we want to wrap. It doesn't need a modifier and has no extra long strings. Any modifier available for the flex layout can be used here."),
														),
												),
											app.Div().
												Class("pf-l-flex__item").
												Body(
													app.A().
														Class("pf-c-table__text pf-m-break-word").
														Href("#").
														Body(
															app.Text("http://breakwordmodifierappliedtoaverylongurlthatwillforcethetabletobreakluckilywehavethepfctabletextelement.com"),
														),
												),
										),
								),
						),
					app.Tr().
						Aria("role", "row").
						Body(
							app.Th().
								Class("pf-m-fit-content").
								Aria("role", "rowheader").
								DataSet("label", "Element").
								Scope("row").
								Body(
									app.Div().
										Class("pf-c-table__text").
										Body(
											app.B().
												Body(
													app.Code().
														Body(
															app.Text(".pf-l-grid"),
														),
												),
										),
								),
							app.Td().
								Aria("role", "cell").
								DataSet("label", "Truncating text").
								Body(
									app.Div().
										Class("pf-l-grid pf-m-gutter").
										Body(
											app.Div().
												Class("pf-l-grid__item pf-m-6-col pf-m-3-col-on-md").
												Body(
													app.Text("Item 1"),
												),
											app.Div().
												Class("pf-l-grid__item pf-m-6-col pf-m-3-col-on-md").
												Body(
													app.Text("Item 2"),
												),
											app.Div().
												Class("pf-l-grid__item pf-m-6-col pf-m-3-col-on-md").
												Body(
													app.Text("Item 3"),
												),
											app.Div().
												Class("pf-l-grid__item pf-m-6-col pf-m-3-col-on-md").
												Body(
													app.Text("Item 4"),
												),
											app.Div().
												Class("pf-l-grid__item").
												Body(
													app.P().
														Class("pf-c-table__text").
														Body(
															app.Text("This is paragraph that we want to wrap. It doesn't need a modifier and has no extra long strings. Any modifier available for the flex layout can be used here."),
														),
												),
											app.Div().
												Class("pf-l-grid__item").
												Body(
													app.A().
														Class("pf-c-table__text pf-m-truncate").
														Href("#").
														Body(
															app.Text("http://breakwordmodifierappliedtoaverylongurlthatwillforcethetabletobreakluckilywehavethepfctabletextelement.com"),
														),
												),
										),
								),
						),
				),
		)
}
