package pfswitch

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LabelAndCheck struct {
	app.Compo
}

func (c *LabelAndCheck) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Label().
				Class("pf-c-switch").
				For("switch-label-check-1").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-label-check-1").
						Aria("labelledby", "switch-label-check-1-on").
						Checked(true),
					app.Span().
						Class("pf-c-switch__toggle").
						Body(
							app.Span().
								Class("pf-c-switch__toggle-icon").
								Body(
									app.I().
										Class("fas fa-check").
										Aria("hidden", true),
								),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-on").
						ID("switch-label-check-1-on").
						Aria("hidden", true).
						Body(
							app.Text("Message when on"),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-off").
						ID("switch-label-check-1-off").
						Aria("hidden", true).
						Body(
							app.Text("Message when off"),
						),
				),
			app.Br(),
			app.Br(),
			app.Label().
				Class("pf-c-switch").
				For("switch-label-check-2").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-label-check-2").
						Aria("labelledby", "switch-label-check-2-off"),
					app.Span().
						Class("pf-c-switch__toggle").
						Body(
							app.Span().
								Class("pf-c-switch__toggle-icon").
								Body(
									app.I().
										Class("fas fa-check").
										Aria("hidden", true),
								),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-on").
						ID("switch-label-check-2-on").
						Aria("hidden", true).
						Body(
							app.Text("Message when on"),
						),
					app.Span().
						Class("pf-c-switch__label pf-m-off").
						ID("switch-label-check-2-off").
						Aria("hidden", true).
						Body(
							app.Text("Message when off"),
						),
				),
		)
}
