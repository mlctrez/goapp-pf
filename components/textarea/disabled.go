package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control").
		Aria("invalid", "false").
		Disabled(true).
		Aria("label", "disabled text area example")
}
