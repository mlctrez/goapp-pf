package emptystate

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithPrimaryElement struct {
	app.Compo
}

func (c *WithPrimaryElement) Render() app.UI {
	return app.Div().
		Class("pf-c-empty-state").
		Body(
			app.Div().
				Class("pf-c-empty-state__content").
				Body(
					app.I().
						Class("fas fa-cubes pf-c-empty-state__icon").
						Aria("hidden", true),
					app.H1().
						Class("pf-c-title pf-m-lg").
						Body(
							app.Text("Empty State"),
						),
					app.Div().
						Class("pf-c-empty-state__body").
						Body(
							app.Text("This represents an the empty state pattern in PatternFly 4. Hopefully it's simple enough to use but flexible enough to meet a variety of needs."),
						),
					app.Div().
						Class("pf-c-empty-state__primary").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link").
								Type("button").
								Body(
									app.Text("Action buttons"),
								),
						),
				),
		)
}
