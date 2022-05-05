package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Checked struct {
	app.Compo
}

func (c *Checked) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-checked").
				Name("exampleRadioChecked").
				Checked(true),
			app.Label().
				Class("pf-c-radio__label").
				For("radio-checked").
				Body(
					app.Text("Radio checked"),
				),
		)
}
