package drawer

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandedInlinePanel struct {
	app.Compo
}

func (c *ExpandedInlinePanel) Render() app.UI {
	return app.Div().
		Class("pf-c-drawer pf-m-expanded pf-m-inline").
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
									app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pretium est a porttitor vehicula. Quisque vel commodo urna. Morbi mattis rutrum ante, id vehicula ex accumsan ut. Morbi viverra, eros vel porttitor facilisis, eros purus aliquet erat, nec lobortis felis elit pulvinar sem. Vivamus vulputate, risus eget commodo eleifend, eros nibh porta quam, vitae lacinia leo libero at magna. Maecenas aliquam sagittis orci, et posuere nisi ultrices sit amet. Aliquam ex odio, malesuada sed posuere quis, pellentesque at mauris. Phasellus venenatis massa ex, eget pulvinar libero auctor pretium. Aliquam erat volutpat. Duis euismod justo in quam ullamcorper, in commodo massa vulputate."),
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
											app.Span().
												Body(
													app.Text("drawer-panel"),
												),
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
										),
								),
						),
				),
		)
}
