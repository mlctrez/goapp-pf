package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type NonSelectable struct {
	app.Compo
}

func (c *NonSelectable) Render() app.UI {
	return app.Div().
		Class("pf-c-card pf-m-non-selectable-raised").
		ID("card-non-selectable-example").
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
