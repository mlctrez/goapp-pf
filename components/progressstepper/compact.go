package progressstepper

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Ol().
		Class("pf-c-progress-stepper pf-m-compact").
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
							app.Div().
								Class("pf-c-progress-stepper__step-description").
								Body(
									app.Text("This is the first thing to happen"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-current pf-m-info").
				Aria("label", "current step,in process step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("pficon pf-icon-resources-full").
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
							app.Div().
								Class("pf-c-progress-stepper__step-description").
								Body(
									app.Text("This is the second thing to happen"),
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
									app.Text("Third step"),
								),
							app.Div().
								Class("pf-c-progress-stepper__step-description").
								Body(
									app.Text("This is the last thing to happen"),
								),
						),
				),
		)
}
