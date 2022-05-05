package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectedFiltersOnMobileFiltersCollapsedExpandableContentExpanded struct {
	app.Compo
}

func (c *SelectedFiltersOnMobileFiltersCollapsedExpandableContentExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-selected-filters-toggle-group-expanded-example").
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
														For("toolbar-selected-filters-toggle-group-expanded-example-bulk-select-toggle-check").
														Body(
															app.Input().
																Type("checkbox").
																ID("toolbar-selected-filters-toggle-group-expanded-example-bulk-select-toggle-check").
																Aria("label", "Select all"),
														),
													app.Button().
														Class("pf-c-dropdown__toggle-button").
														Type("button").
														Aria("expanded", "false").
														ID("toolbar-selected-filters-toggle-group-expanded-example-bulk-select-toggle-button").
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
												Aria("controls", "toolbar-selected-filters-toggle-group-expanded-example-expandable-content").
												Body(
													app.I().
														Class("fas fa-filter").
														Aria("hidden", true),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-overflow-menu").
								Body(
									app.Div().
										Class("pf-c-overflow-menu").
										ID("toolbar-selected-filters-toggle-group-expanded-example-icon-button-overflow-menu").
										Body(
											app.Div().
												Class("pf-c-overflow-menu__control").
												Body(
													app.Div().
														Class("pf-c-dropdown").
														Body(
															app.Button().
																Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																Type("button").
																ID("toolbar-selected-filters-toggle-group-expanded-example-icon-button-overflow-menu-dropdown-toggle").
																Aria("label", "Overflow menu").
																Aria("expanded", "false").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu").
																Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-icon-button-overflow-menu-dropdown-toggle").
																Hidden(true).
																Body(
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Edit"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Clone"),
																				),
																		),
																	app.Li().
																		Body(
																			app.Button().
																				Class("pf-c-dropdown__menu-item").
																				Body(
																					app.Text("Sync"),
																				),
																		),
																),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-toolbar__expandable-content pf-m-expanded").
						ID("toolbar-selected-filters-toggle-group-expanded-example-expandable-content").
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
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-name-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-name-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-select-name-label toolbar-selected-filters-toggle-group-expanded-example-select-name-toggle").
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
														Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-select-name-label").
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
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-label toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-toggle").
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-active").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-active"),
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-canceled").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-canceled"),
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-paused").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-warning").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-restarted").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-status-expanded-restarted"),
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
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-label").
														Hidden(true).
														Body(
															app.Text("Choose one"),
														),
													app.Button().
														Class("pf-c-select__toggle").
														Type("button").
														ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-label toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-toggle").
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-active").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-active"),
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-canceled").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-canceled"),
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
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-paused").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-warning").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-restarted").
																				Name("toolbar-selected-filters-toggle-group-expanded-example-select-checkbox-risk-expanded-restarted"),
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
							app.Div().
								Class("pf-c-toolbar__group pf-m-chip-container").
								Body(
									app.Div().
										Class("pf-c-toolbar__group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item pf-m-chip-group").
												Body(
													app.Div().
														Class("pf-c-chip-group pf-m-category").
														Body(
															app.Div().
																Class("pf-c-chip-group__main").
																Body(
																	app.Span().
																		Class("pf-c-chip-group__label").
																		Aria("hidden", true).
																		ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-status-chip-group-label").
																		Body(
																			app.Text("Status"),
																		),
																	app.Ul().
																		Class("pf-c-chip-group__list").
																		Aria("role", "list").
																		Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-status-chip-group-label").
																		Body(
																			app.Li().
																				Class("pf-c-chip-group__list-item").
																				Body(
																					app.Div().
																						Class("pf-c-chip").
																						Body(
																							app.Span().
																								Class("pf-c-chip__text").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-one").
																								Body(
																									app.Text("Chip one"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-one toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-one").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-one").
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
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-two").
																								Body(
																									app.Text("Chip two"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-two toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-two").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-two").
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
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-three").
																								Body(
																									app.Text("Chip three"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-three toolbar-selected-filters-toggle-group-expanded-example-chip-group-statuschip-three").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-statusremove-chip-three").
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
												),
											app.Div().
												Class("pf-c-toolbar__item pf-m-chip-group").
												Body(
													app.Div().
														Class("pf-c-chip-group pf-m-category").
														Body(
															app.Div().
																Class("pf-c-chip-group__main").
																Body(
																	app.Span().
																		Class("pf-c-chip-group__label").
																		Aria("hidden", true).
																		ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-risk-chip-group-label").
																		Body(
																			app.Text("Risk"),
																		),
																	app.Ul().
																		Class("pf-c-chip-group__list").
																		Aria("role", "list").
																		Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-risk-chip-group-label").
																		Body(
																			app.Li().
																				Class("pf-c-chip-group__list-item").
																				Body(
																					app.Div().
																						Class("pf-c-chip").
																						Body(
																							app.Span().
																								Class("pf-c-chip__text").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-one").
																								Body(
																									app.Text("Chip one"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-one toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-one").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-one").
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
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-two").
																								Body(
																									app.Text("Chip two"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-two toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-two").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-two").
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
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-three").
																								Body(
																									app.Text("Chip three"),
																								),
																							app.Button().
																								Class("pf-c-button pf-m-plain").
																								Type("button").
																								Aria("labelledby", "toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-three toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskchip-three").
																								Aria("label", "Remove").
																								ID("toolbar-selected-filters-toggle-group-expanded-example-chip-group-riskremove-chip-three").
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
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-link pf-m-inline").
												Type("button").
												Body(
													app.Text("Clear all filters"),
												),
										),
								),
						),
				),
		)
}
