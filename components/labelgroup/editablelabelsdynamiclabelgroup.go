package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type EditableLabelsDynamicLabelGroup struct {
	app.Compo
}

func (c *EditableLabelsDynamicLabelGroup) Render() app.UI {
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
												ID("editable-labels-editable-group-example-editable-label-editable-1-editable-content").
												Value("          Editable label 1\n        ").
												Aria("label", "Editable text").
												Body(
													app.Text("Editable label 1"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-editable-group-example-editable-label-editable-1-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-editable-group-example-editable-label-editable-1-button editable-labels-editable-group-example-editable-label-editable-1-text").
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
												ID("editable-labels-editable-group-example-editable-label-editable-2-editable-content").
												Value("          Editable label 2\n        ").
												Aria("label", "Editable text").
												Body(
													app.Text("Editable label 2"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-editable-group-example-editable-label-editable-2-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-editable-group-example-editable-label-editable-2-button editable-labels-editable-group-example-editable-label-editable-2-text").
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
												ID("editable-labels-editable-group-example-editable-label-editable-3-editable-content").
												Value("          Editable label 3\n        ").
												Aria("label", "Editable text").
												Body(
													app.Text("Editable label 3"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("editable-labels-editable-group-example-editable-label-editable-3-button").
												Aria("label", "Remove").
												Aria("labelledby", "editable-labels-editable-group-example-editable-label-editable-3-button editable-labels-editable-group-example-editable-label-editable-3-text").
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
