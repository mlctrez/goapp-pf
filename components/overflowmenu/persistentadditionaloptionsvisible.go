package overflowmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PersistentAdditionalOptionsVisible struct {
	app.Compo
}

func (c *PersistentAdditionalOptionsVisible) Render() app.UI {
	return app.Div().
		Class("pf-c-overflow-menu").
		ID("overflow-menu-persistent-visible-example").
		Body(
			app.Div().
				Class("pf-c-overflow-menu__content").
				Body(
					app.Div().
						Class("pf-c-overflow-menu__group pf-m-button-group").
						Body(
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("button").
										Body(
											app.Text("Primary"),
										),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Secondary"),
										),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-tertiary").
										Type("button").
										Body(
											app.Text("Tertiary"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-overflow-menu__control").
				Body(
					app.Div().
						Class("pf-c-dropdown pf-m-expanded").
						Body(
							app.Button().
								Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
								Type("button").
								ID("overflow-menu-persistent-visible-example-dropdown-toggle").
								Aria("label", "Dropdown for persistent example").
								Aria("expanded", true).
								Body(
									app.I().
										Class("fas fa-ellipsis-v").
										Aria("hidden", true),
								),
							app.Ul().
								Class("pf-c-dropdown__menu").
								Aria("labelledby", "overflow-menu-persistent-visible-example-dropdown-toggle").
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Action 4"),
												),
										),
								),
						),
				),
		)
}
