package tile

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExtraContent struct {
	app.Compo
}

func (c *ExtraContent) Render() app.UI {
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
							app.Text("This is really really long subtext that goes on for so long that it has to wrap to the next line. This is really really long subtext that goes on for so long that it has to wrap to the next line."),
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
							app.Text("This is really really long subtext that goes on for so long that it has to wrap to the next line."),
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
		)
}
