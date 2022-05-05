package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticalExpandableResponsive struct {
	app.Compo
}

func (c *VerticalExpandableResponsive) Render() app.UI {
	return app.Div().
		Class("pf-c-tabs pf-m-expandable pf-m-vertical pf-m-non-expandable-on-md pf-m-expandable-on-lg pf-m-non-expandable-on-xl").
		ID("vertical-expandable-responsive-example").
		Body(
			app.Div().
				Class("pf-c-tabs__toggle").
				Body(
					app.Div().
						Class("pf-c-tabs__toggle-button").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Details").
								Aria("expanded", "false").
								ID("vertical-expandable-responsive-example-toggle-button").
								Aria("labelledby", "vertical-expandable-responsive-example-toggle-text vertical-expandable-responsive-example-toggle-button").
								Body(
									app.Span().
										Class("pf-c-tabs__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-tabs__toggle-text").
										ID("vertical-expandable-responsive-example-toggle-text").
										Body(
											app.Text("Containers"),
										),
								),
						),
				),
			app.Ul().
				Class("pf-c-tabs__list").
				Body(
					app.Li().
						Class("pf-c-tabs__item").
						Body(
							app.Button().
								Class("pf-c-tabs__link").
								ID("vertical-expandable-responsive-example-users-link").
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
								ID("vertical-expandable-responsive-example-containers-link").
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
								ID("vertical-expandable-responsive-example-database-link").
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
								ID("vertical-expandable-responsive-example-server-link").
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
								ID("vertical-expandable-responsive-example-system-link").
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
								ID("vertical-expandable-responsive-example-network-wired-link").
								Body(
									app.Span().
										Class("pf-c-tabs__item-text").
										Body(
											app.Text("Network"),
										),
								),
						),
				),
		)
}
