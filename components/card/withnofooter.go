package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithNoFooter struct {
	app.Compo
}

func (c *WithNoFooter) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-no-footer-example").
		Body(
			app.Div().
				Class("pf-c-card__title").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("This card has no footer"),
				),
		)
}
