package radio

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-radio").
				Body(
					app.Input().
						Class("pf-c-radio__input").
						Type("radio").
						ID("radio-disabled").
						Name("exampleRadioDisabled").
						Disabled(true),
					app.Label().
						Class("pf-c-radio__label pf-m-disabled").
						For("radio-disabled").
						Body(
							app.Text("Radio disabled"),
						),
				),
			app.Div().
				Class("pf-c-radio").
				Body(
					app.Input().
						Class("pf-c-radio__input").
						Type("radio").
						ID("radio-disabled-checked").
						Name("exampleRadioDisabledChecked").
						Disabled(true).
						Checked(true),
					app.Label().
						Class("pf-c-radio__label pf-m-disabled").
						For("radio-disabled-checked").
						Body(
							app.Text("Radio disabled checked"),
						),
				),
		)
}
