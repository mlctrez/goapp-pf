package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WidthAuto struct {
	app.Compo
}

func (c *WidthAuto) Render() app.UI {
	return app.Div().
		Class("pf-c-popover pf-m-right pf-m-width-auto").
		Aria("modal", true).
		Aria("labelledby", "popover-width-auto-header").
		Aria("describedby", "popover-width-auto-body").
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
						ID("popover-width-auto-header").
						Body(
							app.Text("Popover header"),
						),
					app.Div().
						Class("pf-c-popover__body").
						ID("popover-width-auto-body").
						Body(
							app.Text("Popovers body"),
						),
					app.Footer().
						Class("pf-c-popover__footer").
						Body(
							app.Text("Popover footer"),
						),
				),
		)
}
