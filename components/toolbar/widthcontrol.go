package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WidthControl struct {
	app.Compo
}

func (c *WidthControl) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-width-control").
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
										Style("", "").
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
