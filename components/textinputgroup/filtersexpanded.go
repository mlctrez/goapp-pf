package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FiltersExpanded struct {
	app.Compo
}

func (c *FiltersExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-text-input-group").
		Body(
			app.Div().
				Class("pf-c-text-input-group__main").
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
																ID("-chip-groupchip_one_select_collapsed").
																Body(
																	app.Text("Chip one"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_one_select_collapsed -chip-groupchip_one_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_one_select_collapsed").
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
																ID("-chip-groupchip_two_select_collapsed").
																Body(
																	app.Text("Chip two"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_two_select_collapsed -chip-groupchip_two_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_two_select_collapsed").
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
																ID("-chip-groupchip_three_select_collapsed").
																Body(
																	app.Text("Chip three"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_three_select_collapsed -chip-groupchip_three_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_three_select_collapsed").
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
																ID("-chip-groupchip_four_select_collapsed").
																Body(
																	app.Text("Chip four"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_four_select_collapsed -chip-groupchip_four_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_four_select_collapsed").
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
																ID("-chip-groupchip_five_select_collapsed").
																Body(
																	app.Text("Chip five"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_five_select_collapsed -chip-groupchip_five_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_five_select_collapsed").
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
																ID("-chip-groupchip_six_select_collapsed").
																Body(
																	app.Text("Chip six"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_six_select_collapsed -chip-groupchip_six_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_six_select_collapsed").
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
																ID("-chip-groupchip_seven_select_collapsed").
																Body(
																	app.Text("Chip seven"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_seven_select_collapsed -chip-groupchip_seven_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_seven_select_collapsed").
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
																ID("-chip-groupchip_eight_select_collapsed").
																Body(
																	app.Text("Chip eight"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_eight_select_collapsed -chip-groupchip_eight_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_eight_select_collapsed").
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
																ID("-chip-groupchip_nine_select_collapsed").
																Body(
																	app.Text("Chip nine"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_nine_select_collapsed -chip-groupchip_nine_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_nine_select_collapsed").
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
																ID("-chip-groupchip_ten_select_collapsed").
																Body(
																	app.Text("Chip ten"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_ten_select_collapsed -chip-groupchip_ten_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_ten_select_collapsed").
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
																ID("-chip-groupchip_eleven_select_collapsed").
																Body(
																	app.Text("Chip eleven"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_eleven_select_collapsed -chip-groupchip_eleven_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_eleven_select_collapsed").
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
																ID("-chip-groupchip_twelve_select_collapsed").
																Body(
																	app.Text("Chip twelve"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_twelve_select_collapsed -chip-groupchip_twelve_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_twelve_select_collapsed").
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
																ID("-chip-groupchip_thirteen_select_collapsed").
																Body(
																	app.Text("Chip thirteen"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_thirteen_select_collapsed -chip-groupchip_thirteen_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_thirteen_select_collapsed").
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
																ID("-chip-groupchip_fourteen_select_collapsed").
																Body(
																	app.Text("Chip fourteen"),
																),
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("labelledby", "-chip-groupremove_chip_fourteen_select_collapsed -chip-groupchip_fourteen_select_collapsed").
																Aria("label", "Remove").
																ID("-chip-groupremove_chip_fourteen_select_collapsed").
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
					app.Span().
						Class("pf-c-text-input-group__text").
						Body(
							app.Input().
								Class("pf-c-text-input-group__text-input").
								Type("text").
								Value(true).
								Aria("label", "Type to filter"),
						),
				),
			app.Div().
				Class("pf-c-text-input-group__utilities").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Clear input").
						Body(
							app.I().
								Class("fas fa-times fa-fw").
								Aria("hidden", true),
						),
				),
		)
}
