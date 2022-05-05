package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Default struct {
	app.Compo
}

func (c *Default) Render() app.UI {
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
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Link 4"),
								),
						),
				),
		)
}
