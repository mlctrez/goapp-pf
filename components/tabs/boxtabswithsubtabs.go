package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BoxTabsWithSubTabs struct {
	app.Compo
}

func (c *BoxTabsWithSubTabs) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tabs pf-m-box pf-m-scrollable").
				ID("box-parent-example").
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
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-parent-example-users-link").
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
										ID("box-parent-example-containers-link").
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
										ID("box-parent-example-database-link").
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
										ID("box-parent-example-server-link").
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
										ID("box-parent-example-system-link").
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
										ID("box-parent-example-network-wired-link").
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
				),
			app.Div().
				Class("pf-c-tabs pf-m-secondary pf-m-scrollable").
				ID("box-child-example").
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
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-1-link").
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
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-2-link").
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
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-3-link").
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
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-4-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 4"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-5-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Item 5"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-child-example-sub-6-link").
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
						Aria("label", "Scroll right").
						Body(
							app.I().
								Class("fas fa-angle-right").
								Aria("hidden", true),
						),
				),
		)
}
