package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandedElements struct {
	app.Compo
}

func (c *ExpandedElements) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-expanded-elements-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__item pf-m-bulk-select").
								Body(
									app.Div().
										Class("pf-c-dropdown").
										Body(
											app.Div().
												Class("pf-c-dropdown__toggle pf-m-split-button").
												Body(
													app.Label().
														Class("pf-c-dropdown__toggle-check").
														For("toolbar-expanded-elements-example-bulk-select-toggle-check").
														Body(
															app.Input().
																Type("checkbox").
																ID("toolbar-expanded-elements-example-bulk-select-toggle-check").
																Aria("label", "Select all"),
														),
													app.Button().
														Class("pf-c-dropdown__toggle-button").
														Type("button").
														Aria("expanded", "false").
														ID("toolbar-expanded-elements-example-bulk-select-toggle-button").
														Aria("label", "Dropdown toggle").
														Body(
															app.I().
																Class("fas fa-caret-down").
																Aria("hidden", true),
														),
												),
											app.Ul().
												Class("pf-c-dropdown__menu").
												Hidden(true).
												Body(
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Select all"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Select none"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Other action"),
																),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__group pf-m-toggle-group pf-m-show-on-xl").
								Body(
									app.Div().
										Class("pf-c-toolbar__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Show filters").
												Aria("expanded", "false").
												Aria("controls", "toolbar-expanded-elements-example-expandable-content").
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
																ID("toolbar-expanded-elements-example-select-name-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-expanded-elements-example-select-name-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-expanded-elements-example-select-name-label toolbar-expanded-elements-example-select-name-toggle").
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
																Aria("labelledby", "toolbar-expanded-elements-example-select-name-label").
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
														Class("pf-c-select pf-m-expanded").
														Body(
															app.Span().
																ID("toolbar-expanded-elements-example-select-checkbox-status-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-expanded-elements-example-select-checkbox-status-toggle").
																Aria("haspopup", true).
																Aria("expanded", true).
																Aria("labelledby", "toolbar-expanded-elements-example-select-checkbox-status-label toolbar-expanded-elements-example-select-checkbox-status-toggle").
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
																Body(
																	app.FieldSet().
																		Class("pf-c-select__menu-fieldset").
																		Aria("label", "Select input").
																		Body(
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-status-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-status-active").
																						Name("toolbar-expanded-elements-example-select-checkbox-status-active"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Active"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-status-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-status-canceled").
																						Name("toolbar-expanded-elements-example-select-checkbox-status-canceled").
																						Checked(true),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Canceled"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-status-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-status-paused").
																						Name("toolbar-expanded-elements-example-select-checkbox-status-paused").
																						Checked(true),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-status-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-status-warning").
																						Name("toolbar-expanded-elements-example-select-checkbox-status-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-status-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-status-restarted").
																						Name("toolbar-expanded-elements-example-select-checkbox-status-restarted").
																						Checked(true),
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
														Class("pf-c-select pf-m-expanded").
														Body(
															app.Span().
																ID("toolbar-expanded-elements-example-select-checkbox-risk-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-expanded-elements-example-select-checkbox-risk-toggle").
																Aria("haspopup", true).
																Aria("expanded", true).
																Aria("labelledby", "toolbar-expanded-elements-example-select-checkbox-risk-label toolbar-expanded-elements-example-select-checkbox-risk-toggle").
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
																Body(
																	app.FieldSet().
																		Class("pf-c-select__menu-fieldset").
																		Aria("label", "Select input").
																		Body(
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-risk-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-risk-active").
																						Name("toolbar-expanded-elements-example-select-checkbox-risk-active"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Active"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-risk-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-risk-canceled").
																						Name("toolbar-expanded-elements-example-select-checkbox-risk-canceled").
																						Checked(true),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Canceled"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-risk-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-risk-paused").
																						Name("toolbar-expanded-elements-example-select-checkbox-risk-paused").
																						Checked(true),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-risk-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-risk-warning").
																						Name("toolbar-expanded-elements-example-select-checkbox-risk-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-expanded-elements-example-select-checkbox-risk-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-expanded-elements-example-select-checkbox-risk-restarted").
																						Name("toolbar-expanded-elements-example-select-checkbox-risk-restarted").
																						Checked(true),
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
							app.Div().
								Class("pf-c-toolbar__item pf-m-overflow-menu").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("toolbar-expanded-elements-example-icon-button-overflow-menu").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__content").
												Body(
													app.Div().
														Class("pf-c-overflow-menu__group pf-m-icon-button-group").
														Body(
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Edit").
																		Body(
																			app.I().
																				Class("fas fa-edit").
																				Aria("hidden", true),
																		),
																),
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Clone").
																		Body(
																			app.I().
																				Class("fas fa-clone").
																				Aria("hidden", true),
																		),
																),
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Sync").
																		Body(
																			app.I().
																				Class("fas fa-sync").
																				Aria("hidden", true),
																		),
																),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-overflow-menu").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("toolbar-expanded-elements-example-overflow-menu").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__content").
												Body(
													app.Div().
														Class("pf-c-overflow-menu__group pf-m-button-group").
														Body(
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-primary").
																		Type("button").
																		Body(
																			app.Text("Primary"),
																		),
																),
															app.Div().
																Class("pf-c-overflow-menu__item").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-secondary").
																		Type("button").
																		Body(
																			app.Text("Secondary"),
																		),
																),
														),
												),
											app.Div().
												Class("pf-c-overflow-menu__control").
												Body(
													app.Div().
														Class("pf-c-dropdown pf-m-expanded").
														Body(
															app.Button().
																Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																Type("button").
																ID("toolbar-expanded-elements-example-overflow-menu-dropdown-toggle").
																Aria("label", "Overflow menu").
																Aria("expanded", true).
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu").
																Aria("labelledby", "toolbar-expanded-elements-example-overflow-menu-dropdown-toggle").
																Body(
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Tertiary"),
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
						ID("toolbar-expanded-elements-example-expandable-content").
						Hidden(true),
				),
		)
}
