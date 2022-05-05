package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroupAdvancedSearchExpandedWithAutocomplete struct {
	app.Compo
}

func (c *SearchInputGroupAdvancedSearchExpandedWithAutocomplete) Render() app.UI {
	return app.Div().
		Class("ws-example-wrapper").
		Body(
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Div().
						Class("pf-c-text-input-group").
						Body(
							app.Div().
								Class("pf-c-text-input-group__main pf-m-icon").
								Body(
									app.Span().
										Class("pf-c-text-input-group__text").
										Body(
											app.Span().
												Class("pf-c-text-input-group__icon").
												Body(
													app.I().
														Class("fas fa-fw fa-search"),
												),
											app.Input().
												Class("pf-c-text-input-group__text-input").
												Type("text").
												Value("username:root firstname:n").
												Aria("label", "Type to filter"),
										),
								),
							app.Div().
								Class("pf-c-text-input-group__utilities").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Clear input").
										Body(
											app.I().
												Class("fas fa-times fa-fw").
												Aria("hidden", true),
										),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-control pf-m-expanded").
						Type("button").
						Aria("expanded", true).
						Aria("label", "Advanced search").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("submit").
						Aria("label", "Search").
						Body(
							app.I().
								Class("fas fa-arrow-right").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-panel pf-m-raised").
				Body(
					app.Div().
						Class("pf-c-panel__main").
						Body(
							app.Div().
								Class("pf-c-panel__main-body").
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
																For("text-input-group-advanced-search-input-form-with-autocomplete-example-username").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Username"),
																		),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																Value("root").
																ID("text-input-group-advanced-search-input-form-with-autocomplete-example-username").
																Name("text-input-group-advanced-search-input-form-with-autocomplete-example-username"),
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
																For("text-input-group-advanced-search-input-form-with-autocomplete-example-firstname").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("First name"),
																		),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																Value("ned").
																ID("text-input-group-advanced-search-input-form-with-autocomplete-example-firstname").
																Name("text-input-group-advanced-search-input-form-with-autocomplete-example-firstname"),
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
																For("text-input-group-advanced-search-input-form-with-autocomplete-example-group").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Group"),
																		),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("text-input-group-advanced-search-input-form-with-autocomplete-example-group").
																Name("text-input-group-advanced-search-input-form-with-autocomplete-example-group"),
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
																For("text-input-group-advanced-search-input-form-with-autocomplete-example-email").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Email"),
																		),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("text-input-group-advanced-search-input-form-with-autocomplete-example-email").
																Name("text-input-group-advanced-search-input-form-with-autocomplete-example-email"),
														),
												),
											app.Div().
												Class("pf-c-form__group pf-m-action").
												Body(
													app.Div().
														Class("pf-c-form__actions").
														Body(
															app.Button().
																Class("pf-c-button pf-m-primary").
																Type("submit").
																Body(
																	app.Text("Submit"),
																),
															app.Button().
																Class("pf-c-button pf-m-link").
																Type("reset").
																Body(
																	app.Text("Reset"),
																),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-menu").
				Body(
					app.Div().
						Class("pf-c-menu__content").
						Body(
							app.Ul().
								Class("pf-c-menu__list").
								Aria("role", "menu").
								Body(
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("nancy"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("ned"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("neil"),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.Button().
												Class("pf-c-menu__item").
												Type("button").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("nicole"),
																),
														),
												),
										),
								),
						),
				),
		)
}
