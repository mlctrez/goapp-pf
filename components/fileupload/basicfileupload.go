package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicFileUpload struct {
	app.Compo
}

func (c *BasicFileUpload) Render() app.UI {
	return app.Div().
		Class("pf-c-file-upload").
		Body(
			app.Div().
				Class("pf-c-file-upload__file-select").
				Body(
					app.Div().
						Class("pf-c-input-group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								ID("basic-file-upload-text-input").
								Name("basic-file-upload-text-input").
								Aria("label", "Drag and drop a file or upload one").
								ReadOnly(true).
								Placeholder("false").
								Aria("describedby", "basic-file-upload-browse"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("basic-file-upload-browse").
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
						ID("basic-file-upload-file-details").
						Name("basic-file-upload-file-details").
						Aria("label", "Empty text area"),
				),
		)
}
