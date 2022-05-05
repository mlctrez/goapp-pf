package calendarmonth

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type RangeEndDateSelectedStartDateFocused struct {
	app.Compo
}

func (c *RangeEndDateSelectedStartDateFocused) Render() app.UI {
	return app.Div().
		Class("pf-c-calendar-month").
		Body(
			app.Div().
				Class("pf-c-calendar-month__header").
				Body(
					app.Div().
						Class("pf-c-calendar-month__header-nav-control pf-m-prev-month").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Previous month").
								Body(
									app.I().
										Class("fas fa-angle-left").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-calendar-month__header-month").
						Body(
							app.Div().
								Class("pf-c-select").
								Body(
									app.Span().
										ID("calendar-month-range-end-date-selected-month-select-label").
										Hidden(true).
										Body(
											app.Text("Month"),
										),
									app.Button().
										Class("pf-c-select__toggle").
										Type("button").
										ID("calendar-month-range-end-date-selected-month-select-toggle").
										Aria("haspopup", true).
										Aria("expanded", "false").
										Aria("labelledby", "calendar-month-range-end-date-selected-month-select-label calendar-month-range-end-date-selected-month-select-toggle").
										Body(
											app.Div().
												Class("pf-c-select__toggle-wrapper").
												Body(
													app.Span().
														Class("pf-c-select__toggle-text").
														Body(
															app.Text("October"),
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
										Aria("labelledby", "calendar-month-range-end-date-selected-month-select-label").
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
						Class("pf-c-calendar-month__header-year").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("number").
								Value("2020").
								ID("calendar-month-range-end-date-selected-year").
								Aria("label", "Select year"),
						),
					app.Div().
						Class("pf-c-calendar-month__header-nav-control pf-m-next-month").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Next month").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Table().
				Class("pf-c-calendar-month__calendar").
				Body(
					app.THead().
						Class("pf-c-calendar-month__days").
						Body(
							app.Tr().
								Class("pf-c-calendar-month__days-row").
								Body(
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Sunday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("S"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Monday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("M"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Tuesday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("T"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Wednesday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("W"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Thursday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("T"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Friday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("F"),
												),
										),
									app.Th().
										Class("pf-c-calendar-month__day").
										Body(
											app.Span().
												Class("pf-screen-reader").
												Body(
													app.Text("Saturday"),
												),
											app.Span().
												Aria("hidden", true).
												Body(
													app.Text("S"),
												),
										),
								),
						),
					app.TBody().
						Class("pf-c-calendar-month__dates").
						Body(
							app.Tr().
								Class("pf-c-calendar-month__dates-row").
								Body(
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-adjacent-month").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "29 October 2020").
												Body(
													app.Text("29"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-adjacent-month").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "30 October 2020").
												Body(
													app.Text("30"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "1 October 2020").
												Body(
													app.Text("1"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "2 October 2020").
												Body(
													app.Text("2"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "3 October 2020").
												Body(
													app.Text("3"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "4 October 2020").
												Body(
													app.Text("4"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "5 October 2020").
												Body(
													app.Text("5"),
												),
										),
								),
							app.Tr().
								Class("pf-c-calendar-month__dates-row").
								Body(
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "6 October 2020").
												Body(
													app.Text("6"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "7 October 2020").
												Body(
													app.Text("7"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "8 October 2020").
												Body(
													app.Text("8"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-current").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "9 October 2020").
												Body(
													app.Text("9"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "10 October 2020").
												Body(
													app.Text("10"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-selected pf-m-start-range pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date pf-m-focus").
												Type("button").
												Aria("label", "11 October 2020").
												Body(
													app.Text("11"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "12 October 2020").
												Body(
													app.Text("12"),
												),
										),
								),
							app.Tr().
								Class("pf-c-calendar-month__dates-row").
								Body(
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "13 October 2020").
												Body(
													app.Text("13"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "14 October 2020").
												Body(
													app.Text("14"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "15 October 2020").
												Body(
													app.Text("15"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "16 October 2020").
												Body(
													app.Text("16"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "17 October 2020").
												Body(
													app.Text("17"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "18 October 2020").
												Body(
													app.Text("18"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "19 October 2020").
												Body(
													app.Text("19"),
												),
										),
								),
							app.Tr().
								Class("pf-c-calendar-month__dates-row").
								Body(
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "20 October 2020").
												Disabled(true).
												Body(
													app.Text("20"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "21 October 2020").
												Disabled(true).
												Body(
													app.Text("21"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "22 October 2020").
												Disabled(true).
												Body(
													app.Text("22"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "23 October 2020").
												Disabled(true).
												Body(
													app.Text("23"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "24 October 2020").
												Disabled(true).
												Body(
													app.Text("24"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "25 October 2020").
												Disabled(true).
												Body(
													app.Text("25"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range pf-m-disabled").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "26 October 2020").
												Disabled(true).
												Body(
													app.Text("26"),
												),
										),
								),
							app.Tr().
								Class("pf-c-calendar-month__dates-row").
								Body(
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "27 October 2020").
												Body(
													app.Text("27"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-in-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "28 October 2020").
												Body(
													app.Text("28"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-selected pf-m-in-range pf-m-end-range").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "29 October 2020").
												Body(
													app.Text("29"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "30 October 2020").
												Body(
													app.Text("30"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "31 October 2020").
												Body(
													app.Text("31"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-adjacent-month").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "1 October 2020").
												Body(
													app.Text("1"),
												),
										),
									app.Td().
										Class("pf-c-calendar-month__dates-cell pf-m-adjacent-month").
										Body(
											app.Button().
												Class("pf-c-calendar-month__date").
												Type("button").
												Aria("label", "2 October 2020").
												Body(
													app.Text("2"),
												),
										),
								),
						),
				),
		)
}
