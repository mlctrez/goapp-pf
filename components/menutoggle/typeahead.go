package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Typeahead struct {
	app.Compo
}

func (c *Typeahead) Render() app.UI {
	return app.Div().
		Class("pf-c-menu-toggle pf-m-typeahead pf-m-full-width").
		Body(
			app.Div().
				Class("pf-c-text-input-group pf-m-plain").
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
				Class("pf-c-menu-toggle__button").
				Type("button").
				Aria("expanded", "false").
				ID("typeahead-example-toggle-button").
				Aria("label", "Menu toggle").
				Body(
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
		)
}
