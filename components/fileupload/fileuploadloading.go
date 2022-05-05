package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FileUploadLoading struct {
	app.Compo
}

func (c *FileUploadLoading) Render() app.UI {
	return app.Div().
		Class("pf-c-file-upload pf-m-loading").
		Body(
			app.Div().
				Class("pf-c-file-upload__file-select").
				Body(
					app.Div().
						Class("pf-c-input-group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								ID("file-upload-loading-text-input").
								Name("file-upload-loading").
								Aria("label", "Read only filename").
								ReadOnly(true).
								Value("Sample.png").
								Aria("describedby", "file-upload-loading-browse"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Disabled(true).
								ID("file-upload-loading-browse").
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
						ID("file-upload-loading-file-details").
						Name("file-upload-loading-file-details").
						Aria("label", "Text area").
						Body(
							app.Text("Ssh-Rsa AAh3zJFkzjjakCJialksjfB3zJFkzzAAhhMskjjakCJialksjfB3z89z3zJFkz3 +kzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fczJFkzMAjsauoox88aaZXphBx4fc"),
						),
					app.Div().
						Class("pf-c-file-upload__file-details-spinner").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-lg").
								Aria("role", "progressbar").
								Aria("label", "Loading...").
								Body(
									app.Span().
										Class("pf-c-spinner__clipper"),
									app.Span().
										Class("pf-c-spinner__lead-ball"),
									app.Span().
										Class("pf-c-spinner__tail-ball"),
								),
						),
				),
		)
}
