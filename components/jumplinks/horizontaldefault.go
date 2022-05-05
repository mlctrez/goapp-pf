package jumplinks

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontalDefault struct {
	app.Compo
}

func (c *HorizontalDefault) Render() app.UI {
	return app.Nav().
		Class("pf-c-jump-links").
		Body(
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
		)
}
