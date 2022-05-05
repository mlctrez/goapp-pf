package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type UsingTheNavElement struct {
	app.Compo
}

func (c *UsingTheNavElement) Render() app.UI {
	return app.Nav().
		Class("pf-c-tabs pf-m-scrollable").
		Aria("label", "Local").
		ID("default-scroll-nav-example").
		Body(
			app.Button().
				Class("pf-c-tabs__scroll-button").
				Aria("label", "Scroll left").
				Body(
					app.I().
						Class("fas fa-angle-left").
						Aria("hidden", true),
				),
			app.Ul().
				Class("pf-c-tabs__list").
				Body(
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.A().
								Class("pf-c-tabs__link").
								Href("#").
								ID("default-scroll-nav-example-users-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Users"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item pf-m-current").
						Body(
							app.A().
								Class("pf-c-tabs__link").
								Href("#").
								ID("default-scroll-nav-example-containers-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Containers"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.A().
								Class("pf-c-tabs__link").
								Href("#").
								ID("default-scroll-nav-example-database-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Database"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.A().
								Class("pf-c-tabs__link pf-m-disabled").
								Aria("disabled", true).
								TabIndex(-1).
								Href("#").
								ID("default-scroll-nav-example-disabled-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Disabled"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.A().
								Class("pf-c-tabs__link pf-m-aria-disabled").
								Aria("disabled", true).
								Href("#").
								ID("default-scroll-nav-example-aria-disabled-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("ARIA disabled"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.A().
								Class("pf-c-tabs__link").
								Href("#").
								ID("default-scroll-nav-example-network-wired-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Network"),
										),
								),
						),
				),
			app.Button().
				Class("pf-c-tabs__scroll-button").
				Aria("label", "Scroll right").
				Body(
					app.I().
						Class("fas fa-angle-right").
						Aria("hidden", true),
				),
		)
}
