package expandablesection

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hidden struct {
	app.Compo
}

func (c *Hidden) Render() app.UI {
	return app.Div().
		Class("pf-c-expandable-section").
		Body(
			app.Button().
				Type("button").
				Class("pf-c-expandable-section__toggle").
				Aria("expanded", "false").
				Body(
					app.Span().
						Class("pf-c-expandable-section__toggle-icon").
						Body(
							app.I().
								Class("fas fa-angle-right").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-expandable-section__toggle-text").
						Body(
							app.Text("Show more"),
						),
				),
			app.Div().
				Class("pf-c-expandable-section__content").
				Hidden(true).
				Body(
					app.Text("This content is visible only when the component is expanded."),
				),
		)
}
