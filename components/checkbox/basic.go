package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-basic").
				Name("check-basic"),
			app.Label().
				Class("pf-c-check__label").
				For("check-basic").
				Body(
					app.Text("Check"),
				),
		)
}
