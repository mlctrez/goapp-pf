package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Inset struct {
	app.Compo
}

func (c *Inset) Render() app.UI {
	return app.Div().
		Class("pf-c-tabs pf-m-inset-sm-on-md pf-m-inset-lg-on-lg pf-m-inset-2xl-on-xl").
		ID("inset-example").
		Body(
			app.Button().
				Class("pf-c-tabs__scroll-button").
				Disabled(true).
				Aria("hidden", true).
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
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-users-link").
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
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-containers-link").
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
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-database-link").
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
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-server-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Server"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-system-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("System"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("inset-example-network-wired-link").
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
				Disabled(true).
				Aria("hidden", true).
				Aria("label", "Scroll right").
				Body(
					app.I().
						Class("fas fa-angle-right").
						Aria("hidden", true),
				),
		)
}
