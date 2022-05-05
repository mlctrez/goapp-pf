package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CompactLabelsVertical struct {
	app.Compo
}

func (c *CompactLabelsVertical) Render() app.UI {
	return app.Div().
		Class("pf-c-label-group pf-m-vertical").
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
										Class("pf-c-label pf-m-compact").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Span().
														Class("pf-c-label__icon").
														Body(
															app.I().
																Class("fas fa-fw fa-info-circle").
																Aria("hidden", true),
														),
													app.Text("Label"),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-compact pf-m-blue").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Span().
														Class("pf-c-label__icon").
														Body(
															app.I().
																Class("fas fa-fw fa-info-circle").
																Aria("hidden", true),
														),
													app.Text("Label 2"),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-compact pf-m-green").
										Body(
											app.Span().
												Class("pf-c-label__content").
												Body(
													app.Span().
														Class("pf-c-label__icon").
														Body(
															app.I().
																Class("fas fa-fw fa-info-circle").
																Aria("hidden", true),
														),
													app.Text("Label 3"),
												),
										),
								),
						),
				),
		)
}
