package searchinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdvancedSearch struct {
	app.Compo
}

func (c *AdvancedSearch) Render() app.UI {
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
										Aria("label", "username:admin firstname:joe").
										Value("username:root firstname:ned"),
								),
							app.Span().
								Class("pf-c-search-input__utilities").
								Body(
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
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("expanded", "false").
						Aria("label", "Advanced search").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
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
