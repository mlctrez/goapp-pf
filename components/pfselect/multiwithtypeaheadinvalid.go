package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MultiWithTypeaheadInvalid struct {
	app.Compo
}

func (c *MultiWithTypeaheadInvalid) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded pf-m-invalid").
		Body(
			app.Span().
				ID("select-multi-typeahead-invalid-label").
				Hidden(true).
				Body(
					app.Text("Choose states"),
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
																		ID("select-multi-typeahead-invalid-chip_one").
																		Body(
																			app.Text("Arkansas"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-invalid_chip_one select-multi-typeahead-invalid-chip_two").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-invalid_chip_one").
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
																		ID("select-multi-typeahead-invalid-chip_two").
																		Body(
																			app.Text("Massachusetts"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-invalid_chip_two select-multi-typeahead-invalid-chip_two").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-invalid_chip_two").
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
																		ID("select-multi-typeahead-invalid-chip_three").
																		Body(
																			app.Text("New Mexico"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-invalid_chip_three select-multi-typeahead-invalid-chip_three").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-invalid_chip_three").
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
																		ID("select-multi-typeahead-invalid-chip_four").
																		Body(
																			app.Text("Ohio"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-invalid_chip_four select-multi-typeahead-invalid-chip_four").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-invalid_chip_four").
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
																		ID("select-multi-typeahead-invalid-chip_five").
																		Body(
																			app.Text("Washington"),
																		),
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("labelledby", "remove_select-multi-typeahead-invalid_chip_five select-multi-typeahead-invalid-chip_five").
																		Aria("label", "Remove").
																		ID("remove_select-multi-typeahead-invalid_chip_five").
																		Body(
																			app.I().
																				Class("fas fa-times").
																				Aria("hidden", true),
																		),
																),
														),
												),
										),
								),
							app.Input().
								Class("pf-c-form-control pf-c-select__toggle-typeahead").
								Type("text").
								ID("select-multi-typeahead-invalid-typeahead").
								Aria("invalid", true).
								Value("Invalid").
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
					app.Span().
						Class("pf-c-select__toggle-status-icon").
						Body(
							app.I().
								Class("fas fa-exclamation-circle").
								Aria("hidden", true),
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
						ID("select-multi-typeahead-invalid-toggle").
						Aria("haspopup", true).
						Aria("expanded", true).
						Aria("labelledby", "select-multi-typeahead-invalid-label select-multi-typeahead-invalid-toggle").
						Aria("label", "Select").
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-multi-typeahead-invalid-label").
				Aria("role", "listbox").
				Body(
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("Alabama"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("Florida"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("New\n        \u00a0Jersey"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("New\n        \u00a0York"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("North Carolina"),
								),
						),
				),
		)
}
