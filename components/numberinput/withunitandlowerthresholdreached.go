package numberinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithUnitAndLowerThresholdReached struct {
	app.Compo
}

func (c *WithUnitAndLowerThresholdReached) Render() app.UI {
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
						Disabled(true).
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
						Value("0").
						Min("0").
						Name("number-input-unit-lower-threshold-name").
						Aria("label", "Number input"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Plus").
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
