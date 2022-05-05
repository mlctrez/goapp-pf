package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type EditableLabelsLabelActiveDynamicLabelGroup struct {
	app.Compo
}

func (c *EditableLabelsLabelActiveDynamicLabelGroup) Render() app.UI {
	return app.Div().
		Class("pf-c-label-group pf-m-editable").
		Body(
			app.Div().
				Class("pf-c-label-group__main").
				Body(
					app.Ul().
						Class("pf-c-label-group__list").
						Aria("role", "list").
						Aria("label", "Group of labels").
						Body(
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-blue pf-m-editable").
										Body(
											app.Button().
												Class("pf-c-label__content").
												ID("editable-labels-label-active-editable-group-example-editable-label-default-1-editable-content").
												Value("          Editable label 1\n        ").
												Aria("label", "Editable text").
												Body(
													app.Text("Editable label 1"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-label-active-editable-group-example-editable-label-default-1-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-label-active-editable-group-example-editable-label-default-1-button editable-labels-label-active-editable-group-example-editable-label-default-1-text").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-blue pf-m-editable").
										Body(
											app.Button().
												Class("pf-c-label__content").
												ID("editable-labels-label-active-editable-group-example-editable-label-default-2-editable-content").
												Value("          Editable label 2\n        ").
												Aria("label", "Editable text").
												Body(
													app.Text("Editable label 2"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-label-active-editable-group-example-editable-label-default-2-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-label-active-editable-group-example-editable-label-default-2-button editable-labels-label-active-editable-group-example-editable-label-default-2-text").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-blue pf-m-active pf-m-editable pf-m-editable-active").
										Body(
											app.Input().
												Class("pf-c-label__content").
												ID("editable-labels-label-active-editable-group-example-editable-label-active-editable-content").
												Type("text").
												Value("          Editable label 3, active\n        ").
												Aria("label", "Editable text"),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-label-active-editable-group-example-editable-label-active-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-label-active-editable-group-example-editable-label-active-button editable-labels-label-active-editable-group-example-editable-label-active-text").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item pf-m-textarea").
								Body(
									app.Textarea().
										Class("pf-c-label-group__textarea").
										Rows(1).
										TabIndex(0).
										Aria("label", "New label"),
								),
						),
				),
		)
}
