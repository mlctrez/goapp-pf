package wizard

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Finished struct {
	app.Compo
}

func (c *Finished) Render() app.UI {
	return app.Div().
		Class("pf-c-wizard pf-m-finished").
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
				Class("pf-c-wizard__outer-wrap").
				Body(
					app.Div().
						Class("pf-c-wizard__inner-wrap").
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
														Class("pf-c-wizard__nav-link").
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
																		Class("pf-c-wizard__nav-link").
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
											app.Div().
												Class("pf-l-bullseye").
												Body(
													app.Div().
														Class("pf-c-empty-state pf-m-lg").
														Body(
															app.Div().
																Class("pf-c-empty-state__content").
																Body(
																	app.I().
																		Class("fas fa- fa-cogs pf-c-empty-state__icon").
																		Aria("hidden", true),
																	app.H1().
																		Class("pf-c-title pf-m-lg").
																		ID("wizard-finished-empty-state-title").
																		Body(
																			app.Text("Validating credentials"),
																		),
																	app.Div().
																		Class("pf-c-empty-state__body").
																		Body(
																			app.Div().
																				Class("pf-c-progress pf-m-singleline").
																				ID("progress-singleline-example").
																				Body(
																					app.Div().
																						Class("pf-c-progress__description").
																						ID("progress-singleline-example-description"),
																					app.Div().
																						Class("pf-c-progress__status").
																						Aria("hidden", true).
																						Body(
																							app.Span().
																								Class("pf-c-progress__measure").
																								Body(
																									app.Text("33%"),
																								),
																						),
																					app.Div().
																						Class("pf-c-progress__bar").
																						Aria("role", "progressbar").
																						Aria("valuemin", "0").
																						Aria("valuemax", "100").
																						Aria("valuenow", "33").
																						Aria("labelledby", "wizard-finished-empty-state-title").
																						Aria("label", "Progress status").
																						Body(
																							app.Div().
																								Class("pf-c-progress__indicator").
																								Style("width", "33%;"),
																						),
																				),
																		),
																	app.Div().
																		Class("pf-c-empty-state__body").
																		Body(
																			app.Text("Description can be used to further elaborate on the validation step, or give the user a better idea of how long the process will take."),
																		),
																	app.Div().
																		Class("pf-c-empty-state__secondary").
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
