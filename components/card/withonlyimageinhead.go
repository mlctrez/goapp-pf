package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithOnlyImageInHead struct {
	app.Compo
}

func (c *WithOnlyImageInHead) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-image-head-example").
		Body(
			app.Div().
				Class("pf-c-card__header").
				Body(
					app.Div().
						Class("pf-c-card__header-main").
						Body(
							app.Img().
								Src("/assets/images/pf_logo.svg").
								Width(300).
								Alt("Logo"),
						),
				),
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
