package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CallToAction struct {
	app.Compo
}

func (c *CallToAction) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-primary pf-m-display-lg").
				Type("button").
				Body(
					app.Text("Call to action"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-display-lg").
				Type("button").
				Body(
					app.Text("Call to action"),
				),
			app.Button().
				Class("pf-c-button pf-m-tertiary pf-m-display-lg").
				Type("button").
				Body(
					app.Text("Call to action"),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-display-lg").
				Type("button").
				Body(
					app.Text("Call to action"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-arrow-right").
								Aria("hidden", true),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-inline pf-m-display-lg").
				Type("button").
				Body(
					app.Text("Call to action"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-arrow-right").
								Aria("hidden", true),
						),
				),
		)
}
