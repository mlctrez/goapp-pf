package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BadgeToggle struct {
	app.Compo
}

func (c *BadgeToggle) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle pf-m-plain").
				ID("dropdown-badge-toggle-button").
				Aria("expanded", true).
				Type("button").
				Body(
					app.Span().
						Class("pf-c-badge pf-m-read").
						Body(
							app.Text("5"), app.Span().
								Class("pf-c-dropdown__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
			app.Ul().
				Class("pf-c-dropdown__menu").
				Aria("labelledby", "dropdown-badge-toggle-button").
				Body(
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Edit"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Deployment"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Application"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Count"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Application 1"),
								),
						),
				),
		)
}
