package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineEditActionGroup struct {
	app.Compo
}

func (c *InlineEditActionGroup) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit pf-m-inline-editable").
		ID("inline-edit-action-group-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__group pf-m-action-group").
				Body(
					app.Div().
						Class("pf-c-inline-edit__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("Save"),
								),
						),
					app.Div().
						Class("pf-c-inline-edit__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Cancel"),
								),
						),
				),
		)
}
