package simplelist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleList struct {
	app.Compo
}

func (c *SimpleList) Render() app.UI {
	return app.Div().
		Class("pf-c-simple-list").
		Body(
			app.Ul().
				Class("pf-c-simple-list__list").
				Body(
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.Button().
								Class("pf-c-simple-list__item-link pf-m-current").
								Type("button").
								Body(
									app.Text("List item 1"),
								),
						),
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.Button().
								Class("pf-c-simple-list__item-link").
								Type("button").
								Body(
									app.Text("List item 2"),
								),
						),
					app.Li().
						Class("pf-c-simple-list__item").
						Body(
							app.Button().
								Class("pf-c-simple-list__item-link").
								Type("button").
								Body(
									app.Text("List item 3"),
								),
						),
				),
		)
}
