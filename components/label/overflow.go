package label

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Overflow struct {
	app.Compo
}

func (c *Overflow) Render() app.UI {
	return app.Button().
		Class("pf-c-label pf-m-overflow").
		Body(
			app.Span().
				Class("pf-c-label__content").
				Body(
					app.Text("Overflow"),
				),
		)
}
