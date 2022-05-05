package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroupMatchWithResultCount struct {
	app.Compo
}

func (c *SearchInputGroupMatchWithResultCount) Render() app.UI {
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
								Value("John Doe").
								Aria("label", "Type to filter"),
						),
				),
			app.Div().
				Class("pf-c-text-input-group__utilities").
				Body(
					app.Span().
						Class("pf-c-badge pf-m-read").
						Body(
							app.Text("3"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Clear input").
						Body(
							app.I().
								Class("fas fa-times fa-fw").
								Aria("hidden", true),
						),
				),
		)
}
