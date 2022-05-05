package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PlainWithText struct {
	app.Compo
}

func (c *PlainWithText) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-options-menu").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
						Type("button").
						ID("options-menu-disabled-text-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Disabled(true).
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.Text("Custom text"),
								),
							app.Div().
								Class("pf-c-options-menu__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-disabled-text-example-toggle").
						Hidden(true).
						Body(
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 1"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 2"), app.Div().
												Class("pf-c-options-menu__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 3"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-options-menu").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
						Type("button").
						ID("options-menu-plain-text-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.Text("Custom text"),
								),
							app.Div().
								Class("pf-c-options-menu__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-plain-text-example-toggle").
						Hidden(true).
						Body(
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 1"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 2"), app.Div().
												Class("pf-c-options-menu__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 3"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-options-menu pf-m-expanded").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
						Type("button").
						ID("options-menu-plain-text-expanded-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", true).
						Body(
							app.Span().
								Class("pf-c-options-menu__toggle-text").
								Body(
									app.Text("Custom text"),
								),
							app.Div().
								Class("pf-c-options-menu__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-plain-text-expanded-example-toggle").
						Body(
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 1"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 2"), app.Div().
												Class("pf-c-options-menu__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-options-menu__menu-item").
										Type("button").
										Body(
											app.Text("Option 3"),
										),
								),
						),
				),
		)
}
