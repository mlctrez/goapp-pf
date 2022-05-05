package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Right struct {
	app.Compo
}

func (c *Right) Render() app.UI {
	return app.Div().
		Class("pf-c-popover pf-m-right").
		Aria("modal", true).
		Aria("labelledby", "popover-right-header").
		Aria("describedby", "popover-right-body").
		Body(
			app.Div().
				Class("pf-c-popover__arrow"),
			app.Div().
				Class("pf-c-popover__content").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Close").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
					app.H1().
						Class("pf-c-title pf-m-md").
						ID("popover-right-header").
						Body(
							app.Text("Popover header"),
						),
					app.Div().
						Class("pf-c-popover__body").
						ID("popover-right-body").
						Body(
							app.Text("Popovers are triggered by click rather than hover. Click again to close."),
						),
					app.Footer().
						Class("pf-c-popover__footer").
						Body(
							app.Text("Popover footer"),
						),
				),
		)
}
