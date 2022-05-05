package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Input().
		Aria("label", "disabled text input example").
		Class("pf-c-form-control").
		Type("text").
		Aria("invalid", "false").
		Disabled(true).
		DataSet("ouia-component-type", "PF4/TextInput").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
		Value("disabled text input example")
}
