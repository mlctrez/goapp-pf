package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutMeasure struct {
	app.Compo
}

func (c *WithoutMeasure) Render() app.UI {
	return app.Div().
		Class("pf-c-progress").
		ID("progress-no-measure-example").
		Body(
			app.Div().
				Class("pf-c-progress__description").
				ID("progress-no-measure-example-description").
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
				Aria("labelledby", "progress-no-measure-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;"),
				),
		)
}
