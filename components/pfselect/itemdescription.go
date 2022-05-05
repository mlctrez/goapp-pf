package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ItemDescription struct {
	app.Compo
}

func (c *ItemDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-with-description-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-with-description-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-with-description-label select-with-description-toggle").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Span().
								Class("pf-c-select__toggle-text").
								Body(
									app.Text("Select with description"),
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
			app.Ul().
				Class("pf-c-select__menu").
				Aria("role", "listbox").
				Aria("labelledby", "select-with-description-label").
				Body(
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("Menu item"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item pf-m-selected").
								Aria("role", "option").
								Aria("selected", true).
								Body(
									app.Text("Menu item selected"), app.Span().
										Class("pf-c-select__menu-item-icon").
										Body(
											app.I().
												Class("fas fa-check").
												Aria("hidden", true),
										),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item pf-m-description").
								Aria("role", "option").
								Body(
									app.Span().
										Class("pf-c-select__menu-item-main").
										Body(
											app.Text("Menu item description"),
										),
									app.Span().
										Class("pf-c-select__menu-item-description").
										Body(
											app.Text("This is a description."),
										),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item pf-m-description").
								Aria("role", "option").
								Disabled(true).
								Body(
									app.Span().
										Class("pf-c-select__menu-item-main").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-main").
												Body(
													app.Text("Menu item description disabled"),
												),
										),
									app.Span().
										Class("pf-c-select__menu-item-description").
										Body(
											app.Text("This is a description."),
										),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item pf-m-selected pf-m-description").
								Aria("role", "option").
								Aria("selected", true).
								Body(
									app.Span().
										Class("pf-c-select__menu-item-main").
										Body(
											app.Text("Menu item description selected"), app.Span().
												Class("pf-c-select__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
									app.Span().
										Class("pf-c-select__menu-item-description").
										Body(
											app.Text("This is a description."),
										),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item pf-m-selected pf-m-description").
								Aria("role", "option").
								Aria("selected", true).
								Body(
									app.Span().
										Class("pf-c-select__menu-item-main").
										Body(
											app.Text("Menu item long description"), app.Span().
												Class("pf-c-select__menu-item-icon").
												Body(
													app.I().
														Class("fas fa-check").
														Aria("hidden", true),
												),
										),
									app.Span().
										Class("pf-c-select__menu-item-description").
										Body(
											app.Text("This is a really long description that describes the menu item. This is a really long description that describes the menu item."),
										),
								),
						),
				),
		)
}
