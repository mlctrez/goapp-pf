package searchinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithSubmitButton struct {
	app.Compo
}

func (c *WithSubmitButton) Render() app.UI {
	return app.Div().
		Class("pf-c-search-input").
		Body(
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Div().
						Class("pf-c-search-input__bar").
						Body(
							app.Span().
								Class("pf-c-search-input__text").
								Body(
									app.Span().
										Class("pf-c-search-input__icon").
										Body(
											app.I().
												Class("fas fa-search fa-fw").
												Aria("hidden", true),
										),
									app.Input().
										Class("pf-c-search-input__text-input").
										Type("text").
										Placeholder("false").
										Aria("label", "Find by name"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("submit").
						Aria("label", "Search").
						Body(
							app.I().
								Class("fas fa-arrow-right").
								Aria("hidden", true),
						),
				),
		)
}
