package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SplitButtonAction struct {
	app.Compo
}

func (c *SplitButtonAction) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Div().
						Class("pf-c-dropdown__toggle pf-m-split-button pf-m-action").
						Body(
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("label", "Dropdown toggle").
								Body(
									app.Text("Action"),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", "false").
								ID("dropdown-split-button-action-toggle-button").
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
						Class("pf-c-dropdown__toggle pf-m-split-button pf-m-action").
						Body(
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("label", "Dropdown toggle").
								Body(
									app.Text("Action"),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", true).
								ID("dropdown-split-button-action-expanded-toggle-button").
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
			app.Div().
				Class("pf-c-dropdown").
				Body(
					app.Div().
						Class("pf-c-dropdown__toggle pf-m-split-button pf-m-action").
						Body(
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("label", "Settings").
								Body(
									app.I().
										Class("fas fa-cog").
										Aria("hidden", true),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", "false").
								ID("dropdown-split-button-action-icon-toggle-button").
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
						Class("pf-c-dropdown__toggle pf-m-split-button pf-m-action").
						Body(
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("label", "Settings").
								Body(
									app.I().
										Class("fas fa-cog").
										Aria("hidden", true),
								),
							app.Button().
								Class("pf-c-dropdown__toggle-button").
								Type("button").
								Aria("expanded", true).
								ID("dropdown-split-button-action-icon-expanded-toggle-button").
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
										Class("pf-c-dropdown__menu-item pf-m-icon").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-dropdown__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-cog").
														Aria("hidden", true),
												),
											app.Text("Actions"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item pf-m-icon").
										Type("button").
										Disabled(true).
										Body(
											app.Span().
												Class("pf-c-dropdown__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-bell").
														Aria("hidden", true),
												),
											app.Text("Disabled action"),
										),
								),
							app.Li().
								Body(
									app.Button().
										Class("pf-c-dropdown__menu-item pf-m-icon").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-dropdown__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-cubes").
														Aria("hidden", true),
												),
											app.Text("Other action"),
										),
								),
						),
				),
		)
}
