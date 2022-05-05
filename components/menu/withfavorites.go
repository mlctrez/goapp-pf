package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithFavorites struct {
	app.Compo
}

func (c *WithFavorites) Render() app.UI {
	return app.Div().
		Class("pf-c-menu").
		Body(
			app.Div().
				Class("pf-c-menu__content").
				Body(
					app.Section().
						Class("pf-c-menu__group").
						Body(
							app.H1().
								Class("pf-c-menu__group-title").
								Body(
									app.Text("Favorites"),
								),
							app.Ul().
								Class("pf-c-menu__list").
								Aria("role", "menu").
								Body(
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.A().
												Class("pf-c-menu__item").
												Href("#").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 1"),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action pf-m-favorite pf-m-favorited").
												Type("button").
												Aria("label", "Starred").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.A().
												Class("pf-c-menu__item").
												Href("#").
												Aria("role", "menuitem").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 3"),
																),
															app.Span().
																Class("pf-c-menu__item-external-icon").
																Body(
																	app.I().
																		Class("fas fa-external-link-alt").
																		Aria("hidden", true),
																),
															app.Span().
																Class("pf-screen-reader").
																Body(
																	app.Text("(opens new window)"),
																),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action pf-m-favorite").
												Type("button").
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
					app.Section().
						Class("pf-c-menu__group").
						Body(
							app.H1().
								Class("pf-c-menu__group-title").
								Body(
									app.Text("All actions"),
								),
							app.Ul().
								Class("pf-c-menu__list").
								Aria("role", "menu").
								Body(
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.A().
												Class("pf-c-menu__item").
												Href("#").
												Aria("role", "menuitem").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 1"),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action pf-m-favorite").
												Type("button").
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item pf-m-disabled").
										Aria("role", "none").
										Body(
											app.A().
												Class("pf-c-menu__item").
												Href("#").
												Aria("disabled", true).
												TabIndex(-1).
												Aria("role", "menuitem").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 2 disabled"),
																),
															app.Span().
																Class("pf-c-menu__item-external-icon").
																Body(
																	app.I().
																		Class("fas fa-external-link-alt").
																		Aria("hidden", true),
																),
															app.Span().
																Class("pf-screen-reader").
																Body(
																	app.Text("(opens new window)"),
																),
														),
													app.Span().
														Class("pf-c-menu__item-description").
														Body(
															app.Text("This is a description"),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action pf-m-favorite").
												Type("button").
												Disabled(true).
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-menu__list-item").
										Aria("role", "none").
										Body(
											app.A().
												Class("pf-c-menu__item").
												Href("#").
												Aria("role", "menuitem").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Item 3"),
																),
															app.Span().
																Class("pf-c-menu__item-external-icon").
																Body(
																	app.I().
																		Class("fas fa-external-link-alt").
																		Aria("hidden", true),
																),
															app.Span().
																Class("pf-screen-reader").
																Body(
																	app.Text("(opens new window)"),
																),
														),
												),
											app.Button().
												Class("pf-c-menu__item-action pf-m-favorite").
												Type("button").
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-menu__item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
								),
						),
				),
		)
}
