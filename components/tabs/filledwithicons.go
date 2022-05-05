package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FilledWithIcons struct {
	app.Compo
}

func (c *FilledWithIcons) Render() app.UI {
	return app.Div().
		Class("pf-c-tabs pf-m-fill").
		ID("filled-with-icons-example").
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
								ID("filled-with-icons-example-users-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-icon").
										Body(
											app.I().
												Class("fas fa-fas fa-users").
												Aria("hidden", true),
										),
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
								ID("filled-with-icons-example-containers-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-icon").
										Body(
											app.I().
												Class("fas fa-fas fa-box").
												Aria("hidden", true),
										),
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
								ID("filled-with-icons-example-database-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-icon").
										Body(
											app.I().
												Class("fas fa-database").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Database"),
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
