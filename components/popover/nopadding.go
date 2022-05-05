package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NoPadding struct {
	app.Compo
}

func (c *NoPadding) Render() app.UI {
	return app.Div().
		Class("pf-c-popover pf-m-right pf-m-no-padding").
		Aria("modal", true).
		Aria("label", "Popover with no padding example").
		Aria("describedby", "popover-no-padding-body").
		Body(
			app.Div().
				Class("pf-c-popover__arrow"),
			app.Div().
				Class("pf-c-popover__content").
				Body(
					app.Div().
						Class("pf-c-popover__body").
						ID("popover-no-padding-body").
						Body(
							app.Text("This popover has no padding and is intended for use with content that has its own spacing and should touch the edges of the popover container."),
						),
				),
		)
}
