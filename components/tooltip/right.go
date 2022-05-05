package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Right struct {
	app.Compo
}

func (c *Right) Render() app.UI {
	return app.Div().
		Class("pf-c-tooltip pf-m-right").
		Aria("role", "tooltip").
		Body(
			app.Div().
				Class("pf-c-tooltip__arrow"),
			app.Div().
				Class("pf-c-tooltip__content").
				ID("tooltip-right-content").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
				),
		)
}
