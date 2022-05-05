package numberinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VaryingSizes struct {
	app.Compo
}

func (c *VaryingSizes) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-number-input").
				Style("--pf-c-number-input--c-form-control--width-chars", " 1;").
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
								Value("1").
								Name("number-input-sizes-name").
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
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-number-input").
				Style("--pf-c-number-input--c-form-control--width-chars", " 10;").
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
								Value("1234567890").
								Name("number-input-sizes2-name").
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
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-number-input").
				Style("--pf-c-number-input--c-form-control--width-chars", " 5;").
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
								Value("5").
								Name("number-input-sizes3-name").
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
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-number-input").
				Style("--pf-c-number-input--c-form-control--width-chars", " 5;").
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
								Value("12345").
								Name("number-input-sizes4-name").
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
