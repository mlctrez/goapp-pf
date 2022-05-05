package textinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TruncatedOnLeft struct {
	app.Compo
}

func (c *TruncatedOnLeft) Render() app.UI {
	return app.Input().
		Aria("label", "left-truncated text input example").
		Class("pf-c-form-control").
		Type("text").
		Aria("invalid", "false").
		DataSet("ouia-component-type", "PF4/TextInput").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-15").
		Value("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.")
}
