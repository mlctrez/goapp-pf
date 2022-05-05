package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SingleWithTypeaheadExpanded struct {
	app.Compo
}

func (c *SingleWithTypeaheadExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-single-typeahead-expanded-label").
				Hidden(true).
				Body(
					app.Text("New"),
				),
			app.Div().
				Class("pf-c-select__toggle pf-m-typeahead").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Input().
								Class("pf-c-form-control pf-c-select__toggle-typeahead").
								Type("text").
								ID("select-single-typeahead-expanded-typeahead").
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-clear").
						Type("button").
						Aria("label", "Clear all").
						Body(
							app.I().
								Class("fas fa-times-circle").
								Aria("hidden", true),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-button").
						Type("button").
						ID("select-single-typeahead-expanded-toggle").
						Aria("haspopup", true).
						Aria("expanded", true).
						Aria("labelledby", "select-single-typeahead-expanded-label select-single-typeahead-expanded-toggle").
						Aria("label", "Select").
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-single-typeahead-expanded-label").
				Aria("role", "listbox").
				Body(
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Mark().
										Class("pf-c-select__menu-item--match").
										Body(
											app.Text("New"),
										),
									app.Text("Jersey"),
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
									app.Mark().
										Class("pf-c-select__menu-item--match").
										Body(
											app.Text("New"),
										),
									app.Text("Mexico"),
									app.Span().
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
									app.Mark().
										Class("pf-c-select__menu-item--match").
										Body(
											app.Text("New"),
										),
									app.Text("York"),
								),
						),
				),
		)
}
