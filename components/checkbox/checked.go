package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Checked struct {
	app.Compo
}

func (c *Checked) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-checked").
				Name("check-checked").
				Checked(true),
			app.Label().
				Class("pf-c-check__label").
				For("check-checked").
				Body(
					app.Text("Check checked"),
				),
		)
}
