package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MultipleOptions struct {
	app.Compo
}

func (c *MultipleOptions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-options-menu").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle").
						Type("button").
						ID("options-menu-multiple-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.Text("Sort by"),
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
						Aria("labelledby", "options-menu-multiple-example-toggle").
						Hidden(true).
						Body(
							app.Li().
								Body(
									app.Ul().
										Aria("label", "Sort by").
										Body(
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Name"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Date"), app.Div().
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
														Disabled(true).
														Body(
															app.Text("Disabled"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Size"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-divider").
								Aria("role", "separator"),
							app.Li().
								Body(
									app.Ul().
										Aria("label", "Sort direction").
										Body(
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Ascending"), app.Div().
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
															app.Text("Descending"),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-options-menu pf-m-expanded").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle").
						Type("button").
						ID("options-menu-multiple-expanded-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", true).
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.Text("Sort by"),
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
						Aria("labelledby", "options-menu-multiple-expanded-example-toggle").
						Body(
							app.Li().
								Body(
									app.Ul().
										Aria("label", "Sort by").
										Body(
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Name"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Date"), app.Div().
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
														Disabled(true).
														Body(
															app.Text("Disabled"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Size"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-divider").
								Aria("role", "separator"),
							app.Li().
								Body(
									app.Ul().
										Aria("label", "Sort direction").
										Body(
											app.Li().
												Body(
													app.Button().
														Class("pf-c-options-menu__menu-item").
														Type("button").
														Body(
															app.Text("Ascending"), app.Div().
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
															app.Text("Descending"),
														),
												),
										),
								),
						),
				),
		)
}
