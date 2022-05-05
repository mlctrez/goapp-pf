package formcontrol

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Textarea struct {
	app.Compo
}

func (c *Textarea) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Textarea().
				Class("pf-c-form-control").
				Name("textarea-standard").
				ID("textarea-standard").
				Aria("label", "Standard textarea example").
				Body(
					app.Text("Standard"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control").
				ReadOnly(true).
				Name("textarea-readonly").
				ID("textarea-readonly").
				Aria("label", "Readonly textarea example").
				Body(
					app.Text("Readonly"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-success").
				Name("textarea-success").
				ID("textarea-success").
				Aria("label", "Success state textarea example").
				Body(
					app.Text("Success"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-warning").
				Name("textarea-warning").
				ID("textarea-warning").
				Aria("label", "Warning state textarea example").
				Body(
					app.Text("Warning"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control").
				Required(true).
				Name("textarea-error").
				ID("textarea-error").
				Aria("label", "Error state textarea example").
				Aria("invalid", true).
				Body(
					app.Text("Error"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-resize-vertical").
				Name("textarea-resize-vertical").
				ID("textarea-resize-vertical").
				Aria("label", "Resize vertical textarea example").
				Body(
					app.Text("Resizes vertically"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-resize-horizontal").
				Name("textarea-resize-horizontal").
				ID("textarea-resize-horizontal").
				Aria("label", "Resize horizontal textarea example").
				Body(
					app.Text("Resizes horizontally"),
				),
		)
}
