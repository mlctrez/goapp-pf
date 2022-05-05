package contextselector

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-context-selector").
				Body(
					app.Span().
						ID("context-selector-collapsed-example-label").
						Hidden(true).
						Body(
							app.Text("Selected project:"),
						),
					app.Button().
						Class("pf-c-context-selector__toggle").
						Aria("expanded", "false").
						ID("context-selector-collapsed-example-toggle").
						Aria("labelledby", "context-selector-collapsed-example-label context-selector-collapsed-example-toggle").
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
						Hidden(true).
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
						),
				),
			app.Div().
				Class("pf-c-context-selector pf-m-expanded").
				Body(
					app.Span().
						ID("context-selector-expanded-example-label").
						Hidden(true).
						Body(
							app.Text("Selected Project:"),
						),
					app.Button().
						Class("pf-c-context-selector__toggle").
						Aria("expanded", true).
						ID("context-selector-expanded-example-toggle").
						Aria("labelledby", "context-selector-expanded-example-label context-selector-expanded-example-toggle").
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
						),
				),
		)
}
