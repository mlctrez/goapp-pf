package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandedInLightMode struct {
	app.Compo
}

func (c *ExpandedInLightMode) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav pf-m-light").
		Aria("label", "Global").
		Body(
			app.Ul().
				Class("pf-c-nav__list").
				Body(
					app.Li().
						Class("pf-c-nav__item pf-m-expandable pf-m-expanded pf-m-current").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("expandable-example1").
								Aria("expanded", true).
								Body(
									app.Text("Link 1 (current and expanded example)"), app.Span().
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
								Aria("labelledby", "expandable-example1").
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
															app.Text("Current link"),
														),
												),
											app.Li().
												Class("pf-c-divider").
												Aria("role", "separator"),
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
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link pf-m-current").
														Aria("current", "page").
														Body(
															app.Text("Subnav link 3"),
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
								ID("expandable-example2").
								Aria("expanded", true).
								Body(
									app.Text("Link 2 (expanded, but not current example)"), app.Span().
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
								Aria("labelledby", "expandable-example2").
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
						Class("pf-c-nav__item pf-m-expandable").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("expandable-example3").
								Aria("expanded", "false").
								Body(
									app.Text("Link 3"), app.Span().
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
								Aria("labelledby", "expandable-example3").
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
				),
		)
}
