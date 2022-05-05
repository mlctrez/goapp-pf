package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Plain struct {
	app.Compo
}

func (c *Plain) Render() app.UI {
	return app.Div().
		Class("pf-c-card pf-m-plain").
		ID("card-plain-example").
		Body(
			app.Div().
				Class("pf-c-card__title").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
				),
			app.Div().
				Class("pf-c-card__footer").
				Body(
					app.Text("Footer"),
				),
		)
}
