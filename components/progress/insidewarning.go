package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InsideWarning struct {
	app.Compo
}

func (c *InsideWarning) Render() app.UI {
	return app.Div().
		Class("pf-c-progress pf-m-lg pf-m-inside pf-m-warning").
		ID("progress-inside-warning-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-inside-warning-example-description").
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
								Class("fas fa-fw fa-exclamation-triangle").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "100").
				Aria("valuenow", "100").
				Aria("labelledby", "progress-inside-warning-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "100%;").
						Body(
							app.Span().
								Class("pf-c-progress__measure").
								Body(
									app.Text("100%"),
								),
						),
				),
		)
}
