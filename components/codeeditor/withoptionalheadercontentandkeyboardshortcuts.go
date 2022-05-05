package codeeditor

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithOptionalHeaderContentAndKeyboardShortcuts struct {
	app.Compo
}

func (c *WithOptionalHeaderContentAndKeyboardShortcuts) Render() app.UI {
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
						Class("pf-c-code-editor__header-main").
						Body(
							app.Text("Header main content"),
						),
					app.Div().
						Class("pf-c-code-editor__keyboard-shortcuts").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-button__icon pf-m-start").
										Body(
											app.I().
												Class("pf-icon pf-icon-help").
												Aria("hidden", true),
										),
									app.Text("View shortcuts"),
								),
						),
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
