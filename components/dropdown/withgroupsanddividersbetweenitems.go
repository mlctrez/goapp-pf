package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithGroupsAndDividersBetweenItems struct {
	app.Compo
}

func (c *WithGroupsAndDividersBetweenItems) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("dropdown-groups-and-dividers-between-items-button").
				Aria("expanded", true).
				Type("button").
				Body(
					app.Span().
						Class("pf-c-dropdown__toggle-text").
						Body(
							app.Text("Groups"),
						),
					app.Span().
						Class("pf-c-dropdown__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-dropdown__menu").
				Body(
					app.Section().
						Class("pf-c-dropdown__group").
						Body(
							app.Ul().
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
										Class("pf-c-divider").
										Aria("role", "separator"),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Type("button").
												Body(
													app.Text("Action"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-dropdown__group").
						Body(
							app.H1().
								Class("pf-c-dropdown__group-title").
								Body(
									app.Text("Group 2"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-dropdown__menu-item").
												Href("#").
												Body(
													app.Text("Group 2 link"),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Type("button").
												Body(
													app.Text("Group 2 action"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-dropdown__group").
						Body(
							app.H1().
								Class("pf-c-dropdown__group-title").
								Body(
									app.Text("Group 3"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-dropdown__menu-item").
												Href("#").
												Body(
													app.Text("Group 3 link"),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Type("button").
												Body(
													app.Text("Group 3 action"),
												),
										),
								),
						),
				),
		)
}
