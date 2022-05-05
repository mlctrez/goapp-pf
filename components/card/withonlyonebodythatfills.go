package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithOnlyOneBodyThatFills struct {
	app.Compo
}

func (c *WithOnlyOneBodyThatFills) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-body-fill-example").
		Body(
			app.Div().
				Class("pf-c-card__title").
				Body(
					app.Text("Title"),
				),
			app.Div().
				Class("pf-c-card__body pf-m-no-fill").
				Body(
					app.Text("Body pf-m-no-fill"),
				),
			app.Div().
				Class("pf-c-card__body pf-m-no-fill").
				Body(
					app.Text("Body pf-m-no-fill"),
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
