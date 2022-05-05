package pfswitch

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Label().
				Class("pf-c-switch").
				For("switch-with-label-1").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-with-label-1").
						Aria("labelledby", "switch-with-label-1-on").
						Checked(true),
					app.Span().
						Class("pf-c-switch__toggle"),
					app.Span().
						Class("pf-c-switch__label pf-m-on").
						ID("switch-with-label-1-on").
						Aria("hidden", true).
						Body(
							app.Text("Message when on"),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-off").
						ID("switch-with-label-1-off").
						Aria("hidden", true).
						Body(
							app.Text("Message when off"),
						),
				),
			app.Br(),
			app.Br(),
			app.Label().
				Class("pf-c-switch").
				For("switch-with-label-2").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-with-label-2").
						Aria("labelledby", "switch-with-label-2-on"),
					app.Span().
						Class("pf-c-switch__toggle"),
					app.Span().
						Class("pf-c-switch__label pf-m-on").
						ID("switch-with-label-2-on").
						Aria("hidden", true).
						Body(
							app.Text("Message when on"),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-off").
						ID("switch-with-label-2-off").
						Aria("hidden", true).
						Body(
							app.Text("Message when off"),
						),
				),
		)
}
