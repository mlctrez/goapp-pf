package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Button().
		Class("pf-c-menu-toggle").
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
		)
}
