package truncate

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Middle struct {
	app.Compo
}

func (c *Middle) Render() app.UI {
	return app.Div().
		Class("pf-c-truncate--example").
		Body(
			app.Span().
				Class("pf-c-truncate").
				Body(
					app.Span().
						Class("pf-c-truncate__start").
						Body(
							app.Text("redhat_logo_black_and_white_reversed_simple_with_fedora_con"),
						),
					app.Span().
						Class("pf-c-truncate__end").
						Body(
							app.Text("tainer.zip"),
						),
				),
		)
}
