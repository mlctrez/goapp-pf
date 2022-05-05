package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithDescription struct {
	app.Compo
}

func (c *WithDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-with-description").
				Name("check-with-description"),
			app.Label().
				Class("pf-c-check__label").
				For("check-with-description").
				Body(
					app.Text("Check with description"),
				),
			app.Span().
				Class("pf-c-check__description").
				Body(
					app.Text("Single-tenant cloud service hosted and managed by Red Hat that offers high-availability enterprise-grade clusters in a virtual private cloud on AWS od GCP."),
				),
		)
}
