package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonCheckbox struct {
	app.Compo
}

func (c *SplitButtonCheckbox) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Div().
						Class("pf-c-dropdown__toggle pf-m-disabled pf-m-split-button").
						Body(
							app.Label().
								Class("pf-c-dropdown__toggle-check").
								For("dropdown-split-button-disabled-toggle-check").
								Body(
									app.Input().
										Type("checkbox").
										ID("dropdown-split-button-disabled-toggle-check").
										Disabled(true).
										Aria("label", "Select all"),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", "false").
								ID("dropdown-split-button-disabled-toggle-button").
								Aria("label", "Dropdown toggle").
								Disabled(true).
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
											app.Text("Actions"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item").
										Type("button").
										Disabled(true).
										Body(
											app.Text("Disabled action"),
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
				),
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Div().
						Class("pf-c-dropdown__toggle pf-m-split-button").
						Body(
							app.Label().
								Class("pf-c-dropdown__toggle-check").
								For("dropdown-split-button-toggle-check").
								Body(
									app.Input().
										Type("checkbox").
										ID("dropdown-split-button-toggle-check").
										Aria("label", "Select all"),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", "false").
								ID("dropdown-split-button-toggle-button").
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
											app.Text("Actions"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item").
										Type("button").
										Disabled(true).
										Body(
											app.Text("Disabled action"),
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
				),
			app.Div().
				Class("pf-c-dropdown pf-m-expanded").
				Body(
					app.Div().
						Class("pf-c-dropdown__toggle pf-m-split-button").
						Body(
							app.Label().
								Class("pf-c-dropdown__toggle-check").
								For("dropdown-split-button-expanded-toggle-check").
								Body(
									app.Input().
										Type("checkbox").
										ID("dropdown-split-button-expanded-toggle-check").
										Aria("label", "Select all"),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", true).
								ID("dropdown-split-button-expanded-toggle-button").
								Aria("label", "Dropdown toggle").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
					app.Ul().
						Class("pf-c-dropdown__menu").
						Body(
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item").
										Type("button").
										Body(
											app.Text("Actions"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item").
										Type("button").
										Disabled(true).
										Body(
											app.Text("Disabled action"),
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
				),
		)
}
