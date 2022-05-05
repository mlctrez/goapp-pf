package navigation

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LegacyTertiaryOverflow struct {
	app.Compo
}

func (c *LegacyTertiaryOverflow) Render() app.UI {
	return app.Nav().
		Class("pf-c-nav pf-m-tertiary pf-m-scrollable").
		Aria("label", "Local").
		Body(
			app.Button().
				Class("pf-c-nav__scroll-button").
				Disabled(true).
				Aria("label", "Scroll left").
				Body(
					app.I().
						Class("fas fa-angle-left").
						Aria("hidden", true),
				),
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
									app.Text("Tertiary nav item 1"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Tertiary nav item 2"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Tertiary nav item 3"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Tertiary nav item 4"),
								),
						),
					app.Li().
						Class("pf-c-nav__item").
						Body(
							app.A().
								Href("#").
								Class("pf-c-nav__link").
								Body(
									app.Text("Tertiary nav item 5"),
								),
						),
				),
			app.Button().
				Class("pf-c-nav__scroll-button").
				Aria("label", "Scroll right").
				Body(
					app.I().
						Class("fas fa-angle-right").
						Aria("hidden", true),
				),
		)
}
