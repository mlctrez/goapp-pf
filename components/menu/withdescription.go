package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescription struct {
	app.Compo
}

func (c *WithDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-menu").
		Body(
			app.Div().
				Class("pf-c-menu__content").
				Body(
					app.Ul().
						Class("pf-c-menu__list").
						Aria("role", "menu").
						Body(
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-code-branch").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Action 1"),
														),
												),
											app.Span().
												Class("pf-c-menu__item-description").
												Body(
													app.Text("Description"),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item pf-m-disabled").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Disabled(true).
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-code-branch").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Action 2 disabled"),
														),
												),
											app.Span().
												Class("pf-c-menu__item-description").
												Body(
													app.Text("Description"),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item").
										Type("button").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-code-branch").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Action 3"),
														),
												),
											app.Span().
												Class("pf-c-menu__item-description").
												Body(
													app.Text("Nunc non ornare ex, et pretium dui. Duis nec augue at urna elementum blandit tincidunt eget metus. Aenean sed metus id urna dignissim interdum. Aenean vel nisl vitae arcu vehicula pulvinar eget nec turpis. Cras sit amet est est."),
												),
										),
								),
						),
				),
		)
}
