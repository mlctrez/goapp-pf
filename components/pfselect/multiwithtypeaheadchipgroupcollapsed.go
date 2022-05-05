package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MultiWithTypeaheadChipGroupCollapsed struct {
	app.Compo
}

func (c *MultiWithTypeaheadChipGroupCollapsed) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-multi-typeahead-expanded-selected-label").
				Hidden(true).
				Body(
					app.Text("New"),
				),
			app.Div().
				Class("pf-c-select__toggle pf-m-typeahead").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Div().
								Class("pf-c-chip-group").
								Body(
									app.Div().
										Class("pf-c-chip-group__main").
										Body(
											app.Ul().
												Class("pf-c-chip-group__list").
												Aria("role", "list").
												Aria("label", "Chip group list").
												Body(
													app.Li().
														Class("pf-c-chip-group__list-item").
														Body(
															app.Div().
																Class("pf-c-chip").
																Body(
																	app.Span().
																		Class("pf-c-chip__text").
																		ID("select-multi-typeahead-expanded-selected-chip_one").
																		Body(
																			app.Text("Arkansas"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-expanded-selected_chip_one select-multi-typeahead-expanded-selected-chip_one").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-expanded-selected_chip_one").
																		Body(
																			app.I().
																				Class("fas fa-times").
																				Aria("hidden", true),
																		),
																),
														),
													app.Li().
														Class("pf-c-chip-group__list-item").
														Body(
															app.Div().
																Class("pf-c-chip").
																Body(
																	app.Span().
																		Class("pf-c-chip__text").
																		ID("select-multi-typeahead-expanded-selected-chip_two").
																		Body(
																			app.Text("Massachusetts"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-expanded-selected_chip_two select-multi-typeahead-expanded-selected-chip_two").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-expanded-selected_chip_two").
																		Body(
																			app.I().
																				Class("fas fa-times").
																				Aria("hidden", true),
																		),
																),
														),
													app.Li().
														Class("pf-c-chip-group__list-item").
														Body(
															app.Div().
																Class("pf-c-chip").
																Body(
																	app.Span().
																		Class("pf-c-chip__text").
																		ID("select-multi-typeahead-expanded-selected-chip_three").
																		Body(
																			app.Text("New Mexico"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-expanded-selected_chip_three select-multi-typeahead-expanded-selected-chip_three").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-expanded-selected_chip_three").
																		Body(
																			app.I().
																				Class("fas fa-times").
																				Aria("hidden", true),
																		),
																),
														),
													app.Li().
														Class("pf-c-chip-group__list-item").
														Body(
															app.Button().
																Class("pf-c-chip pf-m-overflow").
																Body(
																	app.Span().
																		Class("pf-c-chip__text").
																		Body(
																			app.Text("2 more"),
																		),
																),
														),
												),
										),
								),
							app.Input().
								Class("pf-c-form-control pf-c-select__toggle-typeahead").
								Type("text").
								ID("select-multi-typeahead-expanded-selected-typeahead").
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-clear").
						Type("button").
						Aria("label", "Clear all").
						Body(
							app.I().
								Class("fas fa-times-circle").
								Aria("hidden", true),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-button").
						Type("button").
						ID("select-multi-typeahead-expanded-selected-toggle").
						Aria("haspopup", true).
						Aria("expanded", true).
						Aria("labelledby", "select-multi-typeahead-expanded-selected-label select-multi-typeahead-expanded-selected-toggle").
						Aria("label", "Select").
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-multi-typeahead-expanded-selected-label").
				Aria("role", "listbox").
				Body(
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Mark().
										Class("pf-c-select__menu-item--match").
										Body(
											app.Text("New"),
										),
									app.Text("Jersey"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Mark().
										Class("pf-c-select__menu-item--match").
										Body(
											app.Text("New"),
										),
									app.Text("York"),
								),
						),
				),
		)
}
