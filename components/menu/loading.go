package menu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Loading struct {
	app.Compo
}

func (c *Loading) Render() app.UI {
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
															app.Text("Action"),
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
															app.Text("Link"),
														),
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
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Disabled action"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item pf-m-disabled").
								Aria("role", "none").
								Body(
									app.A().
										Class("pf-c-menu__item").
										Href("#").
										Aria("disabled", true).
										TabIndex(-1).
										Aria("role", "menuitem").
										Body(
											app.Span().
												Class("pf-c-menu__item-main").
												Body(
													app.Span().
														Class("pf-c-menu__item-text").
														Body(
															app.Text("Disabled link"),
														),
												),
										),
								),
							app.Li().
								Class("pf-c-menu__list-item pf-m-loading").
								Aria("role", "none").
								Body(
									app.Span().
										Class("pf-c-spinner pf-m-lg").
										Aria("role", "progressbar").
										Aria("label", "Loading items").
										Body(
											app.Span().
												Class("pf-c-spinner__clipper"),
											app.Span().
												Class("pf-c-spinner__lead-ball"),
											app.Span().
												Class("pf-c-spinner__tail-ball"),
										),
								),
						),
				),
		)
}
