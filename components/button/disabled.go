package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Primary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Secondary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-danger").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Secondary danger disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-tertiary").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Tertiary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-danger").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Danger disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-warning").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Warning disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-button__icon pf-m-start").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
					app.Text("Link disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-danger").
				Type("button").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-button__icon pf-m-start").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
					app.Text("Link danger disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-inline").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Inline link disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Remove").
				Disabled(true).
				Body(
					app.I().
						Class("fas fa-times").
						Aria("hidden", true),
				),
			app.Button().
				Class("pf-c-button pf-m-control").
				Type("button").
				Disabled(true).
				Body(
					app.Text("Control disabled"),
				),
		)
}
