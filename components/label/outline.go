package label

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Outline struct {
	app.Compo
}

func (c *Outline) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-label pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Grey"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline").
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
				Class("pf-c-label pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Grey removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-grey-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-grey-close-button outline-grey-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline").
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
						ID("outline-grey-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-grey-icon-close-button outline-grey-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Grey link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline").
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
						ID("outline-grey-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-grey-link-close-button outline-grey-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-blue pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Blue"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue pf-m-outline").
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
				Class("pf-c-label pf-m-blue pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Blue removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-blue-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-blue-close-button outline-blue-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-blue pf-m-outline").
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
						ID("outline-blue-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-blue-icon-close-button outline-blue-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-blue").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Blue link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-blue").
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
						ID("outline-blue-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-blue-link-close-button outline-blue-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-green pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Green"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green pf-m-outline").
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
				Class("pf-c-label pf-m-green pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Green removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-green-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-green-close-button outline-green-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-green pf-m-outline").
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
						ID("outline-green-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-green-icon-close-button outline-green-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-green").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Green link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-green").
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
						ID("outline-green-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-green-link-close-button outline-green-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-orange pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Orange"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange pf-m-outline").
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
				Class("pf-c-label pf-m-orange pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Orange removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-orange-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-orange-close-button outline-orange-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-orange pf-m-outline").
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
						ID("outline-orange-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-orange-icon-close-button outline-orange-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-orange").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Orange link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-orange").
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
						ID("outline-orange-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-orange-link-close-button outline-orange-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-red pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Red"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red pf-m-outline").
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
				Class("pf-c-label pf-m-red pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Red removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-red-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-red-close-button outline-red-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-red pf-m-outline").
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
						ID("outline-red-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-red-icon-close-button outline-red-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-red").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Red link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-red").
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
						ID("outline-red-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-red-link-close-button outline-red-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-purple pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Purple"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple pf-m-outline").
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
				Class("pf-c-label pf-m-purple pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Purple removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-purple-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-purple-close-button outline-purple-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-purple pf-m-outline").
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
						ID("outline-purple-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-purple-icon-close-button outline-purple-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-purple").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Purple link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-purple").
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
						ID("outline-purple-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-purple-link-close-button outline-purple-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Span().
				Class("pf-c-label pf-m-cyan pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Cyan"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan pf-m-outline").
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
				Class("pf-c-label pf-m-cyan pf-m-outline").
				Body(
					app.Span().
						Class("pf-c-label__content").
						Body(
							app.Text("Cyan removable"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						ID("outline-cyan-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-cyan-close-button outline-cyan-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-cyan pf-m-outline").
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
						ID("outline-cyan-icon-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-cyan-icon-close-button outline-cyan-icon-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-cyan").
				Body(
					app.A().
						Class("pf-c-label__content").
						Href("#").
						Body(
							app.Text("Cyan link"),
						),
				),
			app.Span().
				Class("pf-c-label pf-m-outline pf-m-cyan").
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
						ID("outline-cyan-link-close-button").
						Aria("label", "Remove").
						Aria("labelledby", "outline-cyan-link-close-button outline-cyan-link-close-text").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
		)
}
