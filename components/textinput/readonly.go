package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ReadOnly struct {
	app.Compo
}

func (c *ReadOnly) Render() app.UI {
	return app.Input().
		Aria("label", "readonly input example").
		Class("pf-c-form-control").
		Type("text").
		Aria("invalid", "false").
		ReadOnly(true).
		DataSet("ouia-component-type", "PF4/TextInput").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
		Value("read only text input example")
}
