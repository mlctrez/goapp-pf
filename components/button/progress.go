package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Progress struct {
	app.Compo
}

func (c *Progress) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-progress pf-m-primary").
				Type("button").
				Body(
					app.Text("Primary loader"),
				),
			app.Button().
				Class("pf-c-button pf-m-progress pf-m-in-progress pf-m-primary").
				Type("button").
				Body(
					app.Span().
						Class("pf-c-button__progress").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-md").
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
					app.Text("Primary loading"),
				),
			app.Button().
				Class("pf-c-button pf-m-progress pf-m-secondary").
				Type("button").
				Body(
					app.Text("Secondary loader"),
				),
			app.Button().
				Class("pf-c-button pf-m-progress pf-m-in-progress pf-m-secondary").
				Type("button").
				Body(
					app.Span().
						Class("pf-c-button__progress").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-md").
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
					app.Text("Secondary loading"),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Upload").
				Body(
					app.I().
						Class("fas fa-upload").
						Aria("hidden", true),
				),
			app.Button().
				Class("pf-c-button pf-m-plain pf-m-in-progress").
				Type("button").
				Aria("label", "Upload").
				Body(
					app.I().
						Class("fas fa-upload").
						Aria("hidden", true),
					app.Span().
						Class("pf-c-button__progress").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-md").
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
