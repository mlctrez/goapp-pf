package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithMultipleBodySections struct {
	app.Compo
}

func (c *WithMultipleBodySections) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-multiple-body-example").
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
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
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
