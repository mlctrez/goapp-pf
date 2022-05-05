package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ModalBoxAsGenericContainer struct {
	app.Compo
}

func (c *ModalBoxAsGenericContainer) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("labelledby", "modal-generic-container-description").
		Body(
			app.P().
				ID("modal-generic-container-description").
				Body(
					app.Text("The modal box children elements can be removed, and the modal serves as a generic modal container. One use case of this is when creating a wizard in a modal."),
				),
		)
}
