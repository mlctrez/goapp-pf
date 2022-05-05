package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroupAdvancedSearch struct {
	app.Compo
}

func (c *SearchInputGroupAdvancedSearch) Render() app.UI {
	return app.Div().
		Class("pf-c-input-group").
		Body(
			app.Div().
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
										Value("username:root firstname:ned").
										Aria("label", "Type to filter"),
								),
						),
					app.Div().
						Class("pf-c-text-input-group__utilities").
						Body(
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
		)
}
