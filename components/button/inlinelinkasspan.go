package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineLinkAsSpan struct {
	app.Compo
}

func (c *InlineLinkAsSpan) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), app.Span().
				Class("pf-c-button pf-m-link pf-m-inline").
				Aria("role", "button").
				TabIndex(0).
				Body(
					app.Text("This is long button text that needs to be a span so that it will wrap inline with the text around it."),
				),
			app.Text("Sed hendrerit nisi in cursus maximus. Ut malesuada nisi turpis, in condimentum velit elementum non."),
		)
}
