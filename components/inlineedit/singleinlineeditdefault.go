package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SingleInlineEditDefault struct {
	app.Compo
}

func (c *SingleInlineEditDefault) Render() app.UI {
	return app.Form().
		Class("pf-c-inline-edit").
		ID("single-inline-edit-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__group").
				Body(
					app.Div().
						Class("pf-c-inline-edit__value").
						ID("single-inline-edit-example-label").
						Body(
							app.Text("Static value"),
						),
					app.Div().
						Class("pf-c-inline-edit__action pf-m-enable-editable").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								ID("single-inline-edit-example-edit-button").
								Aria("label", "Edit").
								Aria("labelledby", "single-inline-edit-example-edit-button single-inline-edit-example-label").
								Body(
									app.I().
										Class("fas fa-pencil-alt").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-inline-edit__group").
				Body(
					app.Div().
						Class("pf-c-inline-edit__input").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								Value("Static value").
								Aria("label", "Editable text input"),
						),
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
				),
		)
}
