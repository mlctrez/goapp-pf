package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NavWithFlyout struct {
	app.Compo
}

func (c *NavWithFlyout) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Nav().
				Class("pf-c-nav").
				Aria("label", "Global").
				Body(
					app.Ul().
						Class("pf-c-nav__list").
						Body(
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Clusters"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Overview"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Releases"),
										),
								),
							app.Li().
								Class("pf-c-nav__item pf-m-flyout").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link pf-m-hover").
										Aria("haspopup", true).
										Aria("expanded", true).
										Body(
											app.Text("Subscriptions"), app.Span().
												Class("pf-c-nav__toggle").
												Body(
													app.Span().
														Class("pf-c-nav__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Support cases"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Cluster manager feedback"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Red Hat Marketplace"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Documentation"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-menu pf-m-flyout pf-m-nav").
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
																	app.Text("Container platform"),
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
												Aria("expanded", true).
												Body(
													app.Span().
														Class("pf-c-menu__item-main").
														Body(
															app.Span().
																Class("pf-c-menu__item-text").
																Body(
																	app.Text("Dedicated"),
																),
															app.Span().
																Class("pf-c-menu__item-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right"),
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
																									app.Text("Dedicated (Annual)"),
																								),
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
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-main").
																						Body(
																							app.Span().
																								Class("pf-c-menu__item-text").
																								Body(
																									app.Text("Dedicated (On-Demand)"),
																								),
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
																				Body(
																					app.Span().
																						Class("pf-c-menu__item-main").
																						Body(
																							app.Span().
																								Class("pf-c-menu__item-text").
																								Body(
																									app.Text("Dedicated (On-Demand limits)"),
																								),
																						),
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
