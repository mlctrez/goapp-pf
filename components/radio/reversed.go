package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Reversed struct {
	app.Compo
}

func (c *Reversed) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Label().
				Class("pf-c-radio__label").
				For("radio-rev").
				Body(
					app.Text("Radio reversed"),
				),
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-rev").
				Name("exampleRadioReversed"),
		)
}
