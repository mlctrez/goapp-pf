package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ToggleGroup struct {
	app.Compo
}

func (c *ToggleGroup) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-toggle-group-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group pf-m-toggle-group pf-m-show-on-lg").
								Body(
									app.Div().
										Class("pf-c-toolbar__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Show filters").
												Aria("expanded", "false").
												Aria("controls", "toolbar-toggle-group-example-expandable-content").
												Body(
													app.I().
														Class("fas fa-filter").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item pf-m-search-filter").
										Body(
											app.Div().
												Class("pf-c-input-group").
												Aria("label", "search filter").
												Aria("role", "group").
												Body(
													app.Div().
														Class("pf-c-select").
														Style("width", " 175px").
														Body(
															app.Span().
																ID("toolbar-toggle-group-example-select-name-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-toggle-group-example-select-name-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-toggle-group-example-select-name-label toolbar-toggle-group-example-select-name-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-filter").
																						Aria("hidden", true),
																				),
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Name"),
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
																Aria("labelledby", "toolbar-toggle-group-example-select-name-label").
																Hidden(true).
																Style("width", " 175px").
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
																				Aria("label", "Filter by name"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-toolbar__group pf-m-filter-group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Span().
																ID("toolbar-toggle-group-example-select-checkbox-status-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-toggle-group-example-select-checkbox-status-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-toggle-group-example-select-checkbox-status-label toolbar-toggle-group-example-select-checkbox-status-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Status"),
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
															app.Div().
																Class("pf-c-select__menu").
																Hidden(true).
																Body(
																	app.FieldSet().
																		Class("pf-c-select__menu-fieldset").
																		Aria("label", "Select input").
																		Body(
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item pf-m-description").
																				For("toolbar-toggle-group-example-select-checkbox-status-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-status-active").
																						Name("toolbar-toggle-group-example-select-checkbox-status-active"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Active"),
																						),
																					app.Span().
																						Class("pf-c-check__description").
																						Body(
																							app.Text("This is a description"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item pf-m-description").
																				For("toolbar-toggle-group-example-select-checkbox-status-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-status-canceled").
																						Name("toolbar-toggle-group-example-select-checkbox-status-canceled"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Canceled"),
																						),
																					app.Span().
																						Class("pf-c-check__description").
																						Body(
																							app.Text("This is a really long description that describes the menu item. This is a really long description that describes the menu item."),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-status-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-status-paused").
																						Name("toolbar-toggle-group-example-select-checkbox-status-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-status-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-status-warning").
																						Name("toolbar-toggle-group-example-select-checkbox-status-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-status-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-status-restarted").
																						Name("toolbar-toggle-group-example-select-checkbox-status-restarted"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Restarted"),
																						),
																				),
																		),
																),
														),
												),
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Span().
																ID("toolbar-toggle-group-example-select-checkbox-risk-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-toggle-group-example-select-checkbox-risk-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-toggle-group-example-select-checkbox-risk-label toolbar-toggle-group-example-select-checkbox-risk-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Risk"),
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
															app.Div().
																Class("pf-c-select__menu").
																Hidden(true).
																Body(
																	app.FieldSet().
																		Class("pf-c-select__menu-fieldset").
																		Aria("label", "Select input").
																		Body(
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item pf-m-description").
																				For("toolbar-toggle-group-example-select-checkbox-risk-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-risk-active").
																						Name("toolbar-toggle-group-example-select-checkbox-risk-active"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Active"),
																						),
																					app.Span().
																						Class("pf-c-check__description").
																						Body(
																							app.Text("This is a description"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item pf-m-description").
																				For("toolbar-toggle-group-example-select-checkbox-risk-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-risk-canceled").
																						Name("toolbar-toggle-group-example-select-checkbox-risk-canceled"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Canceled"),
																						),
																					app.Span().
																						Class("pf-c-check__description").
																						Body(
																							app.Text("This is a really long description that describes the menu item. This is a really long description that describes the menu item."),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-risk-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-risk-paused").
																						Name("toolbar-toggle-group-example-select-checkbox-risk-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-risk-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-risk-warning").
																						Name("toolbar-toggle-group-example-select-checkbox-risk-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-toggle-group-example-select-checkbox-risk-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-toggle-group-example-select-checkbox-risk-restarted").
																						Name("toolbar-toggle-group-example-select-checkbox-risk-restarted"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Restarted"),
																						),
																				),
																		),
																),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-toolbar__expandable-content pf-m-hidden").
						ID("toolbar-toggle-group-example-expandable-content").
						Hidden(true),
				),
		)
}
