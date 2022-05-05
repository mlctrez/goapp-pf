package tile

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StackedTiles struct {
	app.Compo
}

func (c *StackedTiles) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tile").
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
				Class("pf-c-tile pf-m-selected").
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
				Class("pf-c-tile pf-m-disabled").
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
				Class("pf-c-tile").
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
				Class("pf-c-tile pf-m-selected").
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
				Class("pf-c-tile pf-m-disabled").
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
				Class("pf-c-tile").
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
				Class("pf-c-tile").
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
