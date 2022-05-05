package hint

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DefaultWithNoTitle struct {
	app.Compo
}

func (c *DefaultWithNoTitle) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-hint").
				Body(
					app.Div().
						Class("pf-c-hint__body").
						Body(
							app.Text("Welcome to the new documentation experience."), app.A().
								Href("#").
								Body(
									app.Text("Learn more about the improved features"),
								),
							app.Text("."),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-hint").
				Body(
					app.Div().
						Class("pf-c-hint__actions").
						Body(
							app.Div().
								Class("pf-c-dropdown").
								Body(
									app.Button().
										Class("pf-c-dropdown__toggle pf-m-plain").
										ID("hint-with-no-title-dropdown-kebab-button").
										Aria("expanded", "false").
										Type("button").
										Aria("label", "Actions").
										Body(
											app.I().
												Class("fas fa-ellipsis-v").
												Aria("hidden", true),
										),
									app.Ul().
										Class("pf-c-dropdown__menu").
										Aria("labelledby", "hint-with-no-title-dropdown-kebab-button").
										Hidden(true).
										Body(
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Link"),
														),
												),
											app.Li().
												Body(
													app.Button().
														Class("pf-c-dropdown__menu-item").
														Type("button").
														Body(
															app.Text("Action"),
														),
												),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item pf-m-disabled").
														Href("#").
														Aria("disabled", true).
														TabIndex(-1).
														Body(
															app.Text("Disabled link"),
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
												Class("pf-c-divider").
												Aria("role", "separator"),
											app.Li().
												Body(
													app.A().
														Class("pf-c-dropdown__menu-item").
														Href("#").
														Body(
															app.Text("Separated link"),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-hint__body").
						Body(
							app.Text("Upgrade to Red Hat Smart Management to remediate all your systems across regions and geographies."),
						),
					app.Div().
						Class("pf-c-hint__footer").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("Try it for 90 days"),
								),
						),
				),
		)
}
