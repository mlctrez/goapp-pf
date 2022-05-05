package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type OptionSingleSelect struct {
	app.Compo
}

func (c *OptionSingleSelect) Render() app.UI {
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
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Option 1"),
														),
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
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Option 2"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.Button().
										Class("pf-c-menu__item pf-m-selected").
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
																Class("fas fa-fw fa-table").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Option 3"),
														),
													app.Span().
														Class("pf-c-menu__item-select-icon").
														Body(
															app.I().
																Class("fas fa-check").
																Aria("hidden", true),
														),
												),
											app.Span().
												Class("pf-c-menu__item-description").
												Body(
													app.Text("Description"),
												),
										),
								),
						),
				),
		)
}
