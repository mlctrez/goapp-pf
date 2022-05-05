package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupTypes struct {
	app.Compo
}

func (c *GroupTypes) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-group-types-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group pf-m-filter-group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Div().
												Class("pf-c-select").
												Body(
													app.Span().
														ID("toolbar-group-types-example-select-checkbox-filter1-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-group-types-example-select-checkbox-filter1-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter1-label toolbar-group-types-example-select-checkbox-filter1-toggle").
														Body(
															app.Div().
																Class("pf-c-select__toggle-wrapper").
																Body(
																	app.Span().
																		Class("pf-c-select__toggle-text").
																		Body(
																			app.Text("Filter 1"),
																		),
																),
															app.Span().
																Class("pf-c-select__toggle-arrow").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
													app.Ul().
														Class("pf-c-select__menu").
														Aria("role", "listbox").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter1-label").
														Hidden(true).
														Body(
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Running"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item pf-m-selected").
																		Aria("role", "option").
																		Aria("selected", true).
																		Body(
																			app.Text("Stopped"), app.Span().
																				Class("pf-c-select__menu-item-icon").
																				Body(
																					app.I().
																						Class("fas fa-check").
																						Aria("hidden", true),
																				),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Down"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Degraded"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Needs maintenance"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Div().
												Class("pf-c-select").
												Body(
													app.Span().
														ID("toolbar-group-types-example-select-checkbox-filter2-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-group-types-example-select-checkbox-filter2-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter2-label toolbar-group-types-example-select-checkbox-filter2-toggle").
														Body(
															app.Div().
																Class("pf-c-select__toggle-wrapper").
																Body(
																	app.Span().
																		Class("pf-c-select__toggle-text").
																		Body(
																			app.Text("Filter 2"),
																		),
																),
															app.Span().
																Class("pf-c-select__toggle-arrow").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
													app.Ul().
														Class("pf-c-select__menu").
														Aria("role", "listbox").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter2-label").
														Hidden(true).
														Body(
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Running"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item pf-m-selected").
																		Aria("role", "option").
																		Aria("selected", true).
																		Body(
																			app.Text("Stopped"), app.Span().
																				Class("pf-c-select__menu-item-icon").
																				Body(
																					app.I().
																						Class("fas fa-check").
																						Aria("hidden", true),
																				),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Down"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Degraded"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Needs maintenance"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Div().
												Class("pf-c-select").
												Body(
													app.Span().
														ID("toolbar-group-types-example-select-checkbox-filter3-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-group-types-example-select-checkbox-filter3-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter3-label toolbar-group-types-example-select-checkbox-filter3-toggle").
														Body(
															app.Div().
																Class("pf-c-select__toggle-wrapper").
																Body(
																	app.Span().
																		Class("pf-c-select__toggle-text").
																		Body(
																			app.Text("Filter 3"),
																		),
																),
															app.Span().
																Class("pf-c-select__toggle-arrow").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
													app.Ul().
														Class("pf-c-select__menu").
														Aria("role", "listbox").
														Aria("labelledby", "toolbar-group-types-example-select-checkbox-filter3-label").
														Hidden(true).
														Body(
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Running"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item pf-m-selected").
																		Aria("role", "option").
																		Aria("selected", true).
																		Body(
																			app.Text("Stopped"), app.Span().
																				Class("pf-c-select__menu-item-icon").
																				Body(
																					app.I().
																						Class("fas fa-check").
																						Aria("hidden", true),
																				),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Down"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Degraded"),
																		),
																),
															app.Li().
																Aria("role", "presentation").
																Body(
																	app.Button().
																		Class("pf-c-select__menu-item").
																		Aria("role", "option").
																		Body(
																			app.Text("Needs maintenance"),
																		),
																),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__group pf-m-icon-button-group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Edit").
												Body(
													app.I().
														Class("fas fa-edit").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Clone").
												Body(
													app.I().
														Class("fas fa-clone").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Sync").
												Body(
													app.I().
														Class("fas fa-sync").
														Aria("hidden", true),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__group pf-m-button-group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-primary").
												Type("button").
												Body(
													app.Text("Action"),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-secondary").
												Type("button").
												Body(
													app.Text("Secondary"),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-tertiary").
												Type("button").
												Body(
													app.Text("Tertiary"),
												),
										),
								),
						),
				),
		)
}
