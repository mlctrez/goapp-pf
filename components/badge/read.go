package badge

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Read struct {
	app.Compo
}

func (c *Read) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-badge pf-m-read").
				Body(
					app.Text("7"),
				),
			app.Span().
				Class("pf-c-badge pf-m-read").
				Body(
					app.Text("24"),
				),
			app.Span().
				Class("pf-c-badge pf-m-read").
				Body(
					app.Text("240"),
				),
			app.Span().
				Class("pf-c-badge pf-m-read").
				Body(
					app.Text("999+"),
				),
		)
}
