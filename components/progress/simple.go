package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Simple struct {
	app.Compo
}

func (c *Simple) Render() app.UI {
	return app.Div().
		Class("pf-c-progress").
		ID("progress-simple-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-simple-example-description").
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
							app.Text("33%"),
						),
				),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "100").
				Aria("valuenow", "33").
				Aria("labelledby", "progress-simple-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;"),
				),
		)
}
