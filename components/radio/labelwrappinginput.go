package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LabelWrappingInput struct {
	app.Compo
}

func (c *LabelWrappingInput) Render() app.UI {
	return app.Label().
		Class("pf-c-radio").
		For("radio-wrap").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-wrap").
				Name("exampleRadioWrap"),
			app.Span().
				Class("pf-c-radio__label").
				Body(
					app.Text("Radio label wraps input"),
				),
		)
}
