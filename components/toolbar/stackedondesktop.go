package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StackedOnDesktop struct {
	app.Compo
}

func (c *StackedOnDesktop) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-stacked-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group pf-m-toggle-group pf-m-show-on-2xl").
								Body(
									app.Div().
										Class("pf-c-toolbar__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Show filters").
												Aria("expanded", "false").
												Aria("controls", "toolbar-stacked-example-expandable-content").
												Body(
													app.I().
														Class("fas fa-filter").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-toolbar__group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item pf-m-label").
												ID("-select-checkbox-resource-label").
												Aria("hidden", true).
												Body(
													app.Text("Resource"),
												),
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("-select-checkbox-resource-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "-select-checkbox-resource-label -select-checkbox-resource-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Pod"),
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
																				For("-select-checkbox-resource-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-resource-active").
																						Name("-select-checkbox-resource-active"),
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
																				For("-select-checkbox-resource-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-resource-canceled").
																						Name("-select-checkbox-resource-canceled"),
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
																				For("-select-checkbox-resource-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-resource-paused").
																						Name("-select-checkbox-resource-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-resource-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-resource-warning").
																						Name("-select-checkbox-resource-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-resource-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-resource-restarted").
																						Name("-select-checkbox-resource-restarted"),
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
										Class("pf-c-toolbar__group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item pf-m-label").
												ID("-select-checkbox-status-label").
												Aria("hidden", true).
												Body(
													app.Text("Status"),
												),
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("-select-checkbox-status-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "-select-checkbox-status-label -select-checkbox-status-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Running"),
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
																				For("-select-checkbox-status-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-status-active").
																						Name("-select-checkbox-status-active"),
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
																				For("-select-checkbox-status-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-status-canceled").
																						Name("-select-checkbox-status-canceled"),
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
																				For("-select-checkbox-status-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-status-paused").
																						Name("-select-checkbox-status-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-status-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-status-warning").
																						Name("-select-checkbox-status-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-status-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-status-restarted").
																						Name("-select-checkbox-status-restarted"),
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
										Class("pf-c-toolbar__group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item pf-m-label").
												ID("-select-checkbox-type-label").
												Aria("hidden", true).
												Body(
													app.Text("Type"),
												),
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("-select-checkbox-type-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "-select-checkbox-type-label -select-checkbox-type-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("Any"),
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
																				For("-select-checkbox-type-active").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-type-active").
																						Name("-select-checkbox-type-active"),
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
																				For("-select-checkbox-type-canceled").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-type-canceled").
																						Name("-select-checkbox-type-canceled"),
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
																				For("-select-checkbox-type-paused").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-type-paused").
																						Name("-select-checkbox-type-paused"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Paused"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-type-warning").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-type-warning").
																						Name("-select-checkbox-type-warning"),
																					app.Span().
																						Class("pf-c-check__label").
																						Body(
																							app.Text("Warning"),
																						),
																				),
																			app.Label().
																				Class("pf-c-check pf-c-select__menu-item").
																				For("-select-checkbox-type-restarted").
																				Body(
																					app.Input().
																						Class("pf-c-check__input").
																						Type("checkbox").
																						ID("-select-checkbox-type-restarted").
																						Name("-select-checkbox-type-restarted"),
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
										ID("toolbar-stacked-example-icon-button-overflow-menu").
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
										ID("toolbar-stacked-example-overflow-menu").
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
																ID("toolbar-stacked-example-overflow-menu-dropdown-toggle").
																Aria("label", "Overflow menu").
																Aria("expanded", "false").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu").
																Aria("labelledby", "toolbar-stacked-example-overflow-menu-dropdown-toggle").
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
					app.Div().
						Class("pf-c-toolbar__expandable-content pf-m-hidden").
						ID("toolbar-stacked-example-expandable-content").
						Hidden(true),
				),
			app.Hr().
				Class("pf-c-divider"),
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
														For("toolbar-stacked-example-bulk-select-toggle-check").
														Body(
															app.Input().
																Type("checkbox").
																ID("toolbar-stacked-example-bulk-select-toggle-check").
																Aria("label", "Select all"),
														),
													app.Button().
														Class("pf-c-dropdown__toggle-button").
														Type("button").
														Aria("expanded", "false").
														ID("toolbar-stacked-example-bulk-select-toggle-button").
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
								Class("pf-c-toolbar__item pf-m-pagination").
								Body(
									app.Div().
										Class("pf-c-pagination").
										Aria("label", "Element pagination").
										Body(
											app.Div().
												Class("pf-c-pagination__total-items").
												Body(
													app.Text("37 items"),
												),
											app.Div().
												Class("pf-c-options-menu").
												Body(
													app.Button().
														Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
														Type("button").
														ID("toolbar-stacked-example-pagination-options-menu-toggle").
														Aria("haspopup", "listbox").
														Aria("expanded", "false").
														Body(
															app.Span().
																Class("pf-c-options-menu__toggle-text").
																Body(
																	app.B().
																		Body(
																			app.Text("1 - 10"),
																		),
																	app.Text("of"),
																	app.B().
																		Body(
																			app.Text("36"),
																		),
																),
															app.Div().
																Class("pf-c-options-menu__toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
													app.Ul().
														Class("pf-c-options-menu__menu").
														Aria("labelledby", "toolbar-stacked-example-pagination-options-menu-toggle").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-options-menu__menu-item").
																		Type("button").
																		Body(
																			app.Text("5 per page"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-options-menu__menu-item").
																		Type("button").
																		Body(
																			app.Text("10 per page"), app.Div().
																				Class("pf-c-options-menu__menu-item-icon").
																				Body(
																					app.I().
																						Class("fas fa-check").
																						Aria("hidden", true),
																				),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-options-menu__menu-item").
																		Type("button").
																		Body(
																			app.Text("20 per page"),
																		),
																),
														),
												),
											app.Nav().
												Class("pf-c-pagination__nav").
												Aria("label", "Pagination").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-disabled").
														Type("button").
														Aria("label", "Go to first page").
														Aria("disabled", true).
														Body(
															app.I().
																Class("fas fa-angle-double-left").
																Aria("hidden", true),
														),
													app.Button().
														Class("pf-c-button pf-m-plain pf-m-disabled").
														Type("button").
														Aria("label", "Go to previous page").
														Aria("disabled", true).
														Body(
															app.I().
																Class("fas fa-angle-left").
																Aria("hidden", true),
														),
													app.Div().
														Class("pf-c-pagination__nav-page-select").
														Aria("label", "Current page 1 of 4").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Aria("label", "Current page").
																Type("number").
																Min("1").
																Max("4").
																Value("1"),
															app.Span().
																Aria("hidden", true).
																Body(
																	app.Text("of 4"),
																),
														),
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Go to next page").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Go to last page").
														Body(
															app.I().
																Class("fas fa-angle-double-right").
																Aria("hidden", true),
														),
												),
										),
								),
						),
				),
		)
}
