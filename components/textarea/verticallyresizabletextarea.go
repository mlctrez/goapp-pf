package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticallyResizableTextArea struct {
	app.Compo
}

func (c *VerticallyResizableTextArea) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control pf-m-resize-vertical").
		Aria("invalid", "false").
		Aria("label", "text vertical resize example")
}
