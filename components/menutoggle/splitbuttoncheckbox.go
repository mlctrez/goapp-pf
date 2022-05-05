package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonCheckbox struct {
	app.Compo
}

func (c *SplitButtonCheckbox) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-menu-toggle pf-m-split-button pf-m-disabled").
				Body(
					app.Label().
						Class("pf-c-check pf-m-standalone").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-disabled-example-input").
								Name("split-button-checkbox-disabled-example-input").
								Aria("label", "Standalone input").
								Disabled(true),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-disabled-example-toggle-button").
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
				Class("pf-c-menu-toggle pf-m-split-button").
				Body(
					app.Label().
						Class("pf-c-check pf-m-standalone").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-example-input").
								Name("split-button-checkbox-example-input").
								Aria("label", "Standalone input"),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", "false").
						ID("split-button-checkbox-example-toggle-button").
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
				Class("pf-c-menu-toggle pf-m-expanded pf-m-split-button").
				Body(
					app.Label().
						Class("pf-c-check pf-m-standalone").
						Body(
							app.Input().
								Class("pf-c-check__input").
								Type("checkbox").
								ID("split-button-checkbox-expanded-example-input").
								Name("split-button-checkbox-expanded-example-input").
								Aria("label", "Standalone input"),
						),
					app.Button().
						Class("pf-c-menu-toggle__button").
						Type("button").
						Aria("expanded", true).
						ID("split-button-checkbox-expanded-example-toggle-button").
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
