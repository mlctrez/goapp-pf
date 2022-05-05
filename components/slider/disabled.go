package slider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-slider pf-m-disabled").
				Style("--pf-c-slider--value", " 62.5%;").
				Body(
					app.Div().
						Class("pf-c-slider__main").
						Body(
							app.Div().
								Class("pf-c-slider__rail").
								Body(
									app.Div().
										Class("pf-c-slider__rail-track"),
								),
							app.Div().
								Class("pf-c-slider__steps").
								Aria("hidden", true).
								Body(
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 0%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("0"),
												),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 12.5%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 25%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("2"),
												),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 37.5%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 50%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("4"),
												),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 62.5%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step").
										Style("--pf-c-slider__step--Left", " 75%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("6"),
												),
										),
									app.Div().
										Class("pf-c-slider__step").
										Style("--pf-c-slider__step--Left", " 87.5%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step").
										Style("--pf-c-slider__step--Left", " 100%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("8"),
												),
										),
								),
							app.Div().
								Class("pf-c-slider__thumb").
								Aria("role", "slider").
								Aria("valuemin", "0").
								Aria("valuemax", "8").
								Aria("valuenow", "5").
								Aria("label", "Value").
								Aria("disabled", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-slider pf-m-disabled").
				Style("--pf-c-slider--value", " 50%;").
				Body(
					app.Div().
						Class("pf-c-slider__main").
						Body(
							app.Div().
								Class("pf-c-slider__rail").
								Body(
									app.Div().
										Class("pf-c-slider__rail-track"),
								),
							app.Div().
								Class("pf-c-slider__steps").
								Aria("hidden", true).
								Body(
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 0%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("0%"),
												),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 25%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step pf-m-active").
										Style("--pf-c-slider__step--Left", " 50%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("50%"),
												),
										),
									app.Div().
										Class("pf-c-slider__step").
										Style("--pf-c-slider__step--Left", " 75%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
										),
									app.Div().
										Class("pf-c-slider__step").
										Style("--pf-c-slider__step--Left", " 100%;").
										Body(
											app.Div().
												Class("pf-c-slider__step-tick"),
											app.Div().
												Class("pf-c-slider__step-label").
												Body(
													app.Text("100%"),
												),
										),
								),
							app.Div().
								Class("pf-c-slider__thumb").
								Aria("role", "slider").
								Aria("valuemin", "0").
								Aria("valuemax", "100").
								Aria("valuenow", "50").
								Aria("label", "Value").
								Aria("disabled", true),
						),
					app.Div().
						Class("pf-c-slider__value").
						Body(
							app.Div().
								Class("pf-c-input-group").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Disabled(true).
										Type("number").
										Value("50").
										Aria("label", "Slider value input"),
									app.Span().
										Class("pf-c-input-group__text pf-m-plain").
										Body(
											app.Text("%"),
										),
								),
						),
				),
		)
}
