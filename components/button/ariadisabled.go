package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AriaDisabled struct {
	app.Compo
}

func (c *AriaDisabled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-primary pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Primary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Secondary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-secondary pf-m-danger pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Secondary danger disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-tertiary pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Tertiary disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-danger pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Danger disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-warning pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Warning disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-link pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
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
				Class("pf-c-button pf-m-link pf-m-danger pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
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
				Class("pf-c-button pf-m-link pf-m-inline pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Inline link disabled"),
				),
			app.Button().
				Class("pf-c-button pf-m-plain pf-m-aria-disabled").
				Type("button").
				Aria("label", "Remove").
				Aria("disabled", true).
				Body(
					app.I().
						Class("fas fa-times").
						Aria("hidden", true),
				),
			app.Button().
				Class("pf-c-button pf-m-control pf-m-aria-disabled").
				Type("button").
				Aria("disabled", true).
				Body(
					app.Text("Control disabled"),
				),
		)
}
