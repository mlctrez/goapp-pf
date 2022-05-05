package clipboardcopy

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineCompactInSentence struct {
	app.Compo
}

func (c *InlineCompactInSentence) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.H4().
				Body(
					app.Strong().
						Body(
							app.Text("Basic"),
						),
				),
			app.Text("Lorem ipsum"),
			app.Div().
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
				),
			app.Text("dolor sit amet."),
			app.Br(),
			app.Br(),
			app.H4().
				Body(
					app.Strong().
						Body(
							app.Text("Long copy string"),
						),
				),
			app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			app.Div().
				Class("pf-c-clipboard-copy pf-m-inline").
				Body(
					app.Span().
						Class("pf-c-clipboard-copy__text").
						Body(
							app.Text("https://app.openshift.io/path/sub-path/sub-sub-path/?runtime=quarkus/12345678901234567890/abcdefghijklmnopqrstuvwxyz1234567890"),
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
				),
			app.Text("Mauris luctus, libero nec dapibus ultricies, urna purus pretium mauris, ullamcorper pharetra lacus nibh vitae enim."),
			app.Br(),
			app.Br(),
			app.H4().
				Body(
					app.Strong().
						Body(
							app.Text("Long copy string in block"),
						),
				),
			app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			app.Div().
				Class("pf-c-clipboard-copy pf-m-inline pf-m-block").
				Body(
					app.Span().
						Class("pf-c-clipboard-copy__text").
						Body(
							app.Text("https://app.openshift.io/path/sub-path/sub-sub-path/?runtime=quarkus/12345678901234567890/abcdefghijklmnopqrstuvwxyz1234567890"),
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
				),
			app.Text("Mauris luctus, libero nec dapibus ultricies, urna purus pretium mauris, ullamcorper pharetra lacus nibh vitae enim."),
		)
}
