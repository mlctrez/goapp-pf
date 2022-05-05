package fileupload

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FileUploadInFormWithError struct {
	app.Compo
}

func (c *FileUploadInFormWithError) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form").
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
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
												ID("file-upload-error-text-input").
												Name("file-upload-error-text-input").
												Aria("label", "File upload error").
												Required(true).
												Value("Sample.png").
												Aria("describedby", "file-upload-error-browse"),
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												ID("file-upload-error-browse").
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
								Aria("describedby", "textAreaHelperText1").
								Aria("invalid", true).
								Body(
									app.Textarea().
										Class("pf-c-form-control pf-m-resize-vertical").
										ID("file-upload-error-file-details").
										Name("file-upload-error-file-details").
										Aria("label", "Empty text area").
										Aria("describedby", "textAreaHelperText1").
										Aria("invalid", true),
								),
						),
					app.P().
						Class("pf-c-form__helper-text pf-m-error").
						ID("textAreaHelperText1").
						Aria("live", "polite").
						Body(
							app.Text("We don't support this file type. Try again with a different file type."),
						),
				),
		)
}
