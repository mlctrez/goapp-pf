package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Mixed struct {
	app.Compo
}

func (c *Mixed) Render() app.UI {
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
									app.Text("Link 1 (not expandable)"),
								),
						),
					app.Li().
						Class("pf-c-nav__item pf-m-expandable pf-m-expanded").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("nav-mixed-link2").
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
								Aria("labelledby", "nav-mixed-link2").
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
						Class("pf-c-nav__item pf-m-expandable pf-m-current").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								ID("nav-mixed-link4").
								Aria("expanded", "false").
								Body(
									app.Text("Link 3 (current, but not expanded example)"), app.Span().
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
								Aria("labelledby", "nav-mixed-link4").
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
														Class("pf-c-nav__link pf-m-current").
														Aria("current", "page").
														Body(
															app.Text("Subnav link 2"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Subnav link 3"),
														),
												),
										),
								),
						),
				),
		)
}
