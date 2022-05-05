package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Bottom struct {
	app.Compo
}

func (c *Bottom) Render() app.UI {
	return app.Div().
		Class("pf-c-tooltip pf-m-bottom").
		Aria("role", "tooltip").
		Body(
			app.Div().
				Class("pf-c-tooltip__arrow"),
			app.Div().
				Class("pf-c-tooltip__content").
				ID("tooltip-bottom-content").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
				),
		)
}
