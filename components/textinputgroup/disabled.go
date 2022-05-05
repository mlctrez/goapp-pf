package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("pf-c-text-input-group pf-m-disabled").
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
								Value("Disabled").
								Disabled(true).
								Aria("label", "Type to filter"),
						),
				),
		)
}
