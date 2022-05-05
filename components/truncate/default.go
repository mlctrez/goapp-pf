package truncate

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Default struct {
	app.Compo
}

func (c *Default) Render() app.UI {
	return app.Div().
		Class("pf-c-truncate--example").
		Body(
			app.Span().
				Class("pf-c-truncate").
				Body(
					app.Span().
						Class("pf-c-truncate__start").
						Body(
							app.Text("Vestibulum interdum risus et enim faucibus, sit amet molestie est accumsan."),
						),
				),
		)
}
