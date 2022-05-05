package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Left struct {
	app.Compo
}

func (c *Left) Render() app.UI {
	return app.Div().
		Class("pf-c-tooltip pf-m-left").
		Aria("role", "tooltip").
		Body(
			app.Div().
				Class("pf-c-tooltip__arrow"),
			app.Div().
				Class("pf-c-tooltip__content").
				ID("tooltip-left-content").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
				),
		)
}
