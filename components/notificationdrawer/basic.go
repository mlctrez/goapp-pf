package notificationdrawer

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-notification-drawer").
		Body(
			app.Div().
				Class("pf-c-notification-drawer__header").
				Body(
					app.H1().
						Class("pf-c-notification-drawer__header-title").
						Body(
							app.Text("Notifications"),
						),
					app.Span().
						Class("pf-c-notification-drawer__header-status").
						Body(
							app.Text("3 unread"),
						),
					app.Div().
						Class("pf-c-notification-drawer__header-action").
						Body(
							app.Div().
								Class("pf-c-dropdown").
								Body(
									app.Button().
										Class("pf-c-dropdown__toggle pf-m-plain").
										ID("notification-drawer-basic-header-action-button").
										Aria("expanded", "false").
										Type("button").
										Aria("label", "Actions").
										Body(
											app.I().
												Class("fas fa-ellipsis-v").
												Aria("hidden", true),
										),
									app.Ul().
										Class("pf-c-dropdown__menu pf-m-align-right").
										Aria("labelledby", "notification-drawer-basic-header-action-button").
										Hidden(true).
										Body(
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Link"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-dropdown__menu-item").
														Type("button").
														Body(
															app.Text("Action"),
														),
												),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item pf-m-disabled").
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
														Class("pf-c-dropdown__menu-item").
														Type("button").
														Disabled(true).
														Body(
															app.Text("Disabled action"),
														),
												),
											app.Li().
												Class("pf-c-divider").
												Aria("role", "separator"),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Separated link"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-notification-drawer__header-action-close").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Close").
										Body(
											app.I().
												Class("fas fa-times").
												Aria("hidden", true),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-notification-drawer__body").
				Body(
					app.Ul().
						Class("pf-c-notification-drawer__list").
						Body(
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-hoverable pf-m-info").
								TabIndex(0).
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-info-circle").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Info notification:"),
														),
													app.Text("Unread\n            info notification title"),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-1-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-1-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This is an info notification description."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("5 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-hoverable pf-m-default").
								TabIndex(0).
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-arrow-circle-up").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Default notification:"),
														),
													app.Text("Unread\n            recommendation notification title. This is a long title to show how the title will wrap if it is long and wraps to multiple lines."),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-2-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-2-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This is a recommendation notification description. This is a long description to show how the title will wrap if it is long and wraps to multiple lines."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("10 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-hoverable pf-m-default").
								TabIndex(0).
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-arrow-circle-up").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Default notification:"),
														),
													app.Text("Unread\n            recommendation notification title. This is a long title to show how the title will wrap if it is long and wraps to multiple lines."),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-3-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-3-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This is a recommendation notification description. This is a long description to show how the title will wrap if it is long and wraps to multiple lines."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("20 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-read pf-m-warning pf-m-hoverable").
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-exclamation-triangle").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Warning notification:"),
														),
													app.Text("Read warning notification title"),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown pf-m-top").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-4-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-4-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This is a warning notification description."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("20 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-read pf-m-success pf-m-hoverable").
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-check-circle").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Success notification:"),
														),
													app.Text("Read success notification title"),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown pf-m-top").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-5-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-5-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This is a success notification description."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("30 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-read pf-m-success pf-m-hoverable").
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-check-circle").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title pf-m-truncate").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Success notification:"),
														),
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent quis odio risus. Ut dictum vitae sapien at posuere. Nullam suscipit massa quis lacus pellentesque scelerisque. Donec non maximus neque, quis ornare nunc. Vivamus in nibh sed libero feugiat feugiat. Nulla lacinia rutrum est, a commodo odio vestibulum suscipit. Nullam id quam et quam porttitor interdum quis nec tellus. Vestibulum arcu dui, pulvinar eu tellus in, semper mattis diam. Sed commodo tincidunt lacus non pulvinar. Curabitur tempor molestie vestibulum. Vivamus vel mi dignissim, efficitur neque eget, efficitur massa. Mauris vitae nunc augue. Donec augue lorem, malesuada et quam vitae, volutpat mattis nisi. Nullam nec venenatis ex, quis lobortis purus. Sed nisl dolor, mattis sit amet tincidunt quis, mollis sed massa."),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown pf-m-top").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-6-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-6-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This example uses \".pf-m-truncate\" to limit the title to a single line and truncate any overflow text with ellipses."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("40 minutes ago"),
										),
								),
							app.Li().
								Class("pf-c-notification-drawer__list-item pf-m-read pf-m-success pf-m-hoverable").
								Body(
									app.Div().
										Class("pf-c-notification-drawer__list-item-header").
										Body(
											app.Span().
												Class("pf-c-notification-drawer__list-item-header-icon").
												Body(
													app.I().
														Class("fas fa-check-circle").
														Aria("hidden", true),
												),
											app.H2().
												Class("pf-c-notification-drawer__list-item-header-title pf-m-truncate").
												Style("--pf-c-notification-drawer__list-item-header-title--max-lines", " 2").
												Body(
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("Success notification:"),
														),
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent quis odio risus. Ut dictum vitae sapien at posuere. Nullam suscipit massa quis lacus pellentesque scelerisque. Donec non maximus neque, quis ornare nunc. Vivamus in nibh sed libero feugiat feugiat. Nulla lacinia rutrum est, a commodo odio vestibulum suscipit. Nullam id quam et quam porttitor interdum quis nec tellus. Vestibulum arcu dui, pulvinar eu tellus in, semper mattis diam. Sed commodo tincidunt lacus non pulvinar. Curabitur tempor molestie vestibulum. Vivamus vel mi dignissim, efficitur neque eget, efficitur massa. Mauris vitae nunc augue. Donec augue lorem, malesuada et quam vitae, volutpat mattis nisi. Nullam nec venenatis ex, quis lobortis purus. Sed nisl dolor, mattis sit amet tincidunt quis, mollis sed massa."),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-action").
										Body(
											app.Div().
												Class("pf-c-dropdown pf-m-top").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("notification-drawer-basicdropdown-kebab-7-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "notification-drawer-basicdropdown-kebab-7-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
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
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-description").
										Body(
											app.Text("This example uses \".pf-m-truncate\" and sets \"--pf-c-notification-drawer__list-item-header-title--max-lines: 2\" to limit title to two lines and truncate any overflow text with ellipses."),
										),
									app.Div().
										Class("pf-c-notification-drawer__list-item-timestamp").
										Body(
											app.Text("50 minutes ago"),
										),
								),
						),
				),
		)
}
