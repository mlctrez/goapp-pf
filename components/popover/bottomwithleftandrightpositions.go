package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BottomWithLeftAndRightPositions struct {
	app.Compo
}

func (c *BottomWithLeftAndRightPositions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-popover pf-m-bottom-left").
				Aria("modal", true).
				Aria("labelledby", "popover-bottom-start-header").
				Aria("describedby", "popover-bottom-start-body").
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
								ID("popover-bottom-start-header").
								Body(
									app.Text("Popover header"),
								),
							app.Div().
								Class("pf-c-popover__body").
								ID("popover-bottom-start-body").
								Body(
									app.Text("This popover is to the bottom and at the start position"),
								),
							app.Footer().
								Class("pf-c-popover__footer").
								Body(
									app.Text("Popover footer"),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-popover pf-m-bottom-right").
				Aria("modal", true).
				Aria("labelledby", "popover-bottom-end-header").
				Aria("describedby", "popover-bottom-end-body").
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
								ID("popover-bottom-end-header").
								Body(
									app.Text("Popover header"),
								),
							app.Div().
								Class("pf-c-popover__body").
								ID("popover-bottom-end-body").
								Body(
									app.Text("This popover is to the bottom and at the end position"),
								),
							app.Footer().
								Class("pf-c-popover__footer").
								Body(
									app.Text("Popover footer"),
								),
						),
				),
		)
}
