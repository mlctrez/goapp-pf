package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expanded struct {
	app.Compo
}

func (c *Expanded) Render() app.UI {
	return app.Button().
		Class("pf-c-menu-toggle pf-m-expanded").
		Type("button").
		Aria("expanded", true).
		Body(
			app.Span().
				Class("pf-c-menu-toggle__text").
				Body(
					app.Text("Expanded"),
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
