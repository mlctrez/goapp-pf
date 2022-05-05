package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectTextUsingRef struct {
	app.Compo
}

func (c *SelectTextUsingRef) Render() app.UI {
	return app.Input().
		Aria("label", "select-all").
		Class("pf-c-form-control").
		Type("text").
		Aria("invalid", "false").
		DataSet("ouia-component-type", "PF4/TextInput").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
		Value("select all on click")
}
