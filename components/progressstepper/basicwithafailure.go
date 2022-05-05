package progressstepper

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicWithAFailure struct {
	app.Compo
}

func (c *BasicWithAFailure) Render() app.UI {
	return app.Ol().
		Class("pf-c-progress-stepper").
		Body(
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-success").
				Aria("label", "completed step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("fas fa-check-circle").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("First step"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-success").
				Aria("label", "completed step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("fas fa-check-circle").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("Second step"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-success").
				Aria("label", "completed step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("fas fa-check-circle").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("Third step"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-current pf-m-danger").
				Aria("label", "current step,step with failure,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("fas fa-exclamation-circle").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("Fourth step"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-pending").
				Aria("label", "pending step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon"),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("Fifth step"),
								),
						),
				),
		)
}
