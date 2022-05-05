package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MenuFooter struct {
	app.Compo
}

func (c *MenuFooter) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-single-footer-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-single-footer-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-single-footer-label select-single-footer-toggle").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Span().
								Class("pf-c-select__toggle-text").
								Body(
									app.Text("Filter by status"),
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
				Body(
					app.Ul().
						Class("pf-c-select__menu-list").
						Aria("role", "listbox").
						Aria("labelledby", "select-single-footer-label").
						Body(
							app.Li().
								Aria("role", "presentation").
								Body(
									app.Button().
										Class("pf-c-select__menu-item").
										Aria("role", "option").
										Body(
											app.Text("Running"),
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
											app.Text("Stopped"), app.Span().
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
										Class("pf-c-select__menu-item").
										Aria("role", "option").
										Body(
											app.Text("Down"),
										),
								),
							app.Li().
								Aria("role", "presentation").
								Body(
									app.Button().
										Class("pf-c-select__menu-item").
										Aria("role", "option").
										Body(
											app.Text("Degraded"),
										),
								),
							app.Li().
								Aria("role", "presentation").
								Body(
									app.Button().
										Class("pf-c-select__menu-item").
										Aria("role", "option").
										Body(
											app.Text("Needs maintenance"),
										),
								),
						),
					app.Div().
						Class("pf-c-select__menu-footer").
						Body(
							app.Button().
								Class("pf-c-button pf-m-link pf-m-inline").
								Type("button").
								Body(
									app.Text("Action"),
								),
						),
				),
		)
}
