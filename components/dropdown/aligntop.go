package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AlignTop struct {
	app.Compo
}

func (c *AlignTop) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-dropdown pf-m-top").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle").
						ID("dropdown-align-top-button").
						Aria("expanded", "false").
						Type("button").
						Body(
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Body(
									app.Text("Top"),
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
						Aria("labelledby", "dropdown-align-top-button").
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
				Class("pf-c-dropdown pf-m-expanded pf-m-top").
				Body(
					app.Button().
						Class("pf-c-dropdown__toggle").
						ID("dropdown-align-top-expanded-button").
						Aria("expanded", true).
						Type("button").
						Body(
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Body(
									app.Text("Top"),
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
						Aria("labelledby", "dropdown-align-top-expanded-button").
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
