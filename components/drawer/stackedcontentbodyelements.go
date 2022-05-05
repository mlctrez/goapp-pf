package drawer

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StackedContentBodyElements struct {
	app.Compo
}

func (c *StackedContentBodyElements) Render() app.UI {
	return app.Div().
		Class("pf-c-drawer pf-m-expanded").
		Body(
			app.Div().
				Class("pf-c-drawer__main").
				Body(
					app.Div().
						Class("pf-c-drawer__content").
						Body(
							app.Div().
								Class("pf-c-drawer__body").
								Body(
									app.Text("content-body"),
								),
							app.Div().
								Class("pf-c-drawer__body pf-m-padding").
								Body(
									app.Text("content-body with padding"),
								),
							app.Div().
								Class("pf-c-drawer__body").
								Body(
									app.Text("content-body"),
								),
						),
					app.Div().
						Class("pf-c-drawer__panel").
						Body(
							app.Div().
								Class("pf-c-drawer__body").
								Body(
									app.Div().
										Class("pf-c-drawer__head").
										Body(
											app.Div().
												Class("pf-c-drawer__actions").
												Body(
													app.Div().
														Class("pf-c-drawer__close").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Close drawer panel").
																Body(
																	app.I().
																		Class("fas fa-times").
																		Aria("hidden", true),
																),
														),
												),
											app.Text("drawer-panel"),
										),
								),
							app.Div().
								Class("pf-c-drawer__body pf-m-no-padding").
								Body(
									app.Text("drawer-panel with no padding"),
								),
							app.Div().
								Class("pf-c-drawer__body").
								Body(
									app.Text("drawer-panel"),
								),
						),
				),
		)
}
