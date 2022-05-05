package skeleton

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PercentageHeightModifiers struct {
	app.Compo
}

func (c *PercentageHeightModifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-skeleton pf-m-height-25"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-33"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-50"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-66"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-75"),
			app.Div().
				Class("pf-c-skeleton pf-m-height-100"),
		)
}
