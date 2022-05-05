package expandablesection

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DetachedToggle struct {
	app.Compo
}

func (c *DetachedToggle) Render() app.UI {
	return app.Div().
		Class("pf-l-stack pf-m-gutter").
		Body(
			app.Div().
				Class("pf-l-stack__item").
				Body(
					app.Div().
						Class("pf-c-expandable-section pf-m-expanded pf-m-detached").
						Body(
							app.Div().
								Class("pf-c-expandable-section__content").
								ID("detached-toggle-content").
								Body(
									app.Text("This content is visible only when the component is expanded."),
								),
						),
				),
			app.Div().
				Class("pf-l-stack__item").
				Body(
					app.Div().
						Class("pf-c-expandable-section pf-m-expanded pf-m-detached").
						Body(
							app.Button().
								Type("button").
								Class("pf-c-expandable-section__toggle").
								Aria("expanded", true).
								Aria("controls", "detached-toggle-content").
								Body(
									app.Span().
										Class("pf-c-expandable-section__toggle-icon pf-m-expand-top").
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
						),
				),
		)
}
