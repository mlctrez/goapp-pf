package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandedWithSubnavTitles struct {
	app.Compo
}

func (c *ExpandedWithSubnavTitles) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav").
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
								Aria("expanded", true).
								Body(
									app.Text("Link 1"), app.Span().
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
								Aria("labelledby", "subnav-title1").
								Body(
									app.H2().
										Class("pf-c-nav__subnav-title pf-screen-reader").
										ID("subnav-title1").
										Body(
											app.Text("Current and expanded example sub-navigation"),
										),
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
					app.Li().
						Class("pf-c-nav__item pf-m-expandable pf-m-expanded").
						Body(
							app.Button().
								Class("pf-c-nav__link").
								Aria("expanded", true).
								Body(
									app.Text("Link 2"), app.Span().
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
								Aria("labelledby", "subnav-title2").
								Body(
									app.H2().
										Class("pf-c-nav__subnav-title pf-screen-reader").
										ID("subnav-title2").
										Body(
											app.Text("Expanded, but not current example sub-navigation"),
										),
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
