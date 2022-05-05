package optionsmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AlignRight struct {
	app.Compo
}

func (c *AlignRight) Render() app.UI {
	return app.Div().
		Class("pf-c-options-menu pf-m-expanded pf-m-align-right").
		Body(
			app.Button().
				Class("pf-c-options-menu__toggle").
				Type("button").
				ID("options-menu-align-right-example-toggle").
				Aria("haspopup", "listbox").
				Aria("expanded", true).
				Body(
					app.Span().
						Class("pf-c-options-menu__toggle-text").
						Body(
							app.Text("Options menu"),
						),
					app.Div().
						Class("pf-c-options-menu__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-options-menu__menu pf-m-align-right").
				Aria("labelledby", "options-menu-align-right-example-toggle").
				Body(
					app.Li().
						Body(
							app.Button().
								Class("pf-c-options-menu__menu-item").
								Type("button").
								Body(
									app.Text("Option 1"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-options-menu__menu-item").
								Type("button").
								Body(
									app.Text("Option 2"), app.Div().
										Class("pf-c-options-menu__menu-item-icon").
										Body(
											app.I().
												Class("fas fa-check").
												Aria("hidden", true),
										),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-options-menu__menu-item").
								Type("button").
								Body(
									app.Text("Option 3"),
								),
						),
				),
		)
}
