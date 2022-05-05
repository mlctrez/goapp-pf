package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ToggleGroupOnMobileFiltersCollapsedExpandableContentExpanded struct {
	app.Compo
}

func (c *ToggleGroupOnMobileFiltersCollapsedExpandableContentExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-toggle-group-collapsed-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group pf-m-toggle-group").
								Body(
									app.Div().
										Class("pf-c-toolbar__toggle pf-m-expanded").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Show filters").
												Aria("expanded", true).
												Aria("controls", "toolbar-toggle-group-collapsed-example-expandable-content").
												Body(
													app.I().
														Class("fas fa-filter").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-toolbar__expandable-content pf-m-expanded").
						ID("toolbar-toggle-group-collapsed-example-expandable-content").
						Body(
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
														ID("toolbar-toggle-group-collapsed-example-select-name-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-toggle-group-collapsed-example-select-name-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-toggle-group-collapsed-example-select-name-label toolbar-toggle-group-collapsed-example-select-name-toggle").
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
														Aria("labelledby", "toolbar-toggle-group-collapsed-example-select-name-label").
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
														ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-label toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-toggle").
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-active").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-active"),
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-canceled").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-canceled"),
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-paused").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-warning").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-restarted").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-status-expanded-restarted"),
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
														ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-label toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-toggle").
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-active").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-active"),
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-canceled").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-canceled"),
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
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-paused").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-warning").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-restarted").
																				Name("toolbar-toggle-group-collapsed-example-select-checkbox-risk-expanded-restarted"),
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
		)
}
