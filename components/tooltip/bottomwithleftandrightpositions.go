package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BottomWithLeftAndRightPositions struct {
	app.Compo
}

func (c *BottomWithLeftAndRightPositions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tooltip pf-m-bottom-left").
				Aria("role", "tooltip").
				Body(
					app.Div().
						Class("pf-c-tooltip__arrow"),
					app.Div().
						Class("pf-c-tooltip__content").
						ID("tooltip-bottom-left-content").
						Body(
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-tooltip pf-m-bottom-right").
				Aria("role", "tooltip").
				Body(
					app.Div().
						Class("pf-c-tooltip__arrow"),
					app.Div().
						Class("pf-c-tooltip__content").
						ID("tooltip-bottom-right-content").
						Body(
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
						),
				),
		)
}
