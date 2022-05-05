package simplelist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GroupedList struct {
	app.Compo
}

func (c *GroupedList) Render() app.UI {
	return app.Div().
		Class("pf-c-simple-list").
		Body(
			app.Section().
				Class("pf-c-simple-list__section").
				Body(
					app.H2().
						Class("pf-c-simple-list__title").
						Body(
							app.Text("Title"),
						),
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
							app.Li().
								Class("pf-c-simple-list__item").
								Body(
									app.Button().
										Class("pf-c-simple-list__item-link").
										Type("button").
										Body(
											app.Text("List item 4"),
										),
								),
						),
				),
			app.Section().
				Class("pf-c-simple-list__section").
				Body(
					app.H2().
						Class("pf-c-simple-list__title").
						Body(
							app.Text("Title"),
						),
					app.Ul().
						Class("pf-c-simple-list__list").
						Body(
							app.Li().
								Class("pf-c-simple-list__item").
								Body(
									app.Button().
										Class("pf-c-simple-list__item-link").
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
							app.Li().
								Class("pf-c-simple-list__item").
								Body(
									app.Button().
										Class("pf-c-simple-list__item-link").
										Type("button").
										Body(
											app.Text("List item 4"),
										),
								),
						),
				),
		)
}
