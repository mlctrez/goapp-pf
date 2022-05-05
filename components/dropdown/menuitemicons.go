package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MenuItemIcons struct {
	app.Compo
}

func (c *MenuItemIcons) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("dropdown-menu-item-icons-button").
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
				Aria("labelledby", "dropdown-menu-item-icons-button").
				Body(
					app.Li().
						Body(
							app.A().
								Class("pf-c-dropdown__menu-item pf-m-icon").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-dropdown__menu-item-icon").
										Body(
											app.Img().
												Src("/assets/images/pf-logo-small.svg").
												Alt("PatternFly logo"),
										),
									app.Text("Link"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item pf-m-icon").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-dropdown__menu-item-icon").
										Body(
											app.I().
												Class("fas fa-cog").
												Aria("hidden", true),
										),
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
