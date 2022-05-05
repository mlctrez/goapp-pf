package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CustomTitle struct {
	app.Compo
}

func (c *CustomTitle) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("labelledby", "modal-custom-title").
		Aria("describedby", "modal-custom-description").
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
						Class("pf-c-title pf-m-4xl").
						ID("modal-custom-title").
						Body(
							app.Text("Custom title"),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("modal-custom-description").
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
