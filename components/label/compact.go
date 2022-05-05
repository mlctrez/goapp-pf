package label

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Compact"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Span().
								Class("pf-c-label__icon").
								Body(
									app.I().
										Class("fas fa-fw fa-info-circle").
										Aria("hidden", true),
								),
							app.Text("Compact icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Compact removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("compact-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "compact-close-button compact-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Span().
								Class("pf-c-label__icon").
								Body(
									app.I().
										Class("fas fa-fw fa-info-circle").
										Aria("hidden", true),
								),
							app.Text("Compact icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("compact-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "compact-icon-close-button compact-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Compact link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Compact link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("compact-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "compact-link-close-button compact-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-compact").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Span().
								Class("pf-c-label__icon").
								Body(
									app.I().
										Class("fas fa-fw fa-info-circle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-label__text").
								Body(
									app.Text("Compact label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("compact-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "compact-icon-close-truncate-button compact-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
		)
}
