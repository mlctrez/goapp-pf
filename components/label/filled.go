package label

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Filled struct {
	app.Compo
}

func (c *Filled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-label").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Grey"),
						),
				),
			app.Span().
				Class("pf-c-label").
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
							app.Text("Grey icon"),
						),
				),
			app.Span().
				Class("pf-c-label").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Grey removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-grey-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-grey-close-button default-grey-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label").
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
							app.Text("Grey icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-grey-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-grey-icon-close-button default-grey-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Grey link"),
						),
				),
			app.Span().
				Class("pf-c-label").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Grey link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-grey-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-grey-link-close-button default-grey-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label").
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
									app.Text("Grey label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-grey-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-grey-icon-close-truncate-button default-grey-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-blue").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Blue"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
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
							app.Text("Blue icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Blue removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-blue-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-blue-close-button default-blue-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
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
							app.Text("Blue icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-blue-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-blue-icon-close-button default-blue-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Blue link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Blue link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-blue-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-blue-link-close-button default-blue-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue").
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
									app.Text("Blue label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-blue-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-blue-icon-close-truncate-button default-blue-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-green").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Green"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
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
							app.Text("Green icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Green removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-green-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-green-close-button default-green-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
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
							app.Text("Green icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-green-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-green-icon-close-button default-green-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Green link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Green link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-green-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-green-link-close-button default-green-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green").
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
									app.Text("Green label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-green-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-green-icon-close-truncate-button default-green-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-orange").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Orange"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
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
							app.Text("Orange icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Orange removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-orange-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-orange-close-button default-orange-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
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
							app.Text("Orange icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-orange-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-orange-icon-close-button default-orange-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Orange link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Orange link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-orange-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-orange-link-close-button default-orange-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange").
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
									app.Text("Orange label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-orange-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-orange-icon-close-truncate-button default-orange-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-red").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Red"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
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
							app.Text("Red icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Red removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-red-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-red-close-button default-red-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
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
							app.Text("Red icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-red-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-red-icon-close-button default-red-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Red link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Red link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-red-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-red-link-close-button default-red-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red").
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
									app.Text("Red label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-red-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-red-icon-close-truncate-button default-red-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-purple").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Purple"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
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
							app.Text("Purple icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Purple removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-purple-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-purple-close-button default-purple-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
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
							app.Text("Purple icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-purple-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-purple-icon-close-button default-purple-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Purple link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Purple link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-purple-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-purple-link-close-button default-purple-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple").
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
									app.Text("Purple label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-purple-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-purple-icon-close-truncate-button default-purple-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-cyan").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Cyan"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
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
							app.Text("Cyan icon"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Cyan removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-cyan-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-cyan-close-button default-cyan-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
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
							app.Text("Cyan icon removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-cyan-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-cyan-icon-close-button default-cyan-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Cyan link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Cyan link removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-cyan-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-cyan-link-close-button default-cyan-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan").
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
									app.Text("Cyan label with icon that truncates"),
								),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("default-cyan-icon-close-truncate-button").
						Aria("label", "Remove").
						Aria("labelledby", "default-cyan-icon-close-truncate-button default-cyan-icon-close-truncate-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
		)
}
