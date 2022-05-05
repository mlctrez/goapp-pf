package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LeftAlignedText struct {
	app.Compo
}

func (c *LeftAlignedText) Render() app.UI {
	return app.Div().
		Class("pf-c-tooltip pf-m-top").
		Aria("role", "tooltip").
		Body(
			app.Div().
				Class("pf-c-tooltip__arrow"),
			app.Div().
				Class("pf-c-tooltip__content pf-m-text-align-left").
				ID("tooltip-text-align-left-example").
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
				),
		)
}
