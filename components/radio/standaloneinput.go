package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StandaloneInput struct {
	app.Compo
}

func (c *StandaloneInput) Render() app.UI {
	return app.Div().
		Class("pf-c-radio pf-m-standalone").
		Body(
			app.Input().
				Class("pf-c-radio__input").
				Type("radio").
				ID("radio-standalone").
				Name("exampleRadioStandalone").
				Aria("label", "Standalone input"),
		)
}
