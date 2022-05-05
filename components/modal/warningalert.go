package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WarningAlert struct {
	app.Compo
}

func (c *WarningAlert) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box pf-m-warning").
		Aria("modal", true).
		Aria("labelledby", "warning-alert-title").
		Aria("describedby", "warning-alert-description").
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
				Class("pf-c-modal-box__header").
				Body(
					app.H1().
						Class("pf-c-modal-box__title pf-m-icon").
						ID("warning-alert-title").
						Body(
							app.Span().
								Class("pf-c-modal-box__title-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-exclamation-triangle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-u-screen-reader").
								Body(
									app.Text("Warning\n        alert:"),
								),
							app.Span().
								Class("pf-c-modal-box__title-text").
								Body(
									app.Text("Warning alert modal title"),
								),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("warning-alert-description").
				Body(
					app.Text("Modal description"),
				),
			app.Footer().
				Class("pf-c-modal-box__footer").
				Body(
					app.Text("Modal footer"),
				),
		)
}
