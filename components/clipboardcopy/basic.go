package clipboardcopy

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-clipboard-copy").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								Value("This is editable").
								ID("basic-editable-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("basic-editable-copy-button").
								Aria("labelledby", "basic-editable-copy-button basic-editable-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-clipboard-copy").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								ReadOnly(true).
								Type("text").
								Value("This is read-only").
								ID("basic-readonly-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("basic-readonly-copy-button").
								Aria("labelledby", "basic-readonly-copy-button basic-readonly-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
				),
		)
}
