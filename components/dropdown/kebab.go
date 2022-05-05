package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Kebab struct {
	app.Compo
}

func (c *Kebab) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle pf-m-plain").
						ID("dropdown-kebab-disabled-button").
						Aria("expanded", "false").
						Type("button").
						Disabled(true).
						Aria("label", "Actions").
						Body(
							app.I().
								Class("fas fa-ellipsis-v").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-kebab-disabled-button").
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
				Class("pf-c-dropdown").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle pf-m-plain").
						ID("dropdown-kebab-button").
						Aria("expanded", "false").
						Type("button").
						Aria("label", "Actions").
						Body(
							app.I().
								Class("fas fa-ellipsis-v").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-kebab-button").
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
						Class("pf-c-dropdown__toggle pf-m-plain").
						ID("dropdown-kebab-expanded-button").
						Aria("expanded", true).
						Type("button").
						Aria("label", "Actions").
						Body(
							app.I().
								Class("fas fa-ellipsis-v").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Aria("labelledby", "dropdown-kebab-expanded-button").
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
