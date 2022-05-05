package truncate

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Start struct {
	app.Compo
}

func (c *Start) Render() app.UI {
	return app.Div().
		Class("pf-c-truncate--example").
		Body(
			app.Span().
				Class("pf-c-truncate").
				Body(
					app.Span().
						Class("pf-c-truncate__end").
						Body(
							app.Text("Vestibulum interdum risus et enim faucibus, sit amet molestie est accumsan.\u200e"),
						),
				),
		)
}
