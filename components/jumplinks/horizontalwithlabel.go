package jumplinks

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontalWithLabel struct {
	app.Compo
}

func (c *HorizontalWithLabel) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Nav().
				Class("pf-c-jump-links").
				Aria("label", "Jump to section").
				Body(
					app.Div().
						Class("pf-c-jump-links__main").
						Body(
							app.Div().
								Class("pf-c-jump-links__header").
								Body(
									app.Div().
										Class("pf-c-jump-links__label").
										Body(
											app.Text("Jump to section"),
										),
								),
							app.Ul().
								Class("pf-c-jump-links__list").
								Body(
									app.Li().
										Class("pf-c-jump-links__item").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Inactive section"),
														),
												),
										),
									app.Li().
										Class("pf-c-jump-links__item pf-m-current").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Active section"),
														),
												),
										),
									app.Li().
										Class("pf-c-jump-links__item").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Inactive section"),
														),
												),
										),
								),
						),
				),
			app.Br(),
			app.Nav().
				Class("pf-c-jump-links pf-m-center").
				Aria("label", "Jump to section").
				Body(
					app.Div().
						Class("pf-c-jump-links__main").
						Body(
							app.Div().
								Class("pf-c-jump-links__header").
								Body(
									app.Div().
										Class("pf-c-jump-links__label").
										Body(
											app.Text("Jump to section"),
										),
								),
							app.Ul().
								Class("pf-c-jump-links__list").
								Body(
									app.Li().
										Class("pf-c-jump-links__item").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Inactive section"),
														),
												),
										),
									app.Li().
										Class("pf-c-jump-links__item pf-m-current").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Active section"),
														),
												),
										),
									app.Li().
										Class("pf-c-jump-links__item").
										Body(
											app.A().
												Class("pf-c-jump-links__link").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-jump-links__link-text").
														Body(
															app.Text("Inactive section"),
														),
												),
										),
								),
						),
				),
		)
}
