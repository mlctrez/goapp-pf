package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DragFileHoverComponent struct {
	app.Compo
}

func (c *DragFileHoverComponent) Render() app.UI {
	return app.Div().
		Class("pf-c-file-upload pf-m-drag-hover").
		Body(
			app.Div().
				Class("pf-c-file-upload__file-select").
				Body(
					app.Div().
						Class("pf-c-input-group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								ID("drag-file-hover-component-text-input").
								Name("drag-file-hover-component-text-input").
								Aria("label", "Drag and drop a file or upload one").
								ReadOnly(true).
								Placeholder("false").
								Aria("describedby", "drag-file-hover-component-browse"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("drag-file-hover-component-browse").
								Body(
									app.Text("Upload"),
								),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Disabled(true).
								Body(
									app.Text("Clear"),
								),
						),
				),
			app.Div().
				Class("pf-c-file-upload__file-details").
				Body(
					app.Textarea().
						Class("pf-c-form-control pf-m-resize-vertical").
						ID("drag-file-hover-component-file-details").
						Name("drag-file-hover-component-file-details").
						Aria("label", "Empty text area"),
				),
		)
}
