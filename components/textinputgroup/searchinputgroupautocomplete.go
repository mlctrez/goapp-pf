package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroupAutocomplete struct {
	app.Compo
}

func (c *SearchInputGroupAutocomplete) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
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
										Value("app").
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
																	app.Text("apple"),
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
																	app.Text("appleby"),
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
																	app.Text("appleseed"),
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
																	app.Text("appleton"),
																),
														),
												),
										),
								),
						),
				),
		)
}
