package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LeftWithTopAndBottomPositions struct {
	app.Compo
}

func (c *LeftWithTopAndBottomPositions) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-popover pf-m-left-top").
				Aria("modal", true).
				Aria("labelledby", "popover-left-start-header").
				Aria("describedby", "popover-left-start-body").
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
								ID("popover-left-start-header").
								Body(
									app.Text("Popover header"),
								),
							app.Div().
								Class("pf-c-popover__body").
								ID("popover-left-start-body").
								Body(
									app.Text("This popover is to the left and at the start position"),
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
				Class("pf-c-popover pf-m-left-bottom").
				Aria("modal", true).
				Aria("labelledby", "popover-left-end-header").
				Aria("describedby", "popover-left-end-body").
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
								ID("popover-left-end-header").
								Body(
									app.Text("Popover header"),
								),
							app.Div().
								Class("pf-c-popover__body").
								ID("popover-left-end-body").
								Body(
									app.Text("This popover is to the left and at the end position"),
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
