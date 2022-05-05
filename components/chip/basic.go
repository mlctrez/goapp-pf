package chip

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-chip").
				Body(
					app.Span().
						Class("pf-c-chip__text").
						ID("chip_one").
						Body(
							app.Text("Chip"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("labelledby", "remove_chip_one chip_one").
						Aria("label", "Remove").
						ID("remove_chip_one").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip").
				Body(
					app.Span().
						Class("pf-c-chip__text").
						ID("chip_two").
						Body(
							app.Text("Really long chip that goes on and on"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("labelledby", "remove_chip_two chip_two").
						Aria("label", "Remove").
						ID("remove_chip_two").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip").
				Body(
					app.Span().
						Class("pf-c-chip__text").
						ID("chip_three").
						Body(
							app.Text("Chip"),
						),
					app.Span().
						Class("pf-c-badge pf-m-read").
						Body(
							app.Text("00"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("labelledby", "remove_chip_three chip_three").
						Aria("label", "Remove").
						ID("remove_chip_three").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip").
				Body(
					app.Span().
						Class("pf-c-chip__text").
						Body(
							app.Text("Read-only chip"),
						),
				),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-chip pf-m-overflow").
				Body(
					app.Span().
						Class("pf-c-chip__text").
						Body(
							app.Text("Overflow chip"),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip pf-m-draggable").
				Body(
					app.Span().
						Class("pf-c-chip__icon").
						Body(
							app.I().
								Class("fas fa-grip-vertical").
								Aria("role", "img").
								Aria("label", "Drag"),
						),
					app.Span().
						Class("pf-c-chip__text").
						Body(
							app.Text("Draggable chip"),
						),
				),
		)
}
