package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonPrimaryAction struct {
	app.Compo
}

func (c *SplitButtonPrimaryAction) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-menu-toggle pf-m-split-button pf-m-action pf-m-disabled pf-m-primary").
				Body(
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Disabled(true).
						Body(
							app.Text("Action"),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-with-toggle-action-primary-disabled-example-toggle-button").
						Aria("label", "Menu toggle").
						Disabled(true).
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
				),
			app.Div().
				Class("pf-c-menu-toggle pf-m-split-button pf-m-action pf-m-primary").
				Body(
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Body(
							app.Text("Action"),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-with-toggle-action-primary-example-toggle-button").
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
				),
			app.Div().
				Class("pf-c-menu-toggle pf-m-expanded pf-m-split-button pf-m-action pf-m-primary").
				Body(
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Body(
							app.Text("Action"),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", true).
						ID("split-button-checkbox-with-toggle-action-primary-expanded-example-toggle-button").
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
				),
		)
}
