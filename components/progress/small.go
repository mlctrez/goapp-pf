package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Small struct {
	app.Compo
}

func (c *Small) Render() app.UI {
	return app.Div().
		Class("pf-c-progress pf-m-sm").
		ID("progress-sm-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-sm-example-description").
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
				Aria("labelledby", "progress-sm-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;"),
				),
		)
}
