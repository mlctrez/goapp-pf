package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescription struct {
	app.Compo
}

func (c *WithDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-description").
				Name("exampleRadioDescription"),
			app.Label().
				Class("pf-c-radio__label").
				For("radio-description").
				Body(
					app.Text("Radio with description"),
				),
			app.Span().
				Class("pf-c-radio__description").
				Body(
					app.Text("Single-tenant cloud service hosted and managed by Red Hat that offers high-availability enterprise-grade clusters in a virtual private cloud on AWS od GCP."),
				),
		)
}
