package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type UploadCompleteNonEditable struct {
	app.Compo
}

func (c *UploadCompleteNonEditable) Render() app.UI {
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
								ID("browsed-file-upload-complete-text-input").
								Name("browsed-file-upload-complete-text-input").
								Aria("label", "Read only filename").
								ReadOnly(true).
								Value("Read only filename").
								Aria("describedby", "browsed-file-upload-complete-browse"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("browsed-file-upload-complete-browse").
								Body(
									app.Text("Upload"),
								),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
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
						ID("browsed-file-upload-complete-file-details").
						Name("browsed-file-upload-complete-file-details").
						Aria("label", "Text area").
						ReadOnly(true).
						Body(
							app.Text("Ssh-Rsa AAh3zJFkzjjakCJialksjfB3zJFkzzAAhhMskjjakCJialksjfB3z89z3zJFkz3 +kzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fc"),
						),
				),
		)
}
