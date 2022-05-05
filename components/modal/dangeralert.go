package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DangerAlert struct {
	app.Compo
}

func (c *DangerAlert) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box pf-m-danger").
		Aria("modal", true).
		Aria("labelledby", "danger-alert-title").
		Aria("describedby", "danger-alert-description").
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
						ID("danger-alert-title").
						Body(
							app.Span().
								Class("pf-c-modal-box__title-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-exclamation-circle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-u-screen-reader").
								Body(
									app.Text("Danger\n        alert:"),
								),
							app.Span().
								Class("pf-c-modal-box__title-text").
								Body(
									app.Text("Danger alert modal title"),
								),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("danger-alert-description").
				Body(
					app.Text("To support screen reader user awareness of the dialog text, the dialog text is wrapped in a div that is referenced by aria-describedby."),
				),
			app.Footer().
				Class("pf-c-modal-box__footer").
				Body(
					app.Text("Modal footer"),
				),
		)
}
