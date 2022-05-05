package slider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Continuous struct {
	app.Compo
}

func (c *Continuous) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-slider").
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
								Class("pf-c-slider__thumb").
								Aria("role", "slider").
								Aria("valuemin", "0").
								Aria("valuemax", "100").
								Aria("valuenow", "50").
								Aria("label", "Value").
								TabIndex(0),
						),
				),
			app.Div().
				Class("pf-c-slider").
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
								TabIndex(0),
						),
				),
		)
}
