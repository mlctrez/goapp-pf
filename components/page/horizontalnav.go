package page

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontalNav struct {
	app.Compo
}

func (c *HorizontalNav) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.Div().
						Class("pf-c-page__header-brand").
						Body(
							app.A().
								Href("#").
								Class("pf-c-page__header-brand-link").
								Body(
									app.Text("Logo"),
								),
						),
					app.Div().
						Class("pf-c-page__header-nav").
						Body(
							app.Text("pf-c-nav"),
						),
					app.Div().
						Class("pf-c-page__header-tools").
						Body(
							app.Text("header-tools"),
						),
				),
			app.Main().
				Class("pf-c-page__main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-dark-100"),
					app.Section().
						Class("pf-c-page__main-section pf-m-dark-200"),
					app.Section().
						Class("pf-c-page__main-section pf-m-light"),
					app.Section().
						Class("pf-c-page__main-section"),
				),
		)
}
