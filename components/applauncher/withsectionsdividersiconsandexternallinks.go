package applauncher

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithSectionsDividersIconsAndExternalLinks struct {
	app.Compo
}

func (c *WithSectionsDividersIconsAndExternalLinks) Render() app.UI {
	return app.Nav().
		Class("pf-c-app-launcher pf-m-expanded").
		Aria("label", "Application launcher").
		Body(
			app.Button().
				Class("pf-c-app-launcher__toggle").
				Type("button").
				ID("-button").
				Aria("expanded", true).
				Aria("label", "Application launcher").
				Body(
					app.I().
						Class("fas fa-th").
						Aria("hidden", true),
				),
			app.Div().
				Class("pf-c-app-launcher__menu").
				Body(
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item").
												Href("#").
												Body(
													app.Span().
														Class("pf-c-app-launcher__menu-item-icon").
														Body(
															app.Img().
																Src("/assets/images/pf-logo-small.svg").
																Alt("PatternFly logo"),
														),
													app.Text("Link not in group"),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-divider").
						Aria("role", "separator"),
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.H1().
								Class("pf-c-app-launcher__group-title").
								Body(
									app.Text("Group 1"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-external").
												Href("#").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-app-launcher__menu-item-icon").
														Body(
															app.Img().
																Src("/assets/images/pf-logo-small.svg").
																Alt("PatternFly logo"),
														),
													app.Text("Group 1 link"),
													app.Span().
														Class("pf-c-app-launcher__menu-item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-external").
												Href("#").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-app-launcher__menu-item-icon").
														Body(
															app.Img().
																Src("/assets/images/pf-logo-small.svg").
																Alt("PatternFly logo"),
														),
													app.Text("Group 1 link"),
													app.Span().
														Class("pf-c-app-launcher__menu-item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
									app.Li().
										Class("pf-c-divider").
										Aria("role", "separator"),
								),
						),
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.H1().
								Class("pf-c-app-launcher__group-title").
								Body(
									app.Text("Group 2"),
								),
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-external").
												Href("#").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-app-launcher__menu-item-icon").
														Body(
															app.Img().
																Src("/assets/images/pf-logo-small.svg").
																Alt("PatternFly logo"),
														),
													app.Text("Group 2 link"),
													app.Span().
														Class("pf-c-app-launcher__menu-item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-external").
												Href("#").
												Target("_blank").
												Body(
													app.Span().
														Class("pf-c-app-launcher__menu-item-icon").
														Body(
															app.Img().
																Src("/assets/images/pf-logo-small.svg").
																Alt("PatternFly logo"),
														),
													app.Text("Group 2 link"),
													app.Span().
														Class("pf-c-app-launcher__menu-item-external-icon").
														Body(
															app.I().
																Class("fas fa-external-link-alt").
																Aria("hidden", true),
														),
													app.Span().
														Class("pf-screen-reader").
														Body(
															app.Text("(opens new window)"),
														),
												),
										),
								),
						),
				),
		)
}
