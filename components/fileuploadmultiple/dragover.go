package fileuploadmultiple

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DragOver struct {
	app.Compo
}

func (c *DragOver) Render() app.UI {
	return app.Div().
		Class("pf-c-multiple-file-upload pf-m-drag-over").
		Body(
			app.Div().
				Class("pf-c-multiple-file-upload__main").
				Body(
					app.Div().
						Class("pf-c-multiple-file-upload__title").
						Body(
							app.Div().
								Class("pf-c-multiple-file-upload__title-icon").
								Body(
									app.I().
										Class("fas fa-upload").
										Aria("hidden", true),
								),
							app.Div().
								Class("pf-c-multiple-file-upload__title-text").
								Body(
									app.Text("Drag and drop files here"), app.Div().
										Class("pf-c-multiple-file-upload__title-text-separator").
										Body(
											app.Text("or"),
										),
								),
						),
					app.Div().
						Class("pf-c-multiple-file-upload__upload").
						Body(
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Upload"),
								),
						),
					app.Div().
						Class("pf-c-multiple-file-upload__info").
						Body(
							app.Text("Accepted file types: JPEG, Doc, PDF, PNG"),
						),
				),
		)
}
