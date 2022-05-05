package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FailureWithoutMeasure struct {
	app.Compo
}

func (c *FailureWithoutMeasure) Render() app.UI {
	return app.Div().
		Class("pf-c-progress pf-m-danger").
		ID("progress-no-measure-failure-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-no-measure-failure-example-description").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-progress__status").
				Aria("hidden", true).
				Body(
					app.Span().
						Class("pf-c-progress__status-icon").
						Body(
							app.I().
								Class("fas fa-fw fa-times-circle").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "100").
				Aria("valuenow", "33").
				Aria("labelledby", "progress-no-measure-failure-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;"),
				),
		)
}
