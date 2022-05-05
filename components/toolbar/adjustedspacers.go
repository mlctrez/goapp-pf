package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdjustedSpacers struct {
	app.Compo
}

func (c *AdjustedSpacers) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-spacer-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__item pf-m-spacer-none").
								Body(
									app.Text("Item"),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-spacer-sm").
								Body(
									app.Text("Item"),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-spacer-md").
								Body(
									app.Text("Item"),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-spacer-lg").
								Body(
									app.Text("Item"),
								),
							app.Hr().
								Class("pf-c-divider pf-m-vertical"),
							app.Div().
								Class("pf-c-toolbar__item pf-m-spacer-none pf-m-spacer-sm-on-md pf-m-spacer-md-on-lg pf-m-spacer-lg-on-xl").
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
