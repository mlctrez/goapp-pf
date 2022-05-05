package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithNoTitle struct {
	app.Compo
}

func (c *WithNoTitle) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-no-title-example").
		Body(
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("This card has no title"),
				),
			app.Div().
				Class("pf-c-card__footer").
				Body(
					app.Text("Footer"),
				),
		)
}
