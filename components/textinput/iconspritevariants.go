package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type IconSpriteVariants struct {
	app.Compo
}

func (c *IconSpriteVariants) Render() app.UI {
	return app.Div().
		ID("ws-react-c-text-input-icon-sprite-variants").
		Class("ws-react-c-text-input pf-u-h-100").
		Body(
			app.Input().
				Aria("label", "success icon sprite text input example").
				Class("pf-c-form-control pf-m-icon-sprite pf-m-success").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "warning icon sprite text input example").
				Class("pf-c-form-control pf-m-icon-sprite pf-m-warning").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-16").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "error icon sprite text input example").
				Class("pf-c-form-control pf-m-icon-sprite").
				Type("text").
				Aria("invalid", true).
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-17").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "calendar icon sprite text input example").
				Class("pf-c-form-control pf-m-icon-sprite pf-m-icon pf-m-calendar").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-18").
				Value(true),
			app.Br(),
			app.Br(),
			app.Input().
				Aria("label", "clock icon sprite text input example").
				Class("pf-c-form-control pf-m-icon-sprite pf-m-icon pf-m-clock").
				Type("text").
				Aria("invalid", "false").
				DataSet("ouia-component-type", "PF4/TextInput").
				DataSet("ouia-safe", true).
				DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-19").
				Value(true),
		)
}
