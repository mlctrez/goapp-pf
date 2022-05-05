package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableNested struct {
	app.Compo
}

func (c *ExpandableNested) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Expandable nested data list example").
		ID("data-list-expandable-nested").
		Body(
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded").
				Aria("labelledby", "data-list-expandable-nested-item-1").
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
												Aria("labelledby", "data-list-expandable-nested-toggle1 data-list-expandable-nested-item1").
												ID("data-list-expandable-nested-toggle1").
												Aria("label", "Toggle details for").
												Aria("expanded", true).
												Aria("controls", "data-list-expandable-nested-content1").
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
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell pf-m-icon").
										Body(
											app.I().
												Class("fas fa-code-branch").
												Aria("hidden", true),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-expandable-nested-item-1").
												Body(
													app.Text("Primary content"),
												),
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
												),
											app.A().
												Href("#").
												Body(
													app.Text("link"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
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
														ID("data-list-expandable-nested-item-1-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-expandable-nested-item-1-dropdown-kebab-button").
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
						ID("data-list-expandable-nested-content1").
						Aria("label", "Primary content details").
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body").
								Body(
									app.Ul().
										Class("pf-c-data-list").
										Aria("role", "list").
										Aria("label", "Expandable nested nested data list example").
										ID("data-list-expandable-nested-nested").
										Body(
											app.Li().
												Class("pf-c-data-list__item pf-m-expanded").
												Aria("labelledby", "data-list-expandable-nested-nested-item-1").
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
																				Aria("labelledby", "data-list-expandable-nested-nested-toggle1 data-list-expandable-nested-nested-item1").
																				ID("data-list-expandable-nested-nested-toggle1").
																				Aria("label", "Toggle details for").
																				Aria("expanded", true).
																				Aria("controls", "data-list-expandable-nested-nested-content1").
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
																),
															app.Div().
																Class("pf-c-data-list__item-content").
																Body(
																	app.Div().
																		Class("pf-c-data-list__cell").
																		Body(
																			app.Span().
																				ID("data-list-expandable-nested-nested-item-1").
																				Body(
																					app.Text("Nested row 1"),
																				),
																		),
																),
														),
													app.Section().
														Class("pf-c-data-list__expandable-content").
														ID("data-list-expandable-nested-nested-content1").
														Aria("label", "Nested row 1 details").
														Body(
															app.Div().
																Class("pf-c-data-list__expandable-content-body").
																Body(
																	app.Text("Nested row 1 expanded content."),
																),
														),
												),
											app.Li().
												Class("pf-c-data-list__item").
												Aria("labelledby", "data-list-expandable-nested-nested-item-2").
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
																				Aria("labelledby", "data-list-expandable-nested-nested-toggle2 data-list-expandable-nested-nested-item2").
																				ID("data-list-expandable-nested-nested-toggle2").
																				Aria("label", "Toggle details for").
																				Aria("expanded", "false").
																				Aria("controls", "data-list-expandable-nested-nested-content2").
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
																),
															app.Div().
																Class("pf-c-data-list__item-content").
																Body(
																	app.Div().
																		Class("pf-c-data-list__cell").
																		Body(
																			app.Div().
																				ID("data-list-expandable-nested-nested-item-2").
																				Body(
																					app.Text("Nested row 2"),
																				),
																		),
																),
														),
													app.Section().
														Class("pf-c-data-list__expandable-content").
														ID("data-list-expandable-nested-nested-content2").
														Aria("label", "Nested row 2 details").
														Hidden(true).
														Body(
															app.Div().
																Class("pf-c-data-list__expandable-content-body").
																Body(
																	app.Text("Nested row 2 expanded content."),
																),
														),
												),
											app.Li().
												Class("pf-c-data-list__item pf-m-expanded").
												Aria("labelledby", "data-list-expandable-nested-nested-item-3").
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
																				Aria("labelledby", "data-list-expandable-nested-nested-toggle3 data-list-expandable-nested-nested-item3").
																				ID("data-list-expandable-nested-nested-toggle3").
																				Aria("label", "Toggle details for").
																				Aria("expanded", true).
																				Aria("controls", "data-list-expandable-nested-nested-content3").
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
																),
															app.Div().
																Class("pf-c-data-list__item-content").
																Body(
																	app.Div().
																		Class("pf-c-data-list__cell").
																		Body(
																			app.Div().
																				ID("data-list-expandable-nested-nested-item-3").
																				Body(
																					app.Text("Nested row 3"),
																				),
																		),
																),
														),
													app.Section().
														Class("pf-c-data-list__expandable-content").
														ID("data-list-expandable-nested-nested-content3").
														Aria("label", "Nested row 3 details").
														Body(
															app.Div().
																Class("pf-c-data-list__expandable-content-body").
																Body(
																	app.Text("Nested row 3 expanded content."),
																),
														),
												),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-expandable-nested-item-2").
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
												Aria("labelledby", "data-list-expandable-nested-toggle2 data-list-expandable-nested-item2").
												ID("data-list-expandable-nested-toggle2").
												Aria("label", "Toggle details for").
												Aria("expanded", "false").
												Aria("controls", "data-list-expandable-nested-content2").
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
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Div().
												ID("data-list-expandable-nested-item-2").
												Body(
													app.Text("Secondary content"),
												),
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
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
														ID("data-list-expandable-nested-item-2-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-expandable-nested-item-2-dropdown-kebab-button").
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
						ID("data-list-expandable-nested-content2").
						Aria("label", "Secondary content details").
						Hidden(true).
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body").
								Body(
									app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded").
				Aria("labelledby", "data-list-expandable-nested-item-3").
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
												Aria("labelledby", "data-list-expandable-nested-toggle3 data-list-expandable-nested-item3").
												ID("data-list-expandable-nested-toggle3").
												Aria("label", "Toggle details for").
												Aria("expanded", true).
												Aria("controls", "data-list-expandable-nested-content3").
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
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell pf-m-icon").
										Body(
											app.I().
												Class("fas fa-code-branch").
												Aria("hidden", true),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Div().
												ID("data-list-expandable-nested-item-3").
												Body(
													app.Text("Tertiary content"),
												),
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
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
														ID("data-list-expandable-nested-item-3-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-expandable-nested-item-3-dropdown-kebab-button").
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
						ID("data-list-expandable-nested-content3").
						Aria("label", "Tertiary content details").
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body pf-m-no-padding").
								Body(
									app.Text("This expanded section has no padding."),
								),
						),
				),
		)
}
