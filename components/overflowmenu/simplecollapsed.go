package overflowmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleCollapsed struct {
	app.Compo
}

func (c *SimpleCollapsed) Render() app.UI {
	return app.Div().
		Class("pf-c-overflow-menu").
		ID("overflow-menu-simple").
		Body(
			app.Div().
				Class("pf-c-overflow-menu__control").
				Body(
					app.Div().
						Class("pf-c-dropdown pf-m-expanded").
						Body(
							app.Button().
								Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
								Type("button").
								ID("overflow-menu-simple-dropdown-toggle").
								Aria("label", "Generic options").
								Aria("expanded", true).
								Body(
									app.I().
										Class("fas fa-ellipsis-v").
										Aria("hidden", true),
								),
							app.Ul().
								Class("pf-c-dropdown__menu").
								Aria("labelledby", "overflow-menu-simple-dropdown-toggle").
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Item 1"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Item 2"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Item 3"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Item 4"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Item 5"),
												),
										),
								),
						),
				),
		)
}
