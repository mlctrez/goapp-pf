package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AlignOnDifferentBreakpoint struct {
	app.Compo
}

func (c *AlignOnDifferentBreakpoint) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("-button").
				Aria("expanded", true).
				Type("button").
				Body(
					app.Span().
						Class("pf-c-dropdown__toggle-text").
						Body(
							app.Text("Dropdown"),
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
				Class("pf-c-dropdown__menu pf-m-align-right-on-lg pf-m-align-left-on-2xl").
				Aria("labelledby", "-button").
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
