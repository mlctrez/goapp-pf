package slider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ThumbValueInput struct {
	app.Compo
}

func (c *ThumbValueInput) Render() app.UI {
	return app.Div().
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
					app.Div().
						Class("pf-c-slider__value pf-m-floating").
						Body(
							app.Div().
								Class("pf-c-input-group").
								Body(
									app.Input().
										Class("pf-c-form-control").
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
