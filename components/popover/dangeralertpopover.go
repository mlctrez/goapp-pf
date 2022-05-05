package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DangerAlertPopover struct {
	app.Compo
}

func (c *DangerAlertPopover) Render() app.UI {
	return app.Div().
		Class("pf-c-popover pf-m-danger pf-m-top").
		Aria("modal", true).
		Aria("labelledby", "popover-top-header").
		Aria("describedby", "popover-top-body").
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
					app.Header().
						Class("pf-c-popover__header").
						Body(
							app.H1().
								Class("pf-c-popover__title pf-m-icon").
								ID("popover-top-header").
								Body(
									app.Span().
										Class("pf-c-popover__title-icon").
										Body(
											app.I().
												Class("fas fa-fw fa-exclamation-circle").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-u-screen-reader").
										Body(
											app.Text("Danger\n          alert:"),
										),
									app.Span().
										Class("pf-c-popover__title-text").
										Body(
											app.Text("Danger popover title"),
										),
								),
						),
					app.Div().
						Class("pf-c-popover__body").
						ID("popover-top-body").
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
