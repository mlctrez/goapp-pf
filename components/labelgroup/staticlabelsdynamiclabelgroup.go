package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StaticLabelsDynamicLabelGroup struct {
	app.Compo
}

func (c *StaticLabelsDynamicLabelGroup) Render() app.UI {
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
										Class("pf-c-label pf-m-green").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Text("Static label 1"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("static-labels-editable-group-example-editable-label-static-1-button").
												Aria("label", "Remove").
												Aria("labelledby", "static-labels-editable-group-example-editable-label-static-1-button static-labels-editable-group-example-editable-label-static-1-text").
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
										Class("pf-c-label pf-m-green").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Text("Static label 2"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("static-labels-editable-group-example-editable-label-static-2-button").
												Aria("label", "Remove").
												Aria("labelledby", "static-labels-editable-group-example-editable-label-static-2-button static-labels-editable-group-example-editable-label-static-2-text").
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
										Class("pf-c-label pf-m-green").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Text("Static label 3"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												ID("static-labels-editable-group-example-editable-label-static-3-button").
												Aria("label", "Remove").
												Aria("labelledby", "static-labels-editable-group-example-editable-label-static-3-button static-labels-editable-group-example-editable-label-static-3-text").
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
