package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Variations struct {
	app.Compo
}

func (c *Variations) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Body(
					app.Text("Primary"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary").
				Type("button").
				Body(
					app.Text("Secondary"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-danger").
				Type("button").
				Body(
					app.Text("Secondary danger"),
				),
			app.Button().
				Class("pf-c-button pf-m-tertiary").
				Type("button").
				Body(
					app.Text("Tertiary"),
				),
			app.Button().
				Class("pf-c-button pf-m-danger").
				Type("button").
				Body(
					app.Text("Danger"),
				),
			app.Button().
				Class("pf-c-button pf-m-warning").
				Type("button").
				Body(
					app.Text("Warning"),
				),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Body(
					app.Span().
						Class("pf-c-button__icon pf-m-start").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
					app.Text("Link"),
				),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Body(
					app.Text("Link"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-danger").
				Type("button").
				Body(
					app.Span().
						Class("pf-c-button__icon pf-m-start").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
					app.Text("Link danger"),
				),
			app.Button().
				Class("pf-c-button pf-m-inline pf-m-link").
				Type("button").
				Body(
					app.Text("Inline link"),
				),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Remove").
				Body(
					app.I().
						Class("fas fa-times").
						Aria("hidden", true),
				),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-button pf-m-control").
				Type("button").
				Body(
					app.Text("Control"),
				),
			app.Button().
				Class("pf-c-button pf-m-control").
				Type("button").
				Aria("label", "Copy input").
				Body(
					app.I().
						Class("fas fa-copy").
						Aria("hidden", true),
				),
			app.Br(),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-button pf-m-primary pf-m-small").
				Type("button").
				Body(
					app.Text("Primary"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-small").
				Type("button").
				Body(
					app.Text("Secondary"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-danger pf-m-small").
				Type("button").
				Body(
					app.Text("Secondary danger"),
				),
			app.Button().
				Class("pf-c-button pf-m-tertiary pf-m-small").
				Type("button").
				Body(
					app.Text("Tertiary"),
				),
			app.Button().
				Class("pf-c-button pf-m-danger pf-m-small").
				Type("button").
				Body(
					app.Text("Danger"),
				),
			app.Button().
				Class("pf-c-button pf-m-warning pf-m-small").
				Type("button").
				Body(
					app.Text("Warning"),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-small").
				Type("button").
				Body(
					app.Text("Link"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-danger pf-m-small").
				Type("button").
				Body(
					app.Text("Link danger"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-plus-circle").
								Aria("hidden", true),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-inline pf-m-link pf-m-small").
				Type("button").
				Body(
					app.Text("Inline link"),
				),
			app.Button().
				Class("pf-c-button pf-m-control pf-m-small").
				Type("button").
				Body(
					app.Text("Control"),
				),
		)
}
