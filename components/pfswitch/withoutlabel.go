package pfswitch

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutLabel struct {
	app.Compo
}

func (c *WithoutLabel) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Label().
				Class("pf-c-switch").
				For("switch-with-icon-1").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-with-icon-1").
						Aria("label", "switch is off").
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
				),
			app.Br(),
			app.Br(),
			app.Label().
				Class("pf-c-switch").
				For("switch-with-icon-2").
				Body(
					app.Input().
						Class("pf-c-switch__input").
						Type("checkbox").
						ID("switch-with-icon-2").
						Aria("label", "switch is off"),
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
				),
		)
}
