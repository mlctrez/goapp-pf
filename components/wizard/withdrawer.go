package wizard

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDrawer struct {
	app.Compo
}

func (c *WithDrawer) Render() app.UI {
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
				Class("pf-c-wizard__toggle").
				Aria("expanded", "false").
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
				Class("pf-c-drawer pf-m-expanded pf-m-inline").
				Body(
					app.Div().
						Class("pf-c-drawer__main").
						Body(
							app.Div().
								Class("pf-c-drawer__content").
								Body(
									app.Div().
										Class("pf-c-wizard__outer-wrap").
										Body(
											app.Div().
												Class("pf-c-wizard__inner-wrap").
												Body(
													app.Main().
														Class("pf-c-wizard__main").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-wizard__main-body").
																Body(
																	app.Nav().
																		Class("pf-c-wizard__nav").
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
																								For("wizard-with-drawer-example-form-field1").
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
																								ID("wizard-with-drawer-example-form-field1").
																								Name("wizard-with-drawer-example-form-field1"),
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
																								For("wizard-with-drawer-example-form-field2").
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
																								ID("wizard-with-drawer-example-form-field2").
																								Name("wizard-with-drawer-example-form-field2"),
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
																								For("wizard-with-drawer-example-form-field3").
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
																								ID("wizard-with-drawer-example-form-field3").
																								Name("wizard-with-drawer-example-form-field3"),
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
																								For("wizard-with-drawer-example-form-field4").
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
																								ID("wizard-with-drawer-example-form-field4").
																								Name("wizard-with-drawer-example-form-field4"),
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
																								For("wizard-with-drawer-example-form-field5").
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
																								ID("wizard-with-drawer-example-form-field5").
																								Name("wizard-with-drawer-example-form-field5"),
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
																								For("wizard-with-drawer-example-form-field6").
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
																								ID("wizard-with-drawer-example-form-field6").
																								Name("wizard-with-drawer-example-form-field6"),
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
																								For("wizard-with-drawer-example-form-field7").
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
																								ID("wizard-with-drawer-example-form-field7").
																								Name("wizard-with-drawer-example-form-field7"),
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
								),
							app.Div().
								Class("pf-c-drawer__panel pf-m-light-200").
								Body(
									app.Div().
										Class("pf-c-drawer__body").
										Body(
											app.Div().
												Class("pf-c-drawer__head").
												Body(
													app.Div().
														Class("pf-c-drawer__actions").
														Body(
															app.Div().
																Class("pf-c-drawer__close").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Close drawer panel").
																		Body(
																			app.I().
																				Class("fas fa-times").
																				Aria("hidden", true),
																		),
																),
														),
													app.Text("drawer-panel"),
												),
										),
								),
						),
				),
		)
}
