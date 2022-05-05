package expandablesection

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expanded struct {
	app.Compo
}

func (c *Expanded) Render() app.UI {
	return app.Div().
		Class("pf-c-expandable-section pf-m-expanded").
		Body(
			app.Button().
				Type("button").
				Class("pf-c-expandable-section__toggle").
				Aria("expanded", true).
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
							app.Text("Show less"),
						),
				),
			app.Div().
				Class("pf-c-expandable-section__content").
				Body(
					app.Text("This content is visible only when the component is expanded."),
				),
		)
}
