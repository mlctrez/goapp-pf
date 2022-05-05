package emptystate

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExtraLarge struct {
	app.Compo
}

func (c *ExtraLarge) Render() app.UI {
	return app.Div().
		Class("pf-c-empty-state pf-m-xl").
		Body(
			app.Div().
				Class("pf-c-empty-state__content").
				Body(
					app.I().
						Class("fas fa-cubes pf-c-empty-state__icon").
						Aria("hidden", true),
					app.H1().
						Class("pf-c-title pf-m-4xl").
						Body(
							app.Text("Empty state"),
						),
					app.Div().
						Class("pf-c-empty-state__body").
						Body(
							app.Text("This represents an the empty state pattern in PatternFly 4. Hopefully it's simple enough to use but flexible enough to meet a variety of needs."),
						),
					app.Button().
						Class("pf-c-button pf-m-primary").
						Type("button").
						Body(
							app.Text("Primary action"),
						),
				),
		)
}
