package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupedNavNoTitlesNoMarginTop struct {
	app.Compo
}

func (c *GroupedNavNoTitlesNoMarginTop) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav").
		Aria("label", "Global").
		Body(
			app.Section().
				Class("pf-c-nav__section pf-m-no-title").
				Aria("label", "Section one").
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
				Class("pf-c-nav__section pf-m-no-title").
				Aria("label", "Section two").
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
											app.Text("Section 2, link 1"),
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
