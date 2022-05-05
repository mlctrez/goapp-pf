package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-simple").
				Name("exampleRadioSimple"),
			app.Label().
				Class("pf-c-radio__label").
				For("radio-simple").
				Body(
					app.Text("Radio"),
				),
		)
}
