package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FiniteStep struct {
	app.Compo
}

func (c *FiniteStep) Render() app.UI {
	return app.Div().
		Class("pf-c-progress").
		ID("progress-finite-step-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-finite-step-example-description").
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
							app.Text("2 of 5 units"),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "5").
				Aria("valuenow", "2").
				Aria("valuetext", "2 of 5 units").
				Aria("labelledby", "progress-finite-step-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "40%;"),
				),
		)
}
