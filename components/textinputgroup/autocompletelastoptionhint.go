package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AutocompleteLastOptionHint struct {
	app.Compo
}

func (c *AutocompleteLastOptionHint) Render() app.UI {
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
								Class("pf-c-text-input-group__text-input pf-m-hint").
								Type("text").
								Value("appleseed").
								Disabled(true).
								Aria("hidden", true),
							app.Input().
								Class("pf-c-text-input-group__text-input").
								Type("text").
								Value("apples").
								Aria("label", "Type to filter"),
						),
				),
		)
}
