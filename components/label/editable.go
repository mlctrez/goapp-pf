package label

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Editable struct {
	app.Compo
}

func (c *Editable) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-label pf-m-blue pf-m-editable").
				Body(
					app.Button().
						Class("pf-c-label__content").
						ID("editable-label-editable-content").
						Value("Editable label").
						Aria("label", "Editable text").
						Body(
							app.Text("Editable label"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("editable-label-button").
						Aria("label", "Remove").
						Aria("labelledby", "editable-label-button editable-label-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue pf-m-editable pf-m-editable-active").
				Body(
					app.Input().
						Class("pf-c-label__content").
						ID("editable-label-active-editable-content").
						Type("text").
						Value("Editable active").
						Aria("label", "Editable text"),
				),
			app.Span().
				Class("pf-c-label pf-m-compact pf-m-blue pf-m-editable").
				Body(
					app.Button().
						Class("pf-c-label__content").
						ID("compact-editable-label-editable-content").
						Value("Compact editable label").
						Aria("label", "Editable text").
						Body(
							app.Text("Compact editable label"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("compact-editable-label-button").
						Aria("label", "Remove").
						Aria("labelledby", "compact-editable-label-button compact-editable-label-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact pf-m-blue pf-m-editable pf-m-editable-active").
				Body(
					app.Input().
						Class("pf-c-label__content").
						ID("compact-editable-label-active-editable-content").
						Type("text").
						Value("Compact editable active").
						Aria("label", "Editable text"),
				),
		)
}
