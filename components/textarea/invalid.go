package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Invalid struct {
	app.Compo
}

func (c *Invalid) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control").
		Aria("invalid", true).
		Required(true).
		Aria("label", "invalid text area example")
}
