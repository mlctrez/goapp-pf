package slider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Actions struct {
	app.Compo
}

func (c *Actions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-slider").
				Style("--pf-c-slider--value", " 50%;").
				Body(
					app.Div().
						Class("pf-c-slider__actions").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Minus").
								Body(
									app.I().
										Class("fas fa-fw fa-minus").
										Aria("hidden", true),
								),
						),
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
					app.Div().
						Class("pf-c-slider__actions").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Plus").
								Body(
									app.I().
										Class("fas fa-fw fa-plus").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Br(),
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
							app.Div().
								Class("pf-c-slider__value pf-m-floating").
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
					app.Div().
						Class("pf-c-slider__actions").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Locked").
								Body(
									app.I().
										Class("fas fa-fw fa-lock").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Br(),
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
					app.Div().
						Class("pf-c-slider__actions").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Lock").
								Body(
									app.I().
										Class("fas fa-fw fa-lock-open").
										Aria("hidden", true),
								),
						),
				),
		)
}
