package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineEditActionGroupIconButtons struct {
	app.Compo
}

func (c *InlineEditActionGroupIconButtons) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit pf-m-inline-editable").
		ID("inline-edit-action-group-icon-buttons-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__group pf-m-action-group pf-m-icon-group").
				Body(
					app.Div().
						Class("pf-c-inline-edit__action pf-m-valid").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Save edits").
								Body(
									app.I().
										Class("fas fa-check").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-inline-edit__action").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Cancel edits").
								Body(
									app.I().
										Class("fas fa-times").
										Aria("hidden", true),
								),
						),
				),
		)
}
