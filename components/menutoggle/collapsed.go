package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Collapsed struct {
	app.Compo
}

func (c *Collapsed) Render() app.UI {
	return app.Button().
		Class("pf-c-menu-toggle").
		Type("button").
		Aria("expanded", "false").
		Body(
			app.Span().
				Class("pf-c-menu-toggle__text").
				Body(
					app.Text("Collapsed"),
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
