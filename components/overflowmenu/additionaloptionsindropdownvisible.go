package overflowmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdditionalOptionsInDropdownVisible struct {
	app.Compo
}

func (c *AdditionalOptionsInDropdownVisible) Render() app.UI {
	return app.Div().
		Class("pf-c-overflow-menu").
		ID("overflow-menu-simple-additional-options-visible").
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
					app.Div().
						Class("pf-c-overflow-menu__group pf-m-icon-button-group").
						Body(
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Align left").
										Body(
											app.I().
												Class("fas fa-align-left").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Align center").
										Body(
											app.I().
												Class("fas fa-align-center").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Align right").
										Body(
											app.I().
												Class("fas fa-align-right").
												Aria("hidden", true),
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
								ID("overflow-menu-simple-additional-options-visible-dropdown-toggle").
								Aria("label", "Dropdown with additional options").
								Aria("expanded", true).
								Body(
									app.I().
										Class("fas fa-ellipsis-v").
										Aria("hidden", true),
								),
							app.Ul().
								Class("pf-c-dropdown__menu").
								Aria("labelledby", "overflow-menu-simple-additional-options-visible-dropdown-toggle").
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-dropdown__menu-item").
												Body(
													app.Text("Action 7"),
												),
										),
								),
						),
				),
		)
}
