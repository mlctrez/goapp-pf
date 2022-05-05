package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Modifiers struct {
	app.Compo
}

func (c *Modifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.H2().
				Class("Preview__section-title").
				Body(
					app.Text("Default fitting - example 1"),
				),
			app.Ul().
				Class("pf-c-data-list").
				Aria("role", "list").
				Aria("label", "Width modifier data list example 1").
				ID("data-list-default-fitting").
				Body(
					app.Li().
						Class("pf-c-data-list__item").
						Aria("labelledby", "data-list-default-fitting-item-1").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-default-fitting-item-1-checkbox").
														Aria("labelledby", "data-list-default-fitting-item-1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																ID("data-list-default-fitting-item-1").
																Body(
																	app.Text("default"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit."),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																Body(
																	app.Text("default"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
																),
														),
												),
										),
								),
						),
				),
			app.H2().
				Class("Preview__section-title").
				Body(
					app.Text("Flex modifiers - example 2"),
				),
			app.Ul().
				Class("pf-c-data-list").
				Aria("role", "list").
				Aria("label", "Width modifier data list example 2").
				ID("data-list-flex-modifiers").
				Body(
					app.Li().
						Class("pf-c-data-list__item").
						Aria("labelledby", "data-list-flex-modifiers-item-1").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-flex-modifiers-item-1-checkbox").
														Aria("labelledby", "data-list-flex-modifiers-item-1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-2").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																ID("data-list-flex-modifiers-item-1").
																Body(
																	app.Text(".pf-m-flex-2"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt."),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-4").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																Body(
																	app.Text(".pf-m-flex-4"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet."),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-action").
										Body(
											app.Div().
												Class("pf-c-data-list__action").
												Body(
													app.Div().
														Class("pf-c-dropdown").
														Body(
															app.Button().
																Class("pf-c-dropdown__toggle pf-m-plain").
																ID("data-list-flex-modifiers-item-1-dropdown-kebab-button").
																Aria("expanded", "false").
																Type("button").
																Aria("label", "Actions").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu pf-m-align-right").
																Aria("labelledby", "data-list-flex-modifiers-item-1-dropdown-kebab-button").
																Hidden(true).
																Body(
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item").
																				Href("#").
																				Body(
																					app.Text("Link"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Type("button").
																				Body(
																					app.Text("Action"),
																				),
																		),
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item pf-m-disabled").
																				Href("#").
																				Aria("disabled", true).
																				TabIndex(-1).
																				Body(
																					app.Text("Disabled link"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Type("button").
																				Disabled(true).
																				Body(
																					app.Text("Disabled action"),
																				),
																		),
																	app.Li().
																		Class("pf-c-divider").
																		Aria("role", "separator"),
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item").
																				Href("#").
																				Body(
																					app.Text("Separated link"),
																				),
																		),
																),
														),
												),
										),
								),
						),
				),
			app.H2().
				Class("Preview__section-title").
				Body(
					app.Text("Flex modifiers - example 3"),
				),
			app.Ul().
				Class("pf-c-data-list").
				Aria("role", "list").
				Aria("label", "Width modifier data list example 3").
				ID("data-list-flex-modifiers-2").
				Body(
					app.Li().
						Class("pf-c-data-list__item pf-m-expanded").
						Aria("labelledby", "data-list-flex-modifiers-2-item-1").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Div().
												Class("pf-c-data-list__toggle").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("labelledby", "data-list-flex-modifiers-2-toggle1 data-list-flex-modifiers-2-item1").
														ID("data-list-flex-modifiers-2-toggle1").
														Aria("label", "Toggle details for").
														Aria("expanded", true).
														Aria("controls", "data-list-flex-modifiers-2-content1").
														Body(
															app.Div().
																Class("pf-c-data-list__toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-flex-modifiers-2-item-1-checkbox").
														Aria("labelledby", "data-list-flex-modifiers-2-item-1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-5").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																ID("data-list-flex-modifiers-2-item-1").
																Body(
																	app.Text(".pf-m-flex-5"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit."),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-2").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																Body(
																	app.Text(".pf-m-flex-2"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit."),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-3").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																Body(
																	app.Text(".pf-m-flex-3"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit."),
																),
														),
												),
											app.Div().
												Class("pf-c-data-list__cell pf-m-flex-3").
												Body(
													app.Div().
														Class("Preview__placeholder").
														Body(
															app.B().
																Body(
																	app.Text(".pf-m-flex-3"),
																),
															app.P().
																Body(
																	app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit."),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-action").
										Body(
											app.Div().
												Class("pf-c-data-list__action").
												Body(
													app.Div().
														Class("pf-c-dropdown").
														Body(
															app.Button().
																Class("pf-c-dropdown__toggle pf-m-plain").
																ID("data-list-flex-modifiers-2-item-1-dropdown-kebab-button").
																Aria("expanded", "false").
																Type("button").
																Aria("label", "Actions").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu pf-m-align-right").
																Aria("labelledby", "data-list-flex-modifiers-2-item-1-dropdown-kebab-button").
																Hidden(true).
																Body(
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item").
																				Href("#").
																				Body(
																					app.Text("Link"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Type("button").
																				Body(
																					app.Text("Action"),
																				),
																		),
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item pf-m-disabled").
																				Href("#").
																				Aria("disabled", true).
																				TabIndex(-1).
																				Body(
																					app.Text("Disabled link"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Type("button").
																				Disabled(true).
																				Body(
																					app.Text("Disabled action"),
																				),
																		),
																	app.Li().
																		Class("pf-c-divider").
																		Aria("role", "separator"),
																	app.Li().
																		Body(
																			app.A().
																				Class("pf-c-dropdown__menu-item").
																				Href("#").
																				Body(
																					app.Text("Separated link"),
																				),
																		),
																),
														),
												),
										),
								),
							app.Section().
								Class("pf-c-data-list__expandable-content").
								ID("data-list-flex-modifiers-2-content1").
								Aria("label", "Primary content details").
								Body(
									app.Div().
										Class("pf-c-data-list__expandable-content-body").
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
										),
								),
						),
				),
		)
}
