package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FullHeight struct {
	app.Compo
}

func (c *FullHeight) Render() app.UI {
	return app.Button().
		Class("pf-c-menu-toggle pf-m-full-height").
		Type("button").
		Aria("expanded", "false").
		Body(
			app.Span().
				Class("pf-c-menu-toggle__text").
				Body(
					app.Text("Full height"),
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
		)
}
