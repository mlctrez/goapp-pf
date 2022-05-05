package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithIcons struct {
	app.Compo
}

func (c *WithIcons) Render() app.UI {
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
															app.Text("From Git"),
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
														Class("pf-c-menu__item-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-layer-group").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Container image"),
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
														Class("pf-c-menu__item-icon").
														Body(
															app.I().
																Class("fas fa-fw fa-cube").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("From DockerFile"),
														),
												),
										),
								),
						),
				),
		)
}
