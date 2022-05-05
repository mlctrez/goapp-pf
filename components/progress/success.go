package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Success struct {
	app.Compo
}

func (c *Success) Render() app.UI {
	return app.Div().
		Class("pf-c-progress pf-m-success").
		ID("progress-success-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-success-example-description").
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
							app.Text("100%"),
						),
					app.Span().
						Class("pf-c-progress__status-icon").
						Body(
							app.I().
								Class("fas fa-fw fa-check-circle").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "100").
				Aria("valuenow", "100").
				Aria("labelledby", "progress-success-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "100%;"),
				),
		)
}
