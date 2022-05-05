package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonCheckboxWithToggleText struct {
	app.Compo
}

func (c *SplitButtonCheckboxWithToggleText) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown").
		Body(
			app.Div().
				Class("pf-c-dropdown__toggle pf-m-split-button").
				Body(
					app.Label().
						Class("pf-c-dropdown__toggle-check").
						For("dropdown-split-button-text-toggle-check").
						Body(
							app.Input().
								Type("checkbox").
								ID("dropdown-split-button-text-toggle-check").
								Aria("label", "Select all").
								Checked(true).
								Aria("labelledby", "dropdown-split-button-text-toggle-check dropdown-split-button-text-toggle-check-text"),
							app.Span().
								Class("pf-c-dropdown__toggle-text").
								Aria("hidden", true).
								ID("dropdown-split-button-text-toggle-check-text").
								Body(
									app.Text("10 selected"),
								),
						),
					app.Button().
						Class("pf-c-dropdown__toggle-button").
						Type("button").
						Aria("expanded", "false").
						ID("dropdown-split-button-text-toggle-button").
						Aria("label", "Dropdown toggle").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-dropdown__menu").
				Hidden(true).
				Body(
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Select all"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Select none"),
								),
						),
					app.Li().
						Body(
							app.Button().
								Class("pf-c-dropdown__menu-item").
								Type("button").
								Body(
									app.Text("Other action"),
								),
						),
				),
		)
}
