package tile

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StackedTilesLarge struct {
	app.Compo
}

func (c *StackedTilesLarge) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tile pf-m-display-lg").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon").
								Body(
									app.I().
										Class("fas fa-bell").
										Aria("hidden", true),
								),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Default"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Div().
				Class("pf-c-tile pf-m-selected pf-m-display-lg").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon").
								Body(
									app.I().
										Class("fas fa-bell").
										Aria("hidden", true),
								),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Selected"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Div().
				Class("pf-c-tile pf-m-disabled pf-m-display-lg").
				TabIndex(-1).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon").
								Body(
									app.I().
										Class("fas fa-bell").
										Aria("hidden", true),
								),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Disabled"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tile pf-m-display-lg").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon"),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Default"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Div().
				Class("pf-c-tile pf-m-display-lg pf-m-selected").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon"),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Selected"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Div().
				Class("pf-c-tile pf-m-disabled pf-m-display-lg").
				TabIndex(-1).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon"),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Disabled"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-tile pf-m-display-lg").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon"),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Default"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
			app.Div().
				Class("pf-c-tile pf-m-display-lg").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tile__header pf-m-stacked").
						Body(
							app.Div().
								Class("pf-c-tile__icon"),
							app.Div().
								Class("pf-c-tile__title").
								Body(
									app.Text("Default"),
								),
						),
					app.Div().
						Class("pf-c-tile__body").
						Body(
							app.Text("Subtext goes here"),
						),
				),
		)
}
