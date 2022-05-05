package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithFiltering struct {
	app.Compo
}

func (c *WithFiltering) Render() app.UI {
	return app.Div().
		Class("pf-c-menu").
		Body(
			app.Div().
				Class("pf-c-menu__search").
				Body(
					app.Div().
						Class("pf-c-menu__search-input").
						Body(
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
														Aria("label", "Search"),
												),
										),
								),
						),
				),
			app.Hr().
				Class("pf-c-divider"),
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
															app.Text("Action 1"),
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
															app.Text("Action 2"),
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
															app.Text("Action 3"),
														),
												),
										),
								),
						),
				),
		)
}
