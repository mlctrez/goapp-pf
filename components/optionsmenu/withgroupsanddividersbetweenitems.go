package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithGroupsAndDividersBetweenItems struct {
	app.Compo
}

func (c *WithGroupsAndDividersBetweenItems) Render() app.UI {
	return app.Div().
		Class("pf-c-options-menu pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-options-menu__toggle").
				Type("button").
				ID("options-menu-groups-and-dividers-between-items-toggle").
				Aria("haspopup", "listbox").
				Aria("expanded", true).
				Body(
					app.Span().
						Class("pf-c-options-menu__toggle-text").
						Body(
							app.Text("Options menu"),
						),
					app.Div().
						Class("pf-c-options-menu__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-options-menu__menu").
				Aria("labelledby", "options-menu-groups-and-dividers-between-items-toggle").
				Body(
					app.Section().
						Class("pf-c-options-menu__group").
						Body(
							app.Ul().
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 1"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 2"), app.Div().
														Class("pf-c-options-menu__menu-item-icon").
														Body(
															app.I().
																Class("fas fa-check").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
								),
						),
					app.Section().
						Class("pf-c-options-menu__group").
						Body(
							app.H1().
								Class("pf-c-options-menu__group-title").
								Body(
									app.Text("Group 1"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 1"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 2"),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
								),
						),
					app.Section().
						Class("pf-c-options-menu__group").
						Body(
							app.H1().
								Class("pf-c-options-menu__group-title").
								Body(
									app.Text("Group 2"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 1"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("Option 2"),
												),
										),
								),
						),
				),
		)
}
