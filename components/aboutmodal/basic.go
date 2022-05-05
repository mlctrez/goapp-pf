package aboutmodal

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-about-modal-box").
		Aria("role", "dialog").
		Aria("modal", true).
		Aria("labelledby", "about-modal-title").
		Body(
			app.Div().
				Class("pf-c-about-modal-box__brand").
				Body(
					app.Img().
						Class("pf-c-about-modal-box__brand-image").
						Src("/assets/images/pf_mini_logo_white.svg").
						Alt("PatternFly brand logo"),
				),
			app.Div().
				Class("pf-c-about-modal-box__close").
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
				),
			app.Div().
				Class("pf-c-about-modal-box__header").
				Body(
					app.H1().
						Class("pf-c-title pf-m-4xl").
						ID("about-modal-title").
						Body(
							app.Text("Product name"),
						),
				),
			app.Div().
				Class("pf-c-about-modal-box__hero"),
			app.Div().
				Class("pf-c-about-modal-box__content").
				Body(
					app.Div().
						Class("pf-c-about-modal-box__body").
						Body(
							app.Text("content"),
						),
					app.P().
						Class("pf-c-about-modal-box__strapline").
						Body(
							app.Text("Trademark and copyright information here"),
						),
				),
		)
}
