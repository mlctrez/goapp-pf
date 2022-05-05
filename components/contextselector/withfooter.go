package contextselector

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithFooter struct {
	app.Compo
}

func (c *WithFooter) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-context-selector pf-m-expanded").
				Body(
					app.Span().
						ID("context-selector-with-footer-example-label").
						Hidden(true).
						Body(
							app.Text("Selected Project:"),
						),
					app.Button().
						Class("pf-c-context-selector__toggle").
						Aria("expanded", true).
						ID("context-selector-with-footer-example-toggle").
						Aria("labelledby", "context-selector-with-footer-example-label context-selector-with-footer-example-toggle").
						Body(
							app.Span().
								Class("pf-c-context-selector__toggle-text").
								Body(
									app.Text("My project"),
								),
							app.Span().
								Class("pf-c-context-selector__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-context-selector__menu").
						Body(
							app.Div().
								Class("pf-c-context-selector__menu-search").
								Body(
									app.Div().
										Class("pf-c-search-input").
										Body(
											app.Div().
												Class("pf-c-search-input__bar").
												Body(
													app.Span().
														Class("pf-c-search-input__text").
														Body(
															app.Span().
																Class("pf-c-search-input__icon").
																Body(
																	app.I().
																		Class("fas fa-search fa-fw").
																		Aria("hidden", true),
																),
															app.Input().
																Class("pf-c-search-input__text-input").
																Type("text").
																Placeholder("false").
																Aria("label", "Search"),
														),
												),
										),
								),
							app.Ul().
								Class("pf-c-context-selector__menu-list").
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-context-selector__menu-list-item").
												Href("#").
												Body(
													app.Text("Link"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Action"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-context-selector__menu-list-item pf-m-disabled").
												Href("#").
												Aria("disabled", true).
												TabIndex(-1).
												Body(
													app.Text("Disabled link"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Disabled(true).
												Body(
													app.Text("Disabled action"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("My project"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("OpenShift cluster"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Production Ansible"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("AWS"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Azure"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("My project"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("OpenShift cluster"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Production Ansible"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("AWS"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Azure"),
												),
										),
								),
							app.Div().
								Class("pf-c-context-selector__menu-footer").
								Body(
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Manage projects"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-context-selector pf-m-expanded").
				Body(
					app.Span().
						ID("context-selector-with-footer-example-two-label").
						Hidden(true).
						Body(
							app.Text("Selected Project:"),
						),
					app.Button().
						Class("pf-c-context-selector__toggle").
						Aria("expanded", true).
						ID("context-selector-with-footer-example-two-toggle").
						Aria("labelledby", "context-selector-with-footer-example-two-label context-selector-with-footer-example-two-toggle").
						Body(
							app.Span().
								Class("pf-c-context-selector__toggle-text").
								Body(
									app.Text("My project"),
								),
							app.Span().
								Class("pf-c-context-selector__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-context-selector__menu").
						Body(
							app.Div().
								Class("pf-c-context-selector__menu-search").
								Body(
									app.Div().
										Class("pf-c-search-input").
										Body(
											app.Div().
												Class("pf-c-search-input__bar").
												Body(
													app.Span().
														Class("pf-c-search-input__text").
														Body(
															app.Span().
																Class("pf-c-search-input__icon").
																Body(
																	app.I().
																		Class("fas fa-search fa-fw").
																		Aria("hidden", true),
																),
															app.Input().
																Class("pf-c-search-input__text-input").
																Type("text").
																Placeholder("false").
																Aria("label", "Search"),
														),
												),
										),
								),
							app.Ul().
								Class("pf-c-context-selector__menu-list").
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-context-selector__menu-list-item").
												Href("#").
												Body(
													app.Text("Link"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Action"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-context-selector__menu-list-item pf-m-disabled").
												Href("#").
												Aria("disabled", true).
												TabIndex(-1).
												Body(
													app.Text("Disabled link"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Disabled(true).
												Body(
													app.Text("Disabled action"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("My project"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("OpenShift cluster"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Production Ansible"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("AWS"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Azure"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("My project"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("OpenShift cluster"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Production Ansible"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("AWS"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-context-selector__menu-list-item").
												Type("button").
												Body(
													app.Text("Azure"),
												),
										),
								),
							app.Div().
								Class("pf-c-context-selector__menu-footer").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link pf-m-inline").
										Type("button").
										Body(
											app.Text("Manage projects"),
										),
								),
						),
				),
		)
}
