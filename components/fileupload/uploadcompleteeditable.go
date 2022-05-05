package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type UploadCompleteEditable struct {
	app.Compo
}

func (c *UploadCompleteEditable) Render() app.UI {
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
								ID("drop-file-text-input").
								Name("drop-file-text-input").
								Aria("label", "Read only filename").
								ReadOnly(true).
								Value("Sample.txt").
								Aria("describedby", "drop-file-browse"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("drop-file-browse").
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
						ID("drop-file-file-details").
						Name("drop-file-file-details").
						Aria("label", "Text area").
						Body(
							app.Text("Ssh-Rsa AAh3zJFkzjjakCJialksjfB3zJFkzzAAhhMskjjakCJialksjfB3z89z3zJFkz3 +kzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fc"),
						),
				),
		)
}
