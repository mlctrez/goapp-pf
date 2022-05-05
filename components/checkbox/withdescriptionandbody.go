package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescriptionAndBody struct {
	app.Compo
}

func (c *WithDescriptionAndBody) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-with-description-body").
				Name("check-with-description-body"),
			app.Label().
				Class("pf-c-check__label").
				For("check-with-description-body").
				Body(
					app.Text("Check with description and body"),
				),
			app.Span().
				Class("pf-c-check__description").
				Body(
					app.Text("Single-tenant cloud service hosted and managed by Red Hat that offers high-availability enterprise-grade clusters in a virtual private cloud on AWS od GCP."),
				),
			app.Span().
				Class("pf-c-check__body").
				Body(
					app.Text("This is where custom content goes."),
				),
		)
}
