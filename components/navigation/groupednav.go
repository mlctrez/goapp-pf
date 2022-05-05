package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupedNav struct {
	app.Compo
}

func (c *GroupedNav) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav").
		Aria("label", "Global").
		Body(
			app.Section().
				Class("pf-c-nav__section").
				Aria("labelledby", "grouped-title1").
				Body(
					app.H2().
						Class("pf-c-nav__section-title").
						ID("grouped-title1").
						Body(
							app.Text("Section title 1"),
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
											app.Text("Link 1"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Link 2"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Link 3"),
										),
								),
						),
				),
			app.Section().
				Class("pf-c-nav__section").
				Aria("labelledby", "grouped-title2").
				Body(
					app.H2().
						Class("pf-c-nav__section-title").
						ID("grouped-title2").
						Body(
							app.Text("Section title 2"),
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
											app.Text("Link 1"),
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
											app.Text("Current link"),
										),
								),
							app.Li().
								Class("pf-c-nav__item").
								Body(
									app.A().
										Href("#").
										Class("pf-c-nav__link").
										Body(
											app.Text("Link 3"),
										),
								),
						),
				),
		)
}
