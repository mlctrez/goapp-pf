package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontalSubnav struct {
	app.Compo
}

func (c *HorizontalSubnav) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav pf-m-horizontal-subnav").
		Aria("label", "Local").
		Body(
			app.Ul().
				Class("pf-c-nav__list").
				Body(
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link pf-m-current").
								Aria("current", "page").
								Body(
									app.Text("Item 1"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Item 2"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Item 3"),
								),
						),
				),
		)
}
