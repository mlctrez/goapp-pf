package actionlist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ActionListWithIcons struct {
	app.Compo
}

func (c *ActionListWithIcons) Render() app.UI {
	return app.Div().
		Class("pf-c-action-list pf-m-icons").
		Body(
			app.Div().
				Class("pf-c-action-list__item").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Close").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-action-list__item").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Kebab").
						Body(
							app.I().
								Class("fas fa-check").
								Aria("hidden", true),
						),
				),
		)
}
