package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithIconimageAndText struct {
	app.Compo
}

func (c *WithIconimageAndText) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-menu-toggle pf-m-secondary").
				Type("button").
				Aria("expanded", "false").
				Body(
					app.Span().
						Class("pf-c-menu-toggle__icon").
						Body(
							app.I().
								Class("fas fa-cog").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Icon"),
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
				Class("pf-c-menu-toggle pf-m-secondary").
				Type("button").
				Aria("expanded", "false").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-menu-toggle__icon").
						Body(
							app.I().
								Class("fas fa-cog").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Icon"),
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
