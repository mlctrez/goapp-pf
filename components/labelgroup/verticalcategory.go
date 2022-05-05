package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticalCategory struct {
	app.Compo
}

func (c *VerticalCategory) Render() app.UI {
	return app.Div().
		Class("pf-c-label-group pf-m-vertical pf-m-category").
		Body(
			app.Div().
				Class("pf-c-label-group__main").
				Body(
					app.Span().
						Class("pf-c-label-group__label").
						Aria("hidden", true).
						ID("label-group-vertical-category-label").
						Body(
							app.Text("Group label"),
						),
					app.Ul().
						Class("pf-c-label-group__list").
						Aria("role", "list").
						Aria("labelledby", "label-group-vertical-category-label").
						Body(
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label").
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
										Class("pf-c-label pf-m-blue").
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
										Class("pf-c-label pf-m-green").
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
