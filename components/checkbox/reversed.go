package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Reversed struct {
	app.Compo
}

func (c *Reversed) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Label().
				Class("pf-c-check__label").
				For("check-reversed").
				Body(
					app.Text("Check reversed"),
				),
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-reversed").
				Name("check-reversed"),
		)
}
