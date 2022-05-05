package wizard

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NavExpandedMobile struct {
	app.Compo
}

func (c *NavExpandedMobile) Render() app.UI {
	return app.Div().
		Class("pf-c-wizard").
		Body(
			app.Div().
				Class("pf-c-wizard__header").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-wizard__close").
						Type("button").
						Aria("label", "Close").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
					app.H1().
						Class("pf-c-title pf-m-3xl pf-c-wizard__title").
						Body(
							app.Text("Wizard title"),
						),
					app.Div().
						Class("pf-c-wizard__description").
						Body(
							app.Text("Here is where the description goes"),
						),
				),
			app.Button().
				Aria("label", "Wizard Header Toggle").
				Class("pf-c-wizard__toggle pf-m-expanded").
				Aria("expanded", true).
				Body(
					app.Span().
						Class("pf-c-wizard__toggle-list").
						Body(
							app.Span().
								Class("pf-c-wizard__toggle-list-item").
								Body(
									app.Span().
										Class("pf-c-wizard__toggle-num").
										Body(
											app.Text("2"),
										),
									app.Text("Configuration"),
									app.I().
										Class("fas fa-angle-right pf-c-wizard__toggle-separator").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-wizard__toggle-list-item").
								Body(
									app.Text("Substep B"),
								),
						),
					app.Span().
						Class("pf-c-wizard__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-wizard__outer-wrap").
				Body(
					app.Div().
						Class("pf-c-wizard__inner-wrap").
						Body(
							app.Nav().
								Class("pf-c-wizard__nav pf-m-expanded").
								Aria("label", "Steps").
								Body(
									app.Ol().
										Class("pf-c-wizard__nav-list").
										Body(
											app.Li().
												Class("pf-c-wizard__nav-item").
												Body(
													app.Button().
														Class("pf-c-wizard__nav-link").
														Body(
															app.Text("Information"),
														),
												),
											app.Li().
												Class("pf-c-wizard__nav-item").
												Body(
													app.Button().
														Class("pf-c-wizard__nav-link pf-m-current").
														Body(
															app.Text("Configuration"),
														),
													app.Ol().
														Class("pf-c-wizard__nav-list").
														Body(
															app.Li().
																Class("pf-c-wizard__nav-item").
																Body(
																	app.Button().
																		Class("pf-c-wizard__nav-link").
																		Body(
																			app.Text("Substep A"),
																		),
																),
															app.Li().
																Class("pf-c-wizard__nav-item").
																Body(
																	app.Button().
																		Class("pf-c-wizard__nav-link pf-m-current").
																		Aria("current", "page").
																		Body(
																			app.Text("Substep B"),
																		),
																),
															app.Li().
																Class("pf-c-wizard__nav-item").
																Body(
																	app.Button().
																		Class("pf-c-wizard__nav-link").
																		Body(
																			app.Text("Substep C"),
																		),
																),
														),
												),
											app.Li().
												Class("pf-c-wizard__nav-item").
												Body(
													app.Button().
														Class("pf-c-wizard__nav-link").
														Body(
															app.Text("Additional"),
														),
												),
											app.Li().
												Class("pf-c-wizard__nav-item").
												Body(
													app.Button().
														Class("pf-c-wizard__nav-link").
														Disabled(true).
														Body(
															app.Text("Review"),
														),
												),
										),
								),
							app.Main().
								Class("pf-c-wizard__main").
								TabIndex(0).
								Body(
									app.Div().
										Class("pf-c-wizard__main-body").
										Body(
											app.Form().
												NoValidate(true).
												Class("pf-c-form").
												Body(
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field1").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 1"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field1").
																		Name("-form-field1"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field2").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 2"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field2").
																		Name("-form-field2"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field3").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 3"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field3").
																		Name("-form-field3"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field4").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 4"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field4").
																		Name("-form-field4"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field5").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 5"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field5").
																		Name("-form-field5"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field6").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 6"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field6").
																		Name("-form-field6"),
																),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-form__group-label").
																Body(
																	app.Label().
																		Class("pf-c-form__label").
																		For("-form-field7").
																		Body(
																			app.Span().
																				Class("pf-c-form__label-text").
																				Body(
																					app.Text("Field 7"),
																				),
																			app.Span().
																				Class("pf-c-form__label-required").
																				Aria("hidden", true).
																				Body(
																					app.Text("*"),
																				),
																		),
																),
															app.Div().
																Class("pf-c-form__group-control").
																Body(
																	app.Input().
																		Class("pf-c-form-control").
																		Required(true).
																		Type("text").
																		ID("-form-field7").
																		Name("-form-field7"),
																),
														),
												),
										),
								),
						),
					app.Footer().
						Class("pf-c-wizard__footer").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("submit").
								Body(
									app.Text("Next"),
								),
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Back"),
								),
							app.Div().
								Class("pf-c-wizard__footer-cancel").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Body(
											app.Text("Cancel"),
										),
								),
						),
				),
		)
}
