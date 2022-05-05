package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontallyResizableTextArea struct {
	app.Compo
}

func (c *HorizontallyResizableTextArea) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control pf-m-resize-horizontal").
		Aria("invalid", "false").
		Aria("label", "text horizontal resize example")
}
