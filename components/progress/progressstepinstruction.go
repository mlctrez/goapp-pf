package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ProgressStepInstruction struct {
	app.Compo
}

func (c *ProgressStepInstruction) Render() app.UI {
	return app.Div().
		Class("pf-c-progress").
		ID("progress-step-instruction-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-step-instruction-example-description").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-progress__status").
				Aria("hidden", true).
				Body(
					app.Span().
						Class("pf-c-progress__measure").
						Body(
							app.Text("Step 2: Copying files"),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "5").
				Aria("valuenow", "2").
				Aria("valuetext", "Step 2: Copying files").
				Aria("labelledby", "progress-step-instruction-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "40%;"),
				),
		)
}
