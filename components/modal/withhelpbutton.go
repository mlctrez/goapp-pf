package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithHelpButton struct {
	app.Compo
}

func (c *WithHelpButton) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("labelledby", "modal-help-title").
		Aria("describedby", "modal-help-description").
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
				Class("pf-c-modal-box__header pf-m-help").
				Body(
					app.Div().
						Class("pf-c-modal-box__header-main").
						Body(
							app.H1().
								Class("pf-c-modal-box__title").
								ID("modal-help-title").
								Body(
									app.Text("Modal title Modal title Modal title Modal title Modal title Modal title Modal title Modal title"),
								),
							app.Div().
								Class("pf-c-modal-box__description").
								ID("modal-help-description").
								Body(
									app.Text("A description is used when you want to provide more info about the modal than the title is able to describe. The content in the description is static and will not scroll with the rest of the modal body."),
								),
						),
					app.Div().
						Class("pf-c-modal-box__header-help").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Help").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
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
