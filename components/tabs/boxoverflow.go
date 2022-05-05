package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BoxOverflow struct {
	app.Compo
}

func (c *BoxOverflow) Render() app.UI {
	return app.Div().
		Class("pf-c-tabs pf-m-box pf-m-scrollable").
		ID("box-overflow-example").
		Body(
			app.Button().
				Class("pf-c-tabs__scroll-button").
				Disabled(true).
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
								ID("box-overflow-example-users-link").
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
								ID("box-overflow-example-containers-link").
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
								ID("box-overflow-example-database-link").
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
								ID("box-overflow-example-server-link").
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
								ID("box-overflow-example-system-link").
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
								ID("box-overflow-example-network-wired-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Network"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("box-overflow-example-cloud-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Hybrid Cloud"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("box-overflow-example-automation-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Automation"),
										),
								),
						),
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("box-overflow-example-files-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Files"),
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
