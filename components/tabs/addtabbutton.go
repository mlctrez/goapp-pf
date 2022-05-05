package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AddTabButton struct {
	app.Compo
}

func (c *AddTabButton) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tabs pf-m-scrollable").
				ID("close-default-example").
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
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-users-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-containers-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-database-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-server-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Server"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-system-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("System"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-default-example-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Network"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
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
					app.Span().
						Class("pf-c-tabs__add").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Add tab").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-tabs pf-m-secondary pf-m-scrollable").
				ID("close-secondary-example").
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
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-users-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-containers-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-database-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-server-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Server"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-system-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("System"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-secondary-example-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Network"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
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
					app.Span().
						Class("pf-c-tabs__add").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Add tab").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tabs pf-m-box pf-m-scrollable").
				ID("close-box-example").
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
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-users-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-containers-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-database-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-server-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Server"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-system-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("System"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-example-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Network"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
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
					app.Span().
						Class("pf-c-tabs__add").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Add tab").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tabs pf-m-box pf-m-color-scheme--light-300 pf-m-scrollable").
				ID("close-box-light-300-example").
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
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-users-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-containers-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-database-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-server-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Server"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-system-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("System"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-action").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("close-box-light-300-example-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Network"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Body(
													app.Span().
														Class("pf-c-tabs__item-close-icon").
														Body(
															app.I().
																Class("fas fa-times").
																Aria("hidden", true),
														),
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
					app.Span().
						Class("pf-c-tabs__add").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Add tab").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
		)
}
