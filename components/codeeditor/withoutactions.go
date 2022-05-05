package codeeditor

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutActions struct {
	app.Compo
}

func (c *WithoutActions) Render() app.UI {
	return app.Div().
		Class("pf-c-code-editor").
		Body(
			app.Div().
				Class("pf-c-code-editor__header").
				Body(
					app.Div().
						Class("pf-c-code-editor__header-main"),
					app.Div().
						Class("pf-c-code-editor__tab").
						Body(
							app.Span().
								Class("pf-c-code-editor__tab-icon").
								Body(
									app.I().
										Class("fas fa-code"),
								),
							app.Span().
								Class("pf-c-code-editor__tab-text").
								Body(
									app.Text("YAML"),
								),
						),
				),
			app.Div().
				Class("pf-c-code-editor__main").
				Body(
					app.Div().
						Class("pf-c-empty-state pf-m-lg").
						Body(
							app.Div().
								Class("pf-c-empty-state__content").
								Body(
									app.Div().
										Class("pf-c-empty-state__icon").
										Body(
											app.I().
												Class("fas fa-code"),
										),
									app.H1().
										Class("pf-c-title pf-m-lg").
										Body(
											app.Text("Start editing"),
										),
									app.Div().
										Class("pf-c-empty-state__body").
										Body(
											app.Text("Drag a file here or browse to upload."),
										),
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("button").
										Body(
											app.Text("Browse"),
										),
									app.Div().
										Class("pf-c-empty-state__secondary").
										Body(
											app.Button().
												Class("pf-c-button pf-m-link").
												Type("button").
												Body(
													app.Text("Start from scratch"),
												),
										),
								),
						),
				),
		)
}
