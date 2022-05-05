package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroup struct {
	app.Compo
}

func (c *SearchInputGroup) Render() app.UI {
	return app.Div().
		Class("pf-c-text-input-group").
		Body(
			app.Div().
				Class("pf-c-text-input-group__main pf-m-icon").
				Body(
					app.Span().
						Class("pf-c-text-input-group__text").
						Body(
							app.Span().
								Class("pf-c-text-input-group__icon").
								Body(
									app.I().
										Class("fas fa-fw fa-search"),
								),
							app.Input().
								Class("pf-c-text-input-group__text-input").
								Type("text").
								Value(true).
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
				),
		)
}
