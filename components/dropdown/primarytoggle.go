package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PrimaryToggle struct {
	app.Compo
}

func (c *PrimaryToggle) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle pf-m-primary").
						ID("dropdown-primary-toggle-button").
						Aria("expanded", "false").
						Type("button").
						Body(
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Body(
									app.Text("Collapsed dropdown"),
								),
							app.Span().
								Class("pf-c-dropdown__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-primary-toggle-button").
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
				Class("pf-c-dropdown pf-m-expanded").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle pf-m-primary").
						ID("dropdown-primary-toggle-expanded-button").
						Aria("expanded", true).
						Type("button").
						Body(
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Body(
									app.Text("Expanded dropdown"),
								),
							app.Span().
								Class("pf-c-dropdown__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-primary-toggle-expanded-button").
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
				Class("pf-c-dropdown").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle pf-m-primary").
						ID("dropdown-primary-toggle-disabled-button").
						Aria("expanded", "false").
						Type("button").
						Disabled(true).
						Body(
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Body(
									app.Text("Disabled"),
								),
							app.Span().
								Class("pf-c-dropdown__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-primary-toggle-disabled-button").
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
		)
}
