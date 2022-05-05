package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type KebabAlignRight struct {
	app.Compo
}

func (c *KebabAlignRight) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle pf-m-plain").
				ID("dropdown-kebab-align-right-button").
				Aria("expanded", true).
				Type("button").
				Aria("label", "Actions").
				Body(
					app.I().
						Class("fas fa-ellipsis-v").
						Aria("hidden", true),
				),
			app.Ul().
				Class("pf-c-dropdown__menu pf-m-align-right").
				Aria("labelledby", "dropdown-kebab-align-right-button").
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
		)
}
