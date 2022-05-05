package skeleton

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PercentageWidthModifiers struct {
	app.Compo
}

func (c *PercentageWidthModifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-skeleton pf-m-width-25"),
			app.Br(),
			app.Div().
				Class("pf-c-skeleton pf-m-width-33"),
			app.Br(),
			app.Div().
				Class("pf-c-skeleton pf-m-width-50"),
			app.Br(),
			app.Div().
				Class("pf-c-skeleton pf-m-width-66"),
			app.Br(),
			app.Div().
				Class("pf-c-skeleton pf-m-width-75"),
			app.Br(),
			app.Div().
				Class("pf-c-skeleton"),
		)
}
