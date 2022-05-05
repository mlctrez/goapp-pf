package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MenuItemFavorites struct {
	app.Compo
}

func (c *MenuItemFavorites) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-favorites-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-favorites-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-favorites-label select-favorites-toggle").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Span().
								Class("pf-c-select__toggle-text").
								Body(
									app.Text("Favorites"),
								),
						),
					app.Span().
						Class("pf-c-select__toggle-arrow").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-favorites-label").
				Body(
					app.Div().
						Class("pf-c-select__menu-search").
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
					app.Hr().
						Class("pf-c-divider"),
					app.Div().
						Class("pf-c-select__menu-group").
						Body(
							app.Div().
								Class("pf-c-select__menu-group-title").
								Aria("hidden", true).
								ID("select-favorites-group-title-1").
								Body(
									app.Text("Favorites"),
								),
							app.Ul().
								Aria("role", "listbox").
								Aria("labelledby", "select-favorites-group-title-1").
								Body(
									app.Li().
										Class("pf-c-select__menu-wrapper pf-m-favorite").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link pf-m-description").
												Aria("role", "option").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-main").
														Body(
															app.Text("Item 1"),
														),
													app.Span().
														Class("pf-c-select__menu-item-description").
														Body(
															app.Text("This is a description."),
														),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-select__menu-wrapper pf-m-favorite").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link").
												Aria("role", "option").
												Body(
													app.Text("Item 4"),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
					app.Div().
						Class("pf-c-select__menu-group").
						Body(
							app.Div().
								Class("pf-c-select__menu-group-title").
								Aria("hidden", true).
								ID("select-favorites-group-title-2").
								Body(
									app.Text("Group 1"),
								),
							app.Ul().
								Aria("role", "listbox").
								Aria("labelledby", "select-favorites-group-title-2").
								Body(
									app.Li().
										Class("pf-c-select__menu-wrapper pf-m-favorite").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link pf-m-description").
												Aria("role", "option").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-main").
														Body(
															app.Text("Item 1"),
														),
													app.Span().
														Class("pf-c-select__menu-item-description").
														Body(
															app.Text("This is a description."),
														),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-select__menu-wrapper").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-selected pf-m-link").
												Aria("role", "option").
												Aria("selected", true).
												Body(
													app.Text("Item 2"), app.Span().
														Class("pf-c-select__menu-item-icon").
														Body(
															app.I().
																Class("fas fa-check").
																Aria("hidden", true),
														),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-select__menu-wrapper pf-m-disabled").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link").
												Aria("role", "option").
												Disabled(true).
												Body(
													app.Text("Item 3 (disabled)"),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Disabled(true).
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
					app.Div().
						Class("pf-c-select__menu-group").
						Body(
							app.Div().
								Class("pf-c-select__menu-group-title").
								Aria("hidden", true).
								ID("select-favorites-group-title-3").
								Body(
									app.Text("Group 2"),
								),
							app.Ul().
								Aria("role", "listbox").
								Aria("labelledby", "select-favorites-group-title-3").
								Body(
									app.Li().
										Class("pf-c-select__menu-wrapper pf-m-favorite").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link").
												Aria("role", "option").
												Body(
													app.Text("Item 4"),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Class("pf-c-select__menu-wrapper").
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-link pf-m-description").
												Aria("role", "option").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-main").
														Body(
															app.Text("Item 5"),
														),
													app.Span().
														Class("pf-c-select__menu-item-description").
														Body(
															app.Text("This is a description."),
														),
												),
											app.Button().
												Class("pf-c-select__menu-item pf-m-action pf-m-favorite-action").
												Aria("role", "option").
												Aria("label", "Not starred").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-action-icon").
														Body(
															app.I().
																Class("fas fa-star").
																Aria("hidden", true),
														),
												),
										),
								),
						),
				),
		)
}
