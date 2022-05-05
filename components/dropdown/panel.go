package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Panel struct {
	app.Compo
}

func (c *Panel) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("dropdown-panel-button").
				Aria("expanded", true).
				Type("button").
				Body(
					app.Span().
						Class("pf-c-dropdown__toggle-text").
						Body(
							app.Text("Expanded dropdown"),
						),
					app.Span().
						Class("pf-c-dropdown__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-dropdown__menu").
				Body(
					app.Text("[Panel contents here]"),
				),
		)
}
