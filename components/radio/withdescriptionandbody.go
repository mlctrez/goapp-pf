package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescriptionAndBody struct {
	app.Compo
}

func (c *WithDescriptionAndBody) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-description-body").
				Name("exampleRadioDescriptionBody"),
			app.Label().
				Class("pf-c-radio__label").
				For("radio-description-body").
				Body(
					app.Text("Radio with description and body"),
				),
			app.Span().
				Class("pf-c-radio__description").
				Body(
					app.Text("Single-tenant cloud service hosted and managed by Red Hat that offers high-availability enterprise-grade clusters in a virtual private cloud on AWS od GCP."),
				),
			app.Span().
				Class("pf-c-radio__body").
				Body(
					app.Text("This is where custom content goes."),
				),
		)
}
