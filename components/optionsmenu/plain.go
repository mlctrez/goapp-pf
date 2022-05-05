package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Plain struct {
	app.Compo
}

func (c *Plain) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-options-menu").
				Body(
					app.Button().
						Class("pf-c-options-menu__toggle pf-m-plain").
						Type("button").
						ID("options-menu-plain-disabled-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Disabled(true).
						Aria("label", "Sort by").
						Body(
							app.I().
								Class("fas fa-sort-amount-down").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-plain-disabled-example-toggle").
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
						Class("pf-c-options-menu__toggle pf-m-plain").
						Type("button").
						ID("options-menu-plain-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", "false").
						Aria("label", "Sort by").
						Body(
							app.I().
								Class("fas fa-sort-amount-down").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-plain-example-toggle").
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
						Class("pf-c-options-menu__toggle pf-m-plain").
						Type("button").
						ID("options-menu-plain-expanded-example-toggle").
						Aria("haspopup", "listbox").
						Aria("expanded", true).
						Aria("label", "Sort by").
						Body(
							app.I().
								Class("fas fa-sort-amount-down").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-options-menu__menu").
						Aria("labelledby", "options-menu-plain-expanded-example-toggle").
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
