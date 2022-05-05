package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutTitle struct {
	app.Compo
}

func (c *WithoutTitle) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("label", "Example of a modal without a title").
		Aria("describedby", "modal-no-title-description").
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
			app.Div().
				Class("pf-c-modal-box__body").
				Body(
					app.Span().
						ID("modal-no-title-description").
						Body(
							app.Text("When static text describing the modal is available, it can be wrapped with an ID referring to the modal's aria-describedby value. Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."),
						),
					app.Text("Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
				),
			app.Footer().
				Class("pf-c-modal-box__footer").
				Body(
					app.Text("Modal footer"),
				),
		)
}
