package skeleton

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TextModifiers struct {
	app.Compo
}

func (c *TextModifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Text("--pf-global--FontSize--4xl"), app.Div().
				Class("pf-c-skeleton pf-m-text-4xl"),
			app.Br(),
			app.Text("--pf-global--FontSize--3xl"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-3xl"),
			app.Br(),
			app.Text("--pf-global--FontSize--2xl"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-2xl"),
			app.Br(),
			app.Text("--pf-global--FontSize--xl"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-xl"),
			app.Br(),
			app.Text("--pf-global--FontSize--lg"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-lg"),
			app.Br(),
			app.Text("--pf-global--FontSize--md"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-md"),
			app.Br(),
			app.Text("--pf-global--FontSize--sm"),
			app.Div().
				Class("pf-c-skeleton pf-m-text-sm"),
		)
}
