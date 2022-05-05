package tabcontent

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Padding struct {
	app.Compo
}

func (c *Padding) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Section().
				Class("pf-c-tab-content").
				ID("tab1-panel").
				Aria("role", "tabpanel").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-tab-content__body pf-m-padding").
						Body(
							app.Text("Panel 1"),
						),
				),
			app.Section().
				Class("pf-c-tab-content").
				ID("tab2-panel").
				Aria("role", "tabpanel").
				TabIndex(0).
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-tab-content__body pf-m-padding").
						Body(
							app.Text("Panel 2"),
						),
				),
			app.Section().
				Class("pf-c-tab-content").
				ID("tab3-panel").
				Aria("role", "tabpanel").
				TabIndex(0).
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-tab-content__body pf-m-padding").
						Body(
							app.Text("Panel 3"),
						),
				),
			app.Section().
				Class("pf-c-tab-content").
				ID("tab4-panel").
				Aria("role", "tabpanel").
				TabIndex(0).
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-tab-content__body pf-m-padding").
						Body(
							app.Text("Panel 4"),
						),
				),
		)
}
