package applauncher

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithSectionsAndDividersBetweenItems struct {
	app.Compo
}

func (c *WithSectionsAndDividersBetweenItems) Render() app.UI {
	return app.Nav().
		Class("pf-c-app-launcher pf-m-expanded").
		Aria("label", "Application launcher").
		Body(
			app.Button().
				Class("pf-c-app-launcher__toggle").
				Type("button").
				ID("-button").
				Aria("expanded", true).
				Aria("label", "Application launcher").
				Body(
					app.I().
						Class("fas fa-th").
						Aria("hidden", true),
				),
			app.Div().
				Class("pf-c-app-launcher__menu").
				Body(
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Text("Link not in group"),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
								),
						),
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.H1().
								Class("pf-c-app-launcher__group-title").
								Body(
									app.Text("Group 1"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Text("Group 1 link"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Text("Group 1 link"),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
								),
						),
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.H1().
								Class("pf-c-app-launcher__group-title").
								Body(
									app.Text("Group 2"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Text("Group 2 link"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Text("Group 2 link"),
												),
										),
								),
						),
				),
		)
}
