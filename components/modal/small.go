package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Small struct {
	app.Compo
}

func (c *Small) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box pf-m-sm").
		Aria("modal", true).
		Aria("labelledby", "modal-sm-title").
		Aria("describedby", "modal-sm-description").
		Body(
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Close dialog").
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
						ID("modal-sm-title").
						Body(
							app.Text("Modal title"),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("modal-sm-description").
				Body(
					app.Text("Static text describing modal purpose. Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod\n    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,\n    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo\n    consequat."),
				),
			app.Footer().
				Class("pf-c-modal-box__footer").
				Body(
					app.Text("Modal footer"),
				),
		)
}
