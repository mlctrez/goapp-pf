package codeeditor

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Default struct {
	app.Compo
}

func (c *Default) Render() app.UI {
	return app.Div().
		Class("pf-c-code-editor").
		Body(
			app.Div().
				Class("pf-c-code-editor__header").
				Body(
					app.Div().
						Class("pf-c-code-editor__controls").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Download code").
								Body(
									app.I().
										Class("fas fa-download"),
								),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Upload code").
								Body(
									app.I().
										Class("fas fa-upload"),
								),
						),
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
									app.Text("HTML"),
								),
						),
				),
			app.Div().
				Class("pf-c-code-editor__main").
				Body(
					app.Code().
						Class("pf-c-code-editor__code").
						Body(
							app.Pre().
								Class("pf-c-code-editor__code-pre").
								Body(
									app.Text("code goes here"),
								),
						),
				),
		)
}
