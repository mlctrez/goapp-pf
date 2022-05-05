package simplelist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleListWithLinks struct {
	app.Compo
}

func (c *SimpleListWithLinks) Render() app.UI {
	return app.Div().
		Class("pf-c-simple-list").
		Body(
			app.Ul().
				Class("pf-c-simple-list__list").
				Body(
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.A().
								Class("pf-c-simple-list__item-link pf-m-current").
								Href("#").
								TabIndex(0).
								Body(
									app.Text("List item 1"),
								),
						),
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.A().
								Class("pf-c-simple-list__item-link").
								Href("#").
								TabIndex(0).
								Body(
									app.Text("List item 2"),
								),
						),
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.A().
								Class("pf-c-simple-list__item-link").
								Href("#").
								TabIndex(0).
								Body(
									app.Text("List item 3"),
								),
						),
				),
		)
}
