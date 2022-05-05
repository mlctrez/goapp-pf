package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Insets struct {
	app.Compo
}

func (c *Insets) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar pf-m-inset-none pf-m-inset-md-on-md pf-m-inset-2xl-on-lg").
		ID("toolbar-inset-example").
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
