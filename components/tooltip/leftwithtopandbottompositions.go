package tooltip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LeftWithTopAndBottomPositions struct {
	app.Compo
}

func (c *LeftWithTopAndBottomPositions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tooltip pf-m-left-top").
				Aria("role", "tooltip").
				Body(
					app.Div().
						Class("pf-c-tooltip__arrow"),
					app.Div().
						Class("pf-c-tooltip__content").
						ID("tooltip-left-top-content").
						Body(
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-tooltip pf-m-left-bottom").
				Aria("role", "tooltip").
				Body(
					app.Div().
						Class("pf-c-tooltip__arrow"),
					app.Div().
						Class("pf-c-tooltip__content").
						ID("tooltip-left-bottom-content").
						Body(
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam id feugiat augue, nec fringilla turpis."),
						),
				),
		)
}
