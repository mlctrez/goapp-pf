package progressstepper

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicWithAnIssue struct {
	app.Compo
}

func (c *BasicWithAnIssue) Render() app.UI {
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
				Class("pf-c-progress-stepper__step pf-m-warning").
				Aria("label", "step with issue,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("fas fa-exclamation-triangle").
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
