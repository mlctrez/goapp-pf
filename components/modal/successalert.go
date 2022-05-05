package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SuccessAlert struct {
	app.Compo
}

func (c *SuccessAlert) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box pf-m-success").
		Aria("modal", true).
		Aria("labelledby", "success-alert-title").
		Aria("describedby", "success-alert-description").
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
						ID("success-alert-title").
						Body(
							app.Span().
								Class("pf-c-modal-box__title-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-check-circle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-u-screen-reader").
								Body(
									app.Text("Success\n        alert:"),
								),
							app.Span().
								Class("pf-c-modal-box__title-text").
								Body(
									app.Text("Success alert modal title"),
								),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("success-alert-description").
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
