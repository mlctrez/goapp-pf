package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Count struct {
	app.Compo
}

func (c *Count) Render() app.UI {
	return app.Button().
		Class("pf-c-menu-toggle").
		Type("button").
		Aria("expanded", "false").
		Body(
			app.Span().
				Class("pf-c-menu-toggle__text").
				Body(
					app.Text("Count"),
				),
			app.Span().
				Class("pf-c-menu-toggle__count").
				Body(
					app.Span().
						Class("pf-c-badge pf-m-unread").
						Body(
							app.Text("4 selected"),
						),
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
