package clipboardcopy

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineCompact struct {
	app.Compo
}

func (c *InlineCompact) Render() app.UI {
	return app.Div().
		Class("pf-c-clipboard-copy pf-m-inline").
		Body(
			app.Span().
				Class("pf-c-clipboard-copy__text").
				Body(
					app.Text("2.3.4-2-redhat"),
				),
			app.Span().
				Class("pf-c-clipboard-copy__actions").
				Body(
					app.Span().
						Class("pf-c-clipboard-copy__actions-item").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Copy to clipboard").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
				),
		)
}
