package pagination

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BottomSticky struct {
	app.Compo
}

func (c *BottomSticky) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pretium est a porttitor vehicula. Quisque vel commodo urna. Morbi mattis rutrum ante, id vehicula ex accumsan ut. Morbi viverra, eros vel porttitor facilisis, eros purus aliquet erat, nec lobortis felis elit pulvinar sem. Vivamus vulputate, risus eget commodo eleifend, eros nibh porta quam, vitae lacinia leo libero at magna. Maecenas aliquam sagittis orci, et posuere nisi ultrices sit amet. Aliquam ex odio, malesuada sed posuere quis, pellentesque at mauris. Phasellus venenatis massa ex, eget pulvinar libero auctor pretium. Aliquam erat volutpat. Duis euismod justo in quam ullamcorper, in commodo massa vulputate."),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pretium est a porttitor vehicula. Quisque vel commodo urna. Morbi mattis rutrum ante, id vehicula ex accumsan ut. Morbi viverra, eros vel porttitor facilisis, eros purus aliquet erat, nec lobortis felis elit pulvinar sem. Vivamus vulputate, risus eget commodo eleifend, eros nibh porta quam, vitae lacinia leo libero at magna. Maecenas aliquam sagittis orci, et posuere nisi ultrices sit amet. Aliquam ex odio, malesuada sed posuere quis, pellentesque at mauris. Phasellus venenatis massa ex, eget pulvinar libero auctor pretium. Aliquam erat volutpat. Duis euismod justo in quam ullamcorper, in commodo massa vulputate."),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pretium est a porttitor vehicula. Quisque vel commodo urna. Morbi mattis rutrum ante, id vehicula ex accumsan ut. Morbi viverra, eros vel porttitor facilisis, eros purus aliquet erat, nec lobortis felis elit pulvinar sem. Vivamus vulputate, risus eget commodo eleifend, eros nibh porta quam, vitae lacinia leo libero at magna. Maecenas aliquam sagittis orci, et posuere nisi ultrices sit amet. Aliquam ex odio, malesuada sed posuere quis, pellentesque at mauris. Phasellus venenatis massa ex, eget pulvinar libero auctor pretium. Aliquam erat volutpat. Duis euismod justo in quam ullamcorper, in commodo massa vulputate."),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus pretium est a porttitor vehicula. Quisque vel commodo urna. Morbi mattis rutrum ante, id vehicula ex accumsan ut. Morbi viverra, eros vel porttitor facilisis, eros purus aliquet erat, nec lobortis felis elit pulvinar sem. Vivamus vulputate, risus eget commodo eleifend, eros nibh porta quam, vitae lacinia leo libero at magna. Maecenas aliquam sagittis orci, et posuere nisi ultrices sit amet. Aliquam ex odio, malesuada sed posuere quis, pellentesque at mauris. Phasellus venenatis massa ex, eget pulvinar libero auctor pretium. Aliquam erat volutpat. Duis euismod justo in quam ullamcorper, in commodo massa vulputate."),
				),
			app.Div().
				Class("pf-c-pagination pf-m-bottom pf-m-sticky").
				Body(
					app.Div().
						Class("pf-c-options-menu pf-m-top").
						Body(
							app.Button().
								Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
								Type("button").
								ID("pagination-options-menu-bottom-sticky-example-toggle").
								Aria("haspopup", "listbox").
								Aria("expanded", "false").
								Body(
									app.Span().
										Class("pf-c-options-menu__toggle-text").
										Body(
											app.B().
												Body(
													app.Text("1 - 10"),
												),
											app.Text("of"),
											app.B().
												Body(
													app.Text("36"),
												),
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
								Class("pf-c-options-menu__menu pf-m-top").
								Aria("labelledby", "pagination-options-menu-bottom-sticky-example-toggle").
								Hidden(true).
								Body(
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("5 per page"),
												),
										),
									app.Li().
										Body(
											app.Button().
												Class("pf-c-options-menu__menu-item").
												Type("button").
												Body(
													app.Text("10 per page"), app.Div().
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
													app.Text("20 per page"),
												),
										),
								),
						),
					app.Nav().
						Class("pf-c-pagination__nav").
						Aria("label", "Pagination").
						Body(
							app.Div().
								Class("pf-c-pagination__nav-control pf-m-first").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
										Aria("label", "Go to first page").
										Body(
											app.I().
												Class("fas fa-angle-double-left").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-pagination__nav-control pf-m-prev").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Disabled(true).
										Aria("label", "Go to previous page").
										Body(
											app.I().
												Class("fas fa-angle-left").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-pagination__nav-page-select").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Aria("label", "Current page").
										Type("number").
										Min("1").
										Max("4").
										Value("1"),
									app.Span().
										Aria("hidden", true).
										Body(
											app.Text("of 4"),
										),
								),
							app.Div().
								Class("pf-c-pagination__nav-control pf-m-next").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Go to next page").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-pagination__nav-control pf-m-last").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Go to last page").
										Body(
											app.I().
												Class("fas fa-angle-double-right").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
