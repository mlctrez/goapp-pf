package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Inside struct {
	app.Compo
}

func (c *Inside) Render() app.UI {
	return app.Div().
		Class("pf-c-progress pf-m-lg pf-m-inside").
		ID("progress-inside-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-inside-example-description").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-progress__status").
				Aria("hidden", true),
			app.Div().
				Class("pf-c-progress__bar").
				Aria("role", "progressbar").
				Aria("valuemin", "0").
				Aria("valuemax", "100").
				Aria("valuenow", "33").
				Aria("labelledby", "progress-inside-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;").
						Body(
							app.Span().
								Class("pf-c-progress__measure").
								Body(
									app.Text("33%"),
								),
						),
				),
		)
}
