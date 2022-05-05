package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithLinks struct {
	app.Compo
}

func (c *WithLinks) Render() app.UI {
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
									app.A().
										Class("pf-c-menu__item").
										Href("#").
										Aria("role", "menuitem").
										Target("_blank").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Link 1"),
														),
													app.Span().
														Class("pf-c-menu__item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.A().
										Class("pf-c-menu__item").
										Href("#").
										Aria("role", "menuitem").
										Target("_blank").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Link 2"),
														),
													app.Span().
														Class("pf-c-menu__item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item").
								Aria("role", "none").
								Body(
									app.A().
										Class("pf-c-menu__item").
										Href("#").
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Link 3"),
														),
												),
										),
								),
						),
				),
		)
}
