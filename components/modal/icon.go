package modal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Icon struct {
	app.Compo
}

func (c *Icon) Render() app.UI {
	return app.Div().
		Class("pf-c-modal-box").
		Aria("modal", true).
		Aria("labelledby", "icon-title").
		Aria("describedby", "icon-description").
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
						ID("icon-title").
						Body(
							app.Span().
								Class("pf-c-modal-box__title-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-bullhorn").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-modal-box__title-text").
								Body(
									app.Text("Modal with icon title"),
								),
						),
				),
			app.Div().
				Class("pf-c-modal-box__body").
				ID("icon-description").
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
