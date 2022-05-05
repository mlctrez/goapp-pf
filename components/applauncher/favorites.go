package applauncher

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Favorites struct {
	app.Compo
}

func (c *Favorites) Render() app.UI {
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
					app.Div().
						Class("pf-c-app-launcher__menu-search").
						Body(
							app.Div().
								Class("pf-c-search-input").
								Body(
									app.Div().
										Class("pf-c-search-input__bar").
										Body(
											app.Span().
												Class("pf-c-search-input__text").
												Body(
													app.Span().
														Class("pf-c-search-input__icon").
														Body(
															app.I().
																Class("fas fa-search fa-fw").
																Aria("hidden", true),
														),
													app.Input().
														Class("pf-c-search-input__text-input").
														Type("text").
														Placeholder("false").
														Aria("label", "Search"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-app-launcher__group").
						Body(
							app.H1().
								Class("pf-c-app-launcher__group-title").
								Body(
									app.Text("Favorites"),
								),
							app.Ul().
								Body(
									app.Li().
										Class("pf-c-app-launcher__menu-wrapper pf-m-external pf-m-favorite").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 2"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
									app.Li().
										Class("pf-c-app-launcher__menu-wrapper pf-m-external pf-m-favorite").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 3"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
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
										Class("pf-c-app-launcher__menu-wrapper pf-m-external").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 1"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
									app.Li().
										Class("pf-c-app-launcher__menu-wrapper pf-m-external pf-m-favorite").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 2"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
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
										Class("pf-c-app-launcher__menu-wrapper pf-m-external pf-m-favorite").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 3"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
									app.Li().
										Class("pf-c-app-launcher__menu-wrapper pf-m-external").
										Body(
											app.A().
												Class("pf-c-app-launcher__menu-item pf-m-link").
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
													app.Text("Link 4"),
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
											app.Button().
												Class("pf-c-app-launcher__menu-item pf-m-action").
												Type("button").
												Aria("label", "Favorite").
												Body(
													app.I().
														Class("fas fa-star").
														Aria("hidden", true),
												),
										),
								),
						),
				),
		)
}
