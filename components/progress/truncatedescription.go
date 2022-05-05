package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TruncateDescription struct {
	app.Compo
}

func (c *TruncateDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-progress").
		ID("progress-truncate-description-example").
		Body(
			app.Div().
				Class("pf-c-progress__description pf-m-truncate").
				ID("progress-truncate-description-example-description").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean quis ultricies lectus, eu lobortis mauris. Morbi pretium arcu id rhoncus mollis. Donec accumsan tincidunt enim nec varius. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Suspendisse potenti."),
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
				Aria("labelledby", "progress-truncate-description-example-description").
				Body(
					app.Div().
						Class("pf-c-progress__indicator").
						Style("width", "33%;"),
				),
		)
}
