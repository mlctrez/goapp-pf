package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-text-input-group").
		Body(
			app.Div().
				Class("pf-c-text-input-group__main").
				Body(
					app.Span().
						Class("pf-c-text-input-group__text").
						Body(
							app.Input().
								Class("pf-c-text-input-group__text-input").
								Type("text").
								Value(true).
								Aria("label", "Type to filter"),
						),
				),
		)
}
