package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control").
		Aria("invalid", "false").
		Aria("label", "text area example")
}
