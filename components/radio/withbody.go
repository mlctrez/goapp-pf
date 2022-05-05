package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithBody struct {
	app.Compo
}

func (c *WithBody) Render() app.UI {
	return app.Div().
		Class("pf-c-radio").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-body").
				Name("exampleRadioBody"),
			app.Label().
				Class("pf-c-radio__label").
				For("radio-body").
				Body(
					app.Text("Radio with body"),
				),
			app.Span().
				Class("pf-c-radio__body").
				Body(
					app.Text("This is where custom content goes."),
				),
		)
}
