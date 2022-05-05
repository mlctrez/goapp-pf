package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StateInvalid struct {
	app.Compo
}

func (c *StateInvalid) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit pf-m-inline-editable").
		ID("inline-edit-state-invalid").
		Body(
			app.Div().
				Class("pf-c-inline-edit__group").
				Body(
					app.Label().
						Class("pf-c-inline-edit__label").
						ID("inline-edit-state-invalid-label").
						For("inline-edit-state-invalid-input").
						Body(
							app.Text("Invalid example"),
						),
					app.Div().
						Class("pf-c-inline-edit__action pf-m-enable-editable").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								ID("inline-edit-state-invalid-edit-button").
								Aria("label", "Edit").
								Aria("labelledby", "inline-edit-state-invalid-label inline-edit-state-invalid-edit-button").
								Body(
									app.I().
										Class("fas fa-pencil-alt").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-inline-edit__value").
				Body(
					app.Text("Static value"),
				),
			app.Div().
				Class("pf-c-inline-edit__group").
				Body(
					app.Div().
						Class("pf-c-inline-edit__input").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Required(true).
								Value("Invalid state").
								Aria("invalid", true).
								Aria("label", "Error state input example"),
						),
					app.Div().
						Class("pf-c-inline-edit__group pf-m-action-group pf-m-icon-group").
						Body(
							app.Div().
								Class("pf-c-inline-edit__action").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
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
				),
		)
}
