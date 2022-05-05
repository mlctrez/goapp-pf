package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonSecondary struct {
	app.Compo
}

func (c *SplitButtonSecondary) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-menu-toggle pf-m-split-button pf-m-disabled pf-m-secondary").
				Body(
					app.Label().
						Class("pf-c-check").
						For("split-button-checkbox-secondary-disabled-example-input").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-secondary-disabled-example-input").
								Name("split-button-checkbox-secondary-disabled-example-input").
								Disabled(true),
							app.Span().
								Class("pf-c-check__label pf-m-disabled").
								Body(
									app.Text("10 selected"),
								),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-secondary-disabled-example-toggle-button").
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
				Class("pf-c-menu-toggle pf-m-split-button pf-m-secondary").
				Body(
					app.Label().
						Class("pf-c-check").
						For("split-button-checkbox-secondary-example-input").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-secondary-example-input").
								Name("split-button-checkbox-secondary-example-input"),
							app.Span().
								Class("pf-c-check__label").
								Body(
									app.Text("10 selected"),
								),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-secondary-example-toggle-button").
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
				Class("pf-c-menu-toggle pf-m-expanded pf-m-split-button pf-m-secondary").
				Body(
					app.Label().
						Class("pf-c-check").
						For("split-button-checkbox-secondary-expanded-example-input").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-secondary-expanded-example-input").
								Name("split-button-checkbox-secondary-expanded-example-input"),
							app.Span().
								Class("pf-c-check__label").
								Body(
									app.Text("10 selected"),
								),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", true).
						ID("split-button-checkbox-secondary-expanded-example-toggle-button").
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
