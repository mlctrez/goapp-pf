package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableThirdLevel struct {
	app.Compo
}

func (c *ExpandableThirdLevel) Render() app.UI {
	return app.Nav().
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
						Class("pf-c-nav__item pf-m-current").
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
						Class("pf-c-nav__item pf-m-expandable").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("expandable-third-level-example-example-1").
								Aria("expanded", "false").
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
							app.Section().
								Class("pf-c-nav__subnav").
								Aria("labelledby", "expandable-third-level-example-example-1").
								Hidden(true).
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
															app.Text("Subnav link 1"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Subnav link 2"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-nav__item pf-m-expandable pf-m-expanded").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("expandable-third-level-example-example-2").
								Aria("expanded", true).
								Body(
									app.Text("Cost management"), app.Span().
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
							app.Section().
								Class("pf-c-nav__subnav").
								Aria("labelledby", "expandable-third-level-example-example-2").
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
															app.Text("Openshift"),
														),
												),
											app.Li().
												Class("pf-c-nav__item pf-m-expandable pf-m-expanded").
												Body(
													app.Button().
														Class("pf-c-nav__link").
														ID("expandable-third-level-example-sub-example-1").
														Aria("expanded", true).
														Body(
															app.Text("Public clouds"), app.Span().
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
													app.Section().
														Class("pf-c-nav__subnav").
														Aria("labelledby", "expandable-third-level-example-sub-example-1").
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
																					app.Text("Amazon Web Services"),
																				),
																		),
																	app.Li().
																		Class("pf-c-nav__item").
																		Body(
																			app.A().
																				Href("#").
																				Class("pf-c-nav__link").
																				Body(
																					app.Text("Microsoft Azure"),
																				),
																		),
																	app.Li().
																		Class("pf-c-nav__item").
																		Body(
																			app.A().
																				Href("#").
																				Class("pf-c-nav__link").
																				Body(
																					app.Text("Google Cloud Services"),
																				),
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
															app.Text("Cost Models"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Cost Explorer"),
														),
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
									app.Text("Support Cases"),
								),
						),
				),
		)
}
