package breadcrumb

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDropdown struct {
	app.Compo
}

func (c *WithDropdown) Render() app.UI {
	return app.Nav().
		Class("pf-c-breadcrumb").
		Aria("label", "breadcrumb").
		Body(
			app.Ol().
				Class("pf-c-breadcrumb__list").
				Body(
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section home"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section title"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-breadcrumb__dropdown").
								Body(
									app.Div().
										Class("pf-c-dropdown pf-m-expanded").
										Body(
											app.Button().
												Class("pf-c-dropdown__toggle pf-m-plain").
												ID("dropdown-badge-toggle-button").
												Aria("expanded", true).
												Type("button").
												Body(
													app.Span().
														Class("pf-c-badge pf-m-read").
														Body(
															app.Text("5"), app.Span().
																Class("pf-c-dropdown__toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
												),
											app.Ul().
												Class("pf-c-dropdown__menu").
												Aria("labelledby", "dropdown-badge-toggle-button").
												Body(
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Section title"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Section title"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Section title"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Section title"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-dropdown__menu-item").
																Type("button").
																Body(
																	app.Text("Section title"),
																),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.H1().
								Class("pf-c-breadcrumb__heading").
								Body(
									app.A().
										Href("#").
										Class("pf-c-breadcrumb__link pf-m-current").
										Aria("current", "page").
										Body(
											app.Text("Section title"),
										),
								),
						),
				),
		)
}
