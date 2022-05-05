package actionlist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ActionListSingleGroup struct {
	app.Compo
}

func (c *ActionListSingleGroup) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-action-list").
				Body(
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("Next"),
								),
						),
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Back"),
								),
						),
				),
			app.Br(),
			app.Text("With kebab"),
			app.Div().
				Class("pf-c-action-list").
				Body(
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("Next"),
								),
						),
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Back"),
								),
						),
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Kebab").
								Body(
									app.I().
										Class("fas fa-ellipsis-v").
										Aria("hidden", true),
								),
						),
				),
		)
}
