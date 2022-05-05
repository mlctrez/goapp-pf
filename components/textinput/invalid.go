package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Invalid struct {
	app.Compo
}

func (c *Invalid) Render() app.UI {
	return app.Input().
		Aria("label", "invalid text input example").
		Class("pf-c-form-control").
		Type("text").
		Aria("invalid", true).
		Required(true).
		DataSet("ouia-component-type", "PF4/TextInput").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
		Value(true)
}
