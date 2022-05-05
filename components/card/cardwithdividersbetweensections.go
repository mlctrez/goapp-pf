package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CardWithDividersBetweenSections struct {
	app.Compo
}

func (c *CardWithDividersBetweenSections) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		Body(
			app.Div().
				Class("pf-c-card__title").
				Body(
					app.Text("Title"),
				),
			app.Hr().
				Class("pf-c-divider"),
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
				),
			app.Hr().
				Class("pf-c-divider"),
			app.Div().
				Class("pf-c-card__body").
				Body(
					app.Text("Body"),
				),
			app.Hr().
				Class("pf-c-divider"),
			app.Div().
				Class("pf-c-card__footer").
				Body(
					app.Text("Footer"),
				),
		)
}
