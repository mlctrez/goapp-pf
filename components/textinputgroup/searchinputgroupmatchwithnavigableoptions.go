package textinputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SearchInputGroupMatchWithNavigableOptions struct {
	app.Compo
}

func (c *SearchInputGroupMatchWithNavigableOptions) Render() app.UI {
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
							app.Text("1 / 3"),
						),
					app.Div().
						Class("pf-c-text-input-group__group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Disabled(true).
								Aria("label", "Next").
								Body(
									app.I().
										Class("fas fa-angle-up fa-fw").
										Aria("hidden", true),
								),
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Next").
								Body(
									app.I().
										Class("fas fa-angle-down fa-fw").
										Aria("hidden", true),
								),
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
