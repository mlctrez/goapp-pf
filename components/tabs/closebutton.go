package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CloseButton struct {
	app.Compo
}

func (c *CloseButton) Render() app.UI {
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-default-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-default-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-default-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-box-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-box-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-box-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-box-light-300-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-box-light-300-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-box-light-300-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tabs pf-m-scrollable").
				ID("close-icons-text-example").
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
										ID("close-icons-text-example-users-link").
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
										ID("close-icons-text-example-containers-link").
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
										ID("close-icons-text-example-database-link").
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-icons-text-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon").
												Body(
													app.I().
														Class("fas fa-server").
														Aria("hidden", true),
												),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-icons-text-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon").
												Body(
													app.I().
														Class("fas fa-laptop").
														Aria("hidden", true),
												),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-icons-text-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon").
												Body(
													app.I().
														Class("fas fa-laptop").
														Aria("hidden", true),
												),
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
										ID("close-icons-text-example-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-icon").
												Body(
													app.I().
														Class("fas fa-project-diagram").
														Aria("hidden", true),
												),
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
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tabs pf-m-fill pf-m-scrollable").
				ID("close-filled-example").
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
										ID("close-filled-example-users-link").
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
										ID("close-filled-example-containers-link").
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
										ID("close-filled-example-database-link").
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-filled-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-filled-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-filled-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
										ID("close-filled-example-network-wired-link").
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
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tabs pf-m-scrollable").
				ID("close-secondary-primary-example").
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
										ID("close-secondary-primary-example-users-link").
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
										ID("close-secondary-primary-example-containers-link").
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
										ID("close-secondary-primary-example-database-link").
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-secondary-primary-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-secondary-primary-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-secondary-primary-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
										ID("close-secondary-primary-example-network-wired-link").
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
				),
			app.Div().
				Class("pf-c-tabs pf-m-secondary pf-m-scrollable").
				ID("close-secondary-secondary-example").
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
										ID("close-secondary-secondary-example-users-link").
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
										ID("close-secondary-secondary-example-containers-link").
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
										ID("close-secondary-secondary-example-database-link").
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-secondary-secondary-example-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("close-secondary-secondary-example-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
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
								Class("pf-c-tabs__item pf-m-action pf-m-disabled").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("close-secondary-secondary-example-close-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Close disabled"),
												),
										),
									app.Span().
										Class("pf-c-tabs__item-close").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Close tab").
												Disabled(true).
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
										ID("close-secondary-secondary-example-network-wired-link").
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
				),
		)
}
