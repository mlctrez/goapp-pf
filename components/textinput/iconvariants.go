package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type IconVariants struct {
	app.Compo
}

func (c *IconVariants) Render() app.UI {
	return app.Div().
		ID("ws-react-c-text-input-icon-variants").
		Class("ws-react-c-text-input pf-u-h-100").
		Body(
			app.Input().
				Aria("label", "text input example").
				Class("pf-c-form-control pf-m-icon pf-m-calendar").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "text input example").
				Class("pf-c-form-control pf-m-icon pf-m-clock").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-16").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "text input example").
				Class("pf-c-form-control pf-m-icon").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-17").
				Value(true).
				Style("", ""),
		)
}
