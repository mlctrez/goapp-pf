package ui

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Page struct {
	app.Compo
	Body []app.UI
}

func (c *Page) Render() app.UI {
	return app.Div().Class("pf-c-page").
		Body(
			app.Header().Class("pf-c-page__header").
				Body(app.Div().Class("pf-c-page__header-brand").
					Body(
						app.A().Href("https://github.com/mlctrez/goapp-pf").
							Class("pf-c-page__header-brand-link").Body(app.Text("GoApp-pf")),
					),
				),
			app.Main().
				Class("pf-c-page__main").Style("height", "80vh").TabIndex(-1).Body(
				app.Section().Class("pf-c-page__main-section pf-m-limit-width pf-m-align-center").Body(
					app.Div().Class("pf-c-page__main-body").Body(
						app.Div().Class("pf-c-card").Body(
							app.Div().Class("pf-c-card__body").
								Body(c.Body...),
						),
					),
				),
			),
		)
}
