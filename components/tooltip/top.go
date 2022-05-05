package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Top struct {
	app.Compo
}

func (c *Top) Render() app.UI {
	return app.Div().
		Class("pf-c-tooltip pf-m-top").
		Aria("role", "tooltip").
		Body(
			app.Div().
				Class("pf-c-tooltip__arrow"),
			app.Div().
				Class("pf-c-tooltip__content").
				ID("tooltip-top-content").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
				),
		)
}
