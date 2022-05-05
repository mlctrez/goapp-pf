package duallistselector

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AvailableItemSelected struct {
	app.Compo
}

func (c *AvailableItemSelected) Render() app.UI {
	return app.Div().
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
												ID("dropdown-kebab-available-item-selected-available-button").
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
												Aria("labelledby", "dropdown-kebab-available-item-selected-available-button").
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
								ID("available-item-selected-available-status-text").
								Body(
									app.Text("1 of 5 items selected"),
								),
						),
					app.Div().
						Class("pf-c-dual-list-selector__menu").
						Body(
							app.Ul().
								Class("pf-c-dual-list-selector__list").
								Aria("role", "listbox").
								Aria("labelledby", "available-item-selected-available-status-text").
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
												Class("pf-c-dual-list-selector__list-item-row pf-m-selected").
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
																			app.Text("Item2"),
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
																			app.Text("Item5"),
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
												ID("dropdown-kebab-available-item-selected-chosen-button").
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
												Aria("labelledby", "dropdown-kebab-available-item-selected-chosen-button").
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
								ID("available-item-selected-chosen-status-text").
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
								Aria("labelledby", "available-item-selected-chosen-status-text").
								Aria("multiselectable", true).
								Aria("activedescendant", true),
						),
				),
		)
}
