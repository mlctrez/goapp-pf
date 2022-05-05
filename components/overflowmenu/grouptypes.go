package overflowmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupTypes struct {
	app.Compo
}

func (c *GroupTypes) Render() app.UI {
	return app.Div().
		Class("pf-c-overflow-menu").
		ID("overflow-menu-button-group-example").
		Body(
			app.Div().
				Class("pf-c-overflow-menu__content").
				Body(
					app.Div().
						Class("pf-c-overflow-menu__group").
						Body(
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 1"),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 2"),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 3"),
								),
						),
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
		)
}
