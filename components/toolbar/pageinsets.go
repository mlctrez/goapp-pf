package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PageInsets struct {
	app.Compo
}

func (c *PageInsets) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar pf-m-page-insets").
		ID("toolbar-page-insets-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Text("Item"),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Text("Item"),
										),
								),
							app.Hr().
								Class("pf-c-divider pf-m-vertical"),
							app.Div().
								Class("pf-c-toolbar__item").
								Body(
									app.Text("Item"),
								),
							app.Div().
								Class("pf-c-toolbar__item").
								Body(
									app.Text("Item"),
								),
						),
				),
		)
}
