package searchinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MatchWithResultCount struct {
	app.Compo
}

func (c *MatchWithResultCount) Render() app.UI {
	return app.Div().
		Class("pf-c-search-input").
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
								Aria("label", "Find by name").
								Value("John Doe"),
						),
					app.Span().
						Class("pf-c-search-input__utilities").
						Body(
							app.Span().
								Class("pf-c-search-input__count").
								Body(
									app.Span().
										Class("pf-c-badge pf-m-read").
										Body(
											app.Text("3"),
										),
								),
							app.Span().
								Class("pf-c-search-input__clear").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Clear").
										Body(
											app.I().
												Class("fas fa-times fa-fw").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
