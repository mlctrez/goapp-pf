package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list pf-m-compact pf-m-grid-sm").
		Aria("role", "list").
		Aria("label", "Compact data list example").
		ID("data-list-compact").
		Body(
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-compact-item-1").
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
												Name("data-list-compact-item-1-checkbox").
												Aria("labelledby", "data-list-compact-item-1"),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-compact-item-1").
												Body(
													app.Text("Primary content"),
												),
											app.Text("Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Secondary content. Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("Tertiary Content"),
												),
											app.Text("Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("More Content"),
												),
											app.Text("Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												Body(
													app.Text("More Content"),
												),
											app.Text("Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
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
														ID("data-list-compact-item-1-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-compact-item-1-dropdown-kebab-button").
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
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-compact-item-2").
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
												Name("data-list-compact-item-2-checkbox").
												Aria("labelledby", "data-list-compact-item-2"),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-compact-item-2").
												Body(
													app.Text("Primary content - lorem ipsum"),
												),
											app.Text("dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Secondary content. Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-action pf-m-hidden-on-lg").
								Body(
									app.Div().
										Class("pf-c-data-list__action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("data-list-compact-item-2-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-compact-item-2-dropdown-kebab-button").
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
							app.Div().
								Class("pf-c-data-list__item-action pf-m-hidden pf-m-visible-on-lg").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("button").
										Body(
											app.Text("Primary"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Secondary"),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-compact-item-3").
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
												Name("data-list-compact-item-3-checkbox").
												Aria("labelledby", "data-list-compact-item-3"),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-compact-item-3").
												Body(
													app.Text("Primary content - lorem ipsum"),
												),
											app.Text("dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Secondary content. Dolor sit amet, consectetur adipisicing elit, sed do eiusmod."),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-action pf-m-hidden-on-xl").
								Body(
									app.Div().
										Class("pf-c-data-list__action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("data-list-compact-item-3-dropdown-kebab-button").
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
														Aria("labelledby", "data-list-compact-item-3-dropdown-kebab-button").
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
							app.Div().
								Class("pf-c-data-list__item-action pf-m-hidden pf-m-visible-on-xl").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("button").
										Body(
											app.Text("Primary"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Secondary"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Secondary"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Secondary"),
										),
								),
						),
				),
		)
}
