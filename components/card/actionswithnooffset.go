package card

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ActionsWithNoOffset struct {
	app.Compo
}

func (c *ActionsWithNoOffset) Render() app.UI {
	return app.Div().
		Class("pf-c-card").
		ID("card-action-no-offset").
		Body(
			app.Div().
				Class("pf-c-card__header").
				Body(
					app.Div().
						Class("pf-c-card__actions pf-m-no-offset").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("Action"),
								),
						),
					app.H1().
						Class("pf-c-title pf-m-2xl").
						ID("card-action-no-offset-check-label").
						Body(
							app.Text("This is a card title"),
						),
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
