package skeleton

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StaticHeightWidthAndShapeModifiers struct {
	app.Compo
}

func (c *StaticHeightWidthAndShapeModifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Text("Small circle"), app.Div().
				Class("pf-c-skeleton pf-m-circle pf-m-width-sm"),
			app.Br(),
			app.Text("Medium circle"),
			app.Div().
				Class("pf-c-skeleton pf-m-circle pf-m-width-md"),
			app.Br(),
			app.Text("Large circle"),
			app.Div().
				Class("pf-c-skeleton pf-m-circle pf-m-width-lg"),
			app.Br(),
			app.Text("Small square"),
			app.Div().
				Class("pf-c-skeleton pf-m-square pf-m-width-sm"),
			app.Br(),
			app.Text("Medium square"),
			app.Div().
				Class("pf-c-skeleton pf-m-square pf-m-width-md"),
			app.Br(),
			app.Text("Large square"),
			app.Div().
				Class("pf-c-skeleton pf-m-square pf-m-width-lg"),
			app.Br(),
			app.Text("Small rectangle"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-sm pf-m-width-md"),
			app.Br(),
			app.Text("Medium rectangle"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-md pf-m-width-lg"),
			app.Br(),
			app.Text("Large rectangle"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-lg"),
		)
}
