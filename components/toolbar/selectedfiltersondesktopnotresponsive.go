package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectedFiltersOnDesktopNotResponsive struct {
	app.Compo
}

func (c *SelectedFiltersOnDesktopNotResponsive) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-selected-filters-example").
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
														For("toolbar-selected-filters-example-bulk-select-toggle-check").
														Body(
															app.Input().
																Type("checkbox").
																ID("toolbar-selected-filters-example-bulk-select-toggle-check").
																Aria("label", "Select all"),
														),
													app.Button().
														Class("pf-c-dropdown__toggle-button").
														Type("button").
														Aria("expanded", "false").
														ID("toolbar-selected-filters-example-bulk-select-toggle-button").
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
								Class("pf-c-toolbar__group pf-m-toggle-group pf-m-show").
								Body(
									app.Div().
										Class("pf-c-toolbar__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Show filters").
												Aria("expanded", "false").
												Aria("controls", "toolbar-selected-filters-example-expandable-content").
												Body(
													app.I().
														Class("fas fa-filter").
														Aria("hidden", true),
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
																ID("toolbar-selected-filters-example-select-checkbox-status-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-selected-filters-example-select-checkbox-status-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-selected-filters-example-select-checkbox-status-label toolbar-selected-filters-example-select-checkbox-status-toggle").
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
																				For("toolbar-selected-filters-example-select-checkbox-status-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-status-active").
																						Name("toolbar-selected-filters-example-select-checkbox-status-active"),
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
																				For("toolbar-selected-filters-example-select-checkbox-status-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-status-canceled").
																						Name("toolbar-selected-filters-example-select-checkbox-status-canceled"),
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
																				For("toolbar-selected-filters-example-select-checkbox-status-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-status-paused").
																						Name("toolbar-selected-filters-example-select-checkbox-status-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-selected-filters-example-select-checkbox-status-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-status-warning").
																						Name("toolbar-selected-filters-example-select-checkbox-status-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-selected-filters-example-select-checkbox-status-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-status-restarted").
																						Name("toolbar-selected-filters-example-select-checkbox-status-restarted"),
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
																ID("toolbar-selected-filters-example-select-checkbox-risk-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("toolbar-selected-filters-example-select-checkbox-risk-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "toolbar-selected-filters-example-select-checkbox-risk-label toolbar-selected-filters-example-select-checkbox-risk-toggle").
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
																				For("toolbar-selected-filters-example-select-checkbox-risk-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-risk-active").
																						Name("toolbar-selected-filters-example-select-checkbox-risk-active"),
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
																				For("toolbar-selected-filters-example-select-checkbox-risk-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-risk-canceled").
																						Name("toolbar-selected-filters-example-select-checkbox-risk-canceled"),
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
																				For("toolbar-selected-filters-example-select-checkbox-risk-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-risk-paused").
																						Name("toolbar-selected-filters-example-select-checkbox-risk-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-selected-filters-example-select-checkbox-risk-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-risk-warning").
																						Name("toolbar-selected-filters-example-select-checkbox-risk-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("toolbar-selected-filters-example-select-checkbox-risk-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("toolbar-selected-filters-example-select-checkbox-risk-restarted").
																						Name("toolbar-selected-filters-example-select-checkbox-risk-restarted"),
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
										ID("toolbar-selected-filters-example-icon-button-overflow-menu").
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
										ID("toolbar-selected-filters-example-overflow-menu").
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
														Class("pf-c-dropdown").
														Body(
															app.Button().
																Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																Type("button").
																ID("toolbar-selected-filters-example-overflow-menu-dropdown-toggle").
																Aria("label", "Overflow menu").
																Aria("expanded", "false").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu").
																Aria("labelledby", "toolbar-selected-filters-example-overflow-menu-dropdown-toggle").
																Hidden(true).
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
				),
			app.Div().
				Class("pf-c-toolbar__content pf-m-chip-container").
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
														ID("toolbar-selected-filters-example-chip-group-status-chip-group-label").
														Body(
															app.Text("Status"),
														),
													app.Ul().
														Class("pf-c-chip-group__list").
														Aria("role", "list").
														Aria("labelledby", "toolbar-selected-filters-example-chip-group-status-chip-group-label").
														Body(
															app.Li().
																Class("pf-c-chip-group__list-item").
																Body(
																	app.Div().
																		Class("pf-c-chip").
																		Body(
																			app.Span().
																				Class("pf-c-chip__text").
																				ID("toolbar-selected-filters-example-chip-group-statuschip-one").
																				Body(
																					app.Text("Chip one"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-statusremove-chip-one toolbar-selected-filters-example-chip-group-statuschip-one").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-statusremove-chip-one").
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
																				ID("toolbar-selected-filters-example-chip-group-statuschip-two").
																				Body(
																					app.Text("Chip two"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-statusremove-chip-two toolbar-selected-filters-example-chip-group-statuschip-two").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-statusremove-chip-two").
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
																				ID("toolbar-selected-filters-example-chip-group-statuschip-three").
																				Body(
																					app.Text("Chip three"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-statusremove-chip-three toolbar-selected-filters-example-chip-group-statuschip-three").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-statusremove-chip-three").
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
														ID("toolbar-selected-filters-example-chip-group-risk-chip-group-label").
														Body(
															app.Text("Risk"),
														),
													app.Ul().
														Class("pf-c-chip-group__list").
														Aria("role", "list").
														Aria("labelledby", "toolbar-selected-filters-example-chip-group-risk-chip-group-label").
														Body(
															app.Li().
																Class("pf-c-chip-group__list-item").
																Body(
																	app.Div().
																		Class("pf-c-chip").
																		Body(
																			app.Span().
																				Class("pf-c-chip__text").
																				ID("toolbar-selected-filters-example-chip-group-riskchip-one").
																				Body(
																					app.Text("Chip one"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-riskremove-chip-one toolbar-selected-filters-example-chip-group-riskchip-one").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-riskremove-chip-one").
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
																				ID("toolbar-selected-filters-example-chip-group-riskchip-two").
																				Body(
																					app.Text("Chip two"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-riskremove-chip-two toolbar-selected-filters-example-chip-group-riskchip-two").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-riskremove-chip-two").
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
																				ID("toolbar-selected-filters-example-chip-group-riskchip-three").
																				Body(
																					app.Text("Chip three"),
																				),
																			app.Button().
																				Class("pf-c-button pf-m-plain").
																				Type("button").
																				Aria("labelledby", "toolbar-selected-filters-example-chip-group-riskremove-chip-three toolbar-selected-filters-example-chip-group-riskchip-three").
																				Aria("label", "Remove").
																				ID("toolbar-selected-filters-example-chip-group-riskremove-chip-three").
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
		)
}
