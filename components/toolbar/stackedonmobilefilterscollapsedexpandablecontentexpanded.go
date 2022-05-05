package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StackedOnMobileFiltersCollapsedExpandableContentExpanded struct {
	app.Compo
}

func (c *StackedOnMobileFiltersCollapsedExpandableContentExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-stacked-collapsed-example").
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
												Aria("controls", "toolbar-stacked-collapsed-example-expandable-content").
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
										ID("toolbar-stacked-collapsed-example-icon-button-overflow-menu").
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
																ID("toolbar-stacked-collapsed-example-icon-button-overflow-menu-dropdown-toggle").
																Aria("label", "Overflow menu").
																Aria("expanded", "false").
																Body(
																	app.I().
																		Class("fas fa-ellipsis-v").
																		Aria("hidden", true),
																),
															app.Ul().
																Class("pf-c-dropdown__menu").
																Aria("labelledby", "toolbar-stacked-collapsed-example-icon-button-overflow-menu-dropdown-toggle").
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
						ID("toolbar-stacked-collapsed-example-expandable-content").
						Body(
							app.Div().
								Class("pf-c-toolbar__group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item pf-m-label").
										ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-label").
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
														ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-label toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-toggle").
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-active").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-active"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-canceled").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-canceled"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-paused").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-warning").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-restarted").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-resource-expanded-restarted"),
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
										ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-label").
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
														ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-stacked-collapsed-example-select-checkbox-status-expanded-label toolbar-stacked-collapsed-example-select-checkbox-status-expanded-toggle").
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-active").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-active"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-canceled").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-canceled"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-paused").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-warning").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-restarted").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-status-expanded-restarted"),
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
										ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-label").
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
														ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-toggle").
														Aria("haspopup", true).
														Aria("expanded", "false").
														Aria("labelledby", "toolbar-stacked-collapsed-example-select-checkbox-type-expanded-label toolbar-stacked-collapsed-example-select-checkbox-type-expanded-toggle").
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-active").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-active").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-active"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-canceled").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-canceled").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-canceled"),
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
																		For("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-paused").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-paused").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-paused"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Paused"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-warning").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-warning").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-warning"),
																			app.Span().
																				Class("pf-c-check__label").
																				Body(
																					app.Text("Warning"),
																				),
																		),
																	app.Label().
																		Class("pf-c-check pf-c-select__menu-item").
																		For("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-restarted").
																		Body(
																			app.Input().
																				Class("pf-c-check__input").
																				Type("checkbox").
																				ID("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-restarted").
																				Name("toolbar-stacked-collapsed-example-select-checkbox-type-expanded-restarted"),
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
														For("toolbar-stacked-collapsed-example-bulk-select-toggle-check").
														Body(
															app.Input().
																Type("checkbox").
																ID("toolbar-stacked-collapsed-example-bulk-select-toggle-check").
																Aria("label", "Select all"),
														),
													app.Button().
														Class("pf-c-dropdown__toggle-button").
														Type("button").
														Aria("expanded", "false").
														ID("toolbar-stacked-collapsed-example-bulk-select-toggle-button").
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
														ID("toolbar-stacked-collapsed-example-pagination-options-menu-toggle").
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
														Aria("labelledby", "toolbar-stacked-collapsed-example-pagination-options-menu-toggle").
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
