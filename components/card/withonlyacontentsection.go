package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithOnlyAContentSection struct {
	app.Compo
}

func (c *WithOnlyAContentSection) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-body-example").
		Body(
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
				),
		)
}
