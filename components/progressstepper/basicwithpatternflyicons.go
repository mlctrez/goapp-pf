package progressstepper

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicWithPatternflyIcons struct {
	app.Compo
}

func (c *BasicWithPatternflyIcons) Render() app.UI {
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
									app.Text("Successful completion"),
								),
						),
				),
			app.Li().
				Class("pf-c-progress-stepper__step pf-m-current").
				Aria("label", "current step,in process step,").
				Body(
					app.Div().
						Class("pf-c-progress-stepper__step-connector").
						Body(
							app.Span().
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("pficon pf-icon-in-progress").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("In process"),
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
								Class("pf-c-progress-stepper__step-icon").
								Body(
									app.I().
										Class("pficon pf-icon-pending").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-progress-stepper__step-main").
						Body(
							app.Div().
								Class("pf-c-progress-stepper__step-title").
								Body(
									app.Text("Pending"),
								),
						),
				),
		)
}
