package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescription struct {
	app.Compo
}

func (c *WithDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("labelledby", "modal-with-description-title").
		Aria("describedby", "modal-with-description-description").
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
						Class("pf-c-modal-box__title").
						ID("modal-with-description-title").
						Body(
							app.Text("Modal title"),
						),
					app.Div().
						Class("pf-c-modal-box__description").
						ID("modal-with-description-description").
						Body(
							app.Text("A description is used when you want to provide more info about the modal than the title is able to describe. The content in the description is static and will not scroll with the rest of the modal body."),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
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
