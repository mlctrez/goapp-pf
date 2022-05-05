package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type PanelMenu struct {
	app.Compo
}

func (c *PanelMenu) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-panel-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-panel-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-panel-label select-panel-toggle").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Span().
								Class("pf-c-select__toggle-text").
								Body(
									app.Text("Filter by status"),
								),
						),
					app.Span().
						Class("pf-c-select__toggle-arrow").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-panel-label").
				Body(
					app.Text("[Panel contents here]"),
				),
		)
}
