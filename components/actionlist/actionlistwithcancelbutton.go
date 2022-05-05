package actionlist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ActionListWithCancelButton struct {
	app.Compo
}

func (c *ActionListWithCancelButton) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Text("In modals, forms, data lists"), app.Div().
				Class("pf-c-action-list").
				Body(
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("Save"),
								),
						),
					app.Div().
						Class("pf-c-action-list__item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link").
								Type("button").
								Body(
									app.Text("Cancel"),
								),
						),
				),
			app.Br(),
			app.Text("In wizards"),
			app.Div().
				Class("pf-c-action-list").
				Body(
					app.Div().
						Class("pf-c-action-list__group").
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
					app.Div().
						Class("pf-c-action-list__group").
						Body(
							app.Div().
								Class("pf-c-action-list__item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Body(
											app.Text("Cancel"),
										),
								),
						),
				),
		)
}
