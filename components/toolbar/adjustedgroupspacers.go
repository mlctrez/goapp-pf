package toolbar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdjustedGroupSpacers struct {
	app.Compo
}

func (c *AdjustedGroupSpacers) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		ID("toolbar-group-spacer-example").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section").
						Body(
							app.Div().
								Class("pf-c-toolbar__group pf-m-space-items-lg").
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
								Class("pf-c-toolbar__group pf-m-space-items-none pf-m-space-items-sm-on-md pf-m-space-items-md-on-lg pf-m-space-items-lg-on-xl").
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
						),
				),
		)
}
