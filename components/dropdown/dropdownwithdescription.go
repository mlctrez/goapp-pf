package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DropdownWithDescription struct {
	app.Compo
}

func (c *DropdownWithDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("dropdown-with-description-button").
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
				Aria("labelledby", "dropdown-with-description-button").
				Body(
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item pf-m-description").
								Type("button").
								Body(
									app.Div().
										Class("pf-c-dropdown__menu-item-main").
										Body(
											app.Text("Menu item default"),
										),
									app.Div().
										Class("pf-c-dropdown__menu-item-description").
										Body(
											app.Text("This is a description"),
										),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item pf-m-description").
								Type("button").
								Body(
									app.Div().
										Class("pf-c-dropdown__menu-item-main").
										Body(
											app.Text("Menu item with long description"),
										),
									app.Div().
										Class("pf-c-dropdown__menu-item-description").
										Body(
											app.Text("This is a really long description that describes the menu item."),
										),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item pf-m-description").
								Type("button").
								Disabled(true).
								Body(
									app.Div().
										Class("pf-c-dropdown__menu-item-main").
										Body(
											app.Text("Menu item disabled"),
										),
									app.Div().
										Class("pf-c-dropdown__menu-item-description").
										Body(
											app.Text("This is a description"),
										),
								),
						),
					app.Li().
						Body(
							app.A().
								Class("pf-c-dropdown__menu-item pf-m-icon pf-m-description").
								Href("#").
								Body(
									app.Div().
										Class("pf-c-dropdown__menu-item-main").
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
									app.Div().
										Class("pf-c-dropdown__menu-item-description").
										Body(
											app.Text("This is a description"),
										),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item pf-m-icon pf-m-description").
								Type("button").
								Body(
									app.Div().
										Class("pf-c-dropdown__menu-item-main").
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
									app.Div().
										Class("pf-c-dropdown__menu-item-description").
										Body(
											app.Text("This is a description"),
										),
								),
						),
				),
		)
}
