package numberinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithUnitAndUpperThresholdReached struct {
	app.Compo
}

func (c *WithUnitAndUpperThresholdReached) Render() app.UI {
	return app.Div().
		Class("pf-c-number-input").
		Body(
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Minus").
						Body(
							app.Span().
								Class("pf-c-number-input__icon").
								Body(
									app.I().
										Class("fas fa-minus").
										Aria("hidden", true),
								),
						),
					app.Input().
						Class("pf-c-form-control").
						Type("number").
						Value("100").
						Max("100").
						Name("number-input-unit-upper-threshold-name").
						Aria("label", "Number input"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Plus").
						Disabled(true).
						Body(
							app.Span().
								Class("pf-c-number-input__icon").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
			app.Span().
				Class("pf-c-number-input__unit").
				Body(
					app.Text("%"),
				),
		)
}
