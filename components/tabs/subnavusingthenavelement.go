package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SubNavUsingTheNavElement struct {
	app.Compo
}

func (c *SubNavUsingTheNavElement) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Nav().
				Class("pf-c-tabs").
				Aria("label", "Local").
				ID("primary-nav-example").
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
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("primary-nav-example-users-link").
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
										ID("primary-nav-example-containers-link").
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
										ID("primary-nav-example-database-link").
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
										Class("pf-c-tabs__link").
										Href("#").
										ID("primary-nav-example-server-link").
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
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("primary-nav-example-system-link").
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
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("primary-nav-example-network-wired-link").
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
				),
			app.Nav().
				Class("pf-c-tabs pf-m-secondary").
				Aria("label", "Local secondary").
				ID("secondary-nav-example").
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
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("secondary-nav-example-sub-1-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 1"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("secondary-nav-example-sub-2-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 2"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current").
								Body(
									app.A().
										Class("pf-c-tabs__link").
										Href("#").
										ID("secondary-nav-example-sub-3-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 3"),
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
										ID("secondary-nav-example-sub-disabled-link").
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
										ID("secondary-nav-example-sub-aria-disabled-link").
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
										ID("secondary-nav-example-sub-6-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 6"),
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
				),
		)
}
