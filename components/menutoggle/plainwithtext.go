package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PlainWithText struct {
	app.Compo
}

func (c *PlainWithText) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-menu-toggle pf-m-plain pf-m-text").
				Type("button").
				Aria("expanded", "false").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Disabled"),
						),
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
			app.Button().
				Class("pf-c-menu-toggle pf-m-plain pf-m-text").
				Type("button").
				Aria("expanded", "false").
				Body(
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Custom text"),
						),
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
		)
}
