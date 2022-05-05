package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AutoResizing struct {
	app.Compo
}

func (c *AutoResizing) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control").
		Aria("invalid", "false").
		Aria("label", "auto resizing text area example")
}
