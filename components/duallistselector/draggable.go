package duallistselector

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Draggable struct {
	app.Compo
}

func (c *Draggable) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				ID("draggable-help").
				Body(
					app.Text("Activate the reorder button and use the arrow keys to reorder the list or use your mouse to drag/reorder. Press escape to cancel the reordering."),
				),
			app.Div().
				Class("pf-c-dual-list-selector").
				Body(
					app.Div().
						Class("pf-c-dual-list-selector__pane pf-m-available").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__header").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__title").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__title-text").
												Body(
													app.Text("Available options"),
												),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__tools").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__tools-filter").
										Body(
											app.Div().
												Class("pf-c-search-input").
												Body(
													app.Div().
														Class("pf-c-search-input__bar").
														Body(
															app.Span().
																Class("pf-c-search-input__text").
																Body(
																	app.Span().
																		Class("pf-c-search-input__icon").
																		Body(
																			app.I().
																				Class("fas fa-search fa-fw").
																				Aria("hidden", true),
																		),
																	app.Input().
																		Class("pf-c-search-input__text-input").
																		Type("text").
																		Placeholder("false").
																		Aria("label", "Filter options"),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-dual-list-selector__tools-actions").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Sort").
												Body(
													app.I().
														Class("fas fa-sort-amount-down").
														Aria("hidden", true),
												),
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("dropdown-kebab-draggable-available-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu").
														Aria("labelledby", "dropdown-kebab-draggable-available-button").
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
								Class("pf-c-dual-list-selector__status").
								Body(
									app.Span().
										Class("pf-c-dual-list-selector__status-text").
										ID("draggable-available-status-text").
										Body(
											app.Text("0 of 5 items selected"),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__menu").
								Body(
									app.Ul().
										Class("pf-c-dual-list-selector__list").
										Aria("role", "listbox").
										Aria("labelledby", "draggable-available-status-text").
										Aria("multiselectable", true).
										Aria("activedescendant", true).
										Body(
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row").
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				Body(
																					app.Text("Item1"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row").
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				Body(
																					app.Text("Item4"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row").
														Body(
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				Body(
																					app.Text("Item6"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__controls").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__controls-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
										Aria("label", "Add selected").
										Body(
											app.I().
												Class("fas fa-fw fa-angle-right"),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__controls-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Add all").
										Body(
											app.I().
												Class("fas fa-fw fa-angle-double-right"),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__controls-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
										Aria("label", "Remove all").
										Body(
											app.I().
												Class("fas fa-fw fa-angle-double-left"),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__controls-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
										Aria("label", "Remove selected").
										Body(
											app.I().
												Class("fas fa-fw fa-angle-left"),
										),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__pane pf-m-chosen").
						Body(
							app.Div().
								Class("pf-c-dual-list-selector__header").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__title").
										Body(
											app.Div().
												Class("pf-c-dual-list-selector__title-text").
												Body(
													app.Text("Chosen options"),
												),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__tools").
								Body(
									app.Div().
										Class("pf-c-dual-list-selector__tools-filter").
										Body(
											app.Div().
												Class("pf-c-search-input").
												Body(
													app.Div().
														Class("pf-c-search-input__bar").
														Body(
															app.Span().
																Class("pf-c-search-input__text").
																Body(
																	app.Span().
																		Class("pf-c-search-input__icon").
																		Body(
																			app.I().
																				Class("fas fa-search fa-fw").
																				Aria("hidden", true),
																		),
																	app.Input().
																		Class("pf-c-search-input__text-input").
																		Type("text").
																		Placeholder("false").
																		Aria("label", "Filter options"),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-dual-list-selector__tools-actions").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Sort").
												Body(
													app.I().
														Class("fas fa-sort-amount-down").
														Aria("hidden", true),
												),
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("dropdown-kebab-draggable-chosen-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu").
														Aria("labelledby", "dropdown-kebab-draggable-chosen-button").
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
								Class("pf-c-dual-list-selector__status").
								Body(
									app.Span().
										Class("pf-c-dual-list-selector__status-text").
										ID("draggable-chosen-status-text").
										Body(
											app.Text("0 of 0 items selected"),
										),
								),
							app.Div().
								Class("pf-c-dual-list-selector__menu").
								Body(
									app.Ul().
										Class("pf-c-dual-list-selector__list").
										Aria("role", "listbox").
										Aria("labelledby", "draggable-chosen-status-text").
										Aria("multiselectable", true).
										Aria("activedescendant", true).
										Body(
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__draggable").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Disabled(true).
																		Aria("pressed", "false").
																		Aria("label", "Reorder").
																		ID("draggable-list-item-2-draggable-button").
																		Aria("labelledby", "draggable-list-item-2-draggable-button draggable-list-item-2-item-text").
																		Aria("describedby", "draggable-help").
																		Body(
																			app.I().
																				Class("fas fa-grip-vertical").
																				Aria("hidden", true),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				ID("draggable-list-item-2-item-text").
																				Body(
																					app.Text("Item2 - draggable icon disabled"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__draggable").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("pressed", "false").
																		Aria("label", "Reorder").
																		ID("draggable-list-item-3-draggable-button").
																		Aria("labelledby", "draggable-list-item-3-draggable-button draggable-list-item-3-item-text").
																		Aria("describedby", "draggable-help").
																		Body(
																			app.I().
																				Class("fas fa-grip-vertical").
																				Aria("hidden", true),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				ID("draggable-list-item-3-item-text").
																				Body(
																					app.Text("Item3"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row pf-m-ghost-row").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__draggable").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Disabled(true).
																		Aria("pressed", "false").
																		Aria("label", "Reorder").
																		ID("draggable-list-item-5-draggable-button").
																		Aria("labelledby", "draggable-list-item-5-draggable-button draggable-list-item-5-item-text").
																		Aria("describedby", "draggable-help").
																		Body(
																			app.I().
																				Class("fas fa-grip-vertical").
																				Aria("hidden", true),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				ID("draggable-list-item-5-item-text").
																				Body(
																					app.Text("Item5 - ghost row"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-dual-list-selector__list-item").
												Aria("role", "option").
												Body(
													app.Div().
														Class("pf-c-dual-list-selector__list-item-row pf-m-selected").
														Body(
															app.Div().
																Class("pf-c-dual-list-selector__draggable").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("pressed", "false").
																		Aria("label", "Reorder").
																		ID("draggable-list-item-7-draggable-button").
																		Aria("labelledby", "draggable-list-item-7-draggable-button draggable-list-item-7-item-text").
																		Aria("describedby", "draggable-help").
																		Body(
																			app.I().
																				Class("fas fa-grip-vertical").
																				Aria("hidden", true),
																		),
																),
															app.Span().
																Class("pf-c-dual-list-selector__item").
																Body(
																	app.Span().
																		Class("pf-c-dual-list-selector__item-main").
																		Body(
																			app.Span().
																				Class("pf-c-dual-list-selector__item-text").
																				ID("draggable-list-item-7-item-text").
																				Body(
																					app.Text("Item7 - selected"),
																				),
																		),
																	app.Span().
																		Class("pf-c-dual-list-selector__item-count").
																		Body(
																			app.Span().
																				Class("pf-c-badge pf-m-read"),
																		),
																),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-screen-reader").
				Aria("live", "assertive").
				Body(
					app.Text("This is the aria-live section that provides real-time feedback to the user."),
				),
		)
}
