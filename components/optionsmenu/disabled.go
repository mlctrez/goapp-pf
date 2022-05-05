package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("pf-c-options-menu").
		Body(
			app.Button().
				Class("pf-c-options-menu__toggle").
				Type("button").
				ID("options-menu-single-disabled-example-toggle").
				Aria("haspopup", "listbox").
				Aria("expanded", "false").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-options-menu__toggle-text").
						Body(
							app.Text("Disabled options menu"),
						),
					app.Div().
						Class("pf-c-options-menu__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
		)
}
