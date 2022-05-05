package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineEditToggle struct {
	app.Compo
}

func (c *InlineEditToggle) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit").
		ID("inline-edit-toggle-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__action pf-m-enable-editable").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("inline-edit-toggle-example-edit-button").
						Aria("label", "Edit").
						Aria("labelledby", "inline-edit-toggle-example-edit-button inline-edit-toggle-example-label").
						Body(
							app.I().
								Class("fas fa-pencil-alt").
								Aria("hidden", true),
						),
				),
		)
}
