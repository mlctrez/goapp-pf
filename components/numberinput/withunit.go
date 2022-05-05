package numberinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithUnit struct {
	app.Compo
}

func (c *WithUnit) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
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
								Value("90").
								Name("number-input-unit-name").
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
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-number-input").
				Body(
					app.Span().
						Class("pf-c-number-input__unit").
						Body(
							app.Text("$"),
						),
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
								Value("1.00").
								Name("number-input-unit2-name").
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
				),
		)
}
