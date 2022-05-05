package applauncher

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Collapsed struct {
	app.Compo
}

func (c *Collapsed) Render() app.UI {
	return app.Nav().
		Class("pf-c-app-launcher").
		Aria("label", "Application launcher").
		Body(
			app.Button().
				Class("pf-c-app-launcher__toggle").
				Type("button").
				ID("-button").
				Aria("expanded", "false").
				Aria("label", "Application launcher").
				Body(
					app.I().
						Class("fas fa-th").
						Aria("hidden", true),
				),
			app.Ul().
				Class("pf-c-app-launcher__menu").
				Aria("labelledby", "-button").
				Hidden(true).
				Body(
					app.Li().
						Body(
							app.A().
								Class("pf-c-app-launcher__menu-item").
								Href("#").
								Body(
									app.Text("Link"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-app-launcher__menu-item").
								Type("button").
								Body(
									app.Text("Action"),
								),
						),
					app.Li().
						Class("pf-c-divider").
						Aria("role", "separator"),
					app.Li().
						Body(
							app.A().
								Class("pf-c-app-launcher__menu-item pf-m-disabled").
								Href("#").
								Aria("disabled", true).
								TabIndex(-1).
								Body(
									app.Text("Disabled link"),
								),
						),
				),
		)
}
