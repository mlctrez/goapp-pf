package labelgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CategoryRemovable struct {
	app.Compo
}

func (c *CategoryRemovable) Render() app.UI {
	return app.Div().
		Class("pf-c-label-group pf-m-category").
		Body(
			app.Div().
				Class("pf-c-label-group__main").
				Body(
					app.Span().
						Class("pf-c-label-group__label").
						Aria("hidden", true).
						ID("label-group-category-removable-label").
						Body(
							app.Text("Group label"),
						),
					app.Ul().
						Class("pf-c-label-group__list").
						Aria("role", "list").
						Aria("labelledby", "label-group-category-removable-label").
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
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-cyan").
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
													app.Text("Label 4"),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-orange").
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
													app.Text("Label 5"),
												),
										),
								),
							app.Li().
								Class("pf-c-label-group__list-item").
								Body(
									app.Span().
										Class("pf-c-label pf-m-red").
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
													app.Text("Label 6"),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-label-group__close").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("labelledby", "label-group-category-removable-button label-group-category-removable-label").
						Aria("label", "Close label group").
						ID("label-group-category-removable-button").
						Body(
							app.I().
								Class("fas fa-times-circle").
								Aria("hidden", true),
						),
				),
		)
}
