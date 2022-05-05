package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectMultiWithTypeahead struct {
	app.Compo
}

func (c *SelectMultiWithTypeahead) Render() app.UI {
	return app.Div().
		Class("pf-c-select").
		Body(
			app.Span().
				ID("select-multi-typeahead-label").
				Hidden(true).
				Body(
					app.Text("Choose states"),
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
								ID("select-multi-typeahead-typeahead").
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-button").
						Type("button").
						ID("select-multi-typeahead-toggle").
						Aria("haspopup", true).
						Aria("expanded", "false").
						Aria("labelledby", "select-multi-typeahead-label select-multi-typeahead-toggle").
						Aria("label", "Select").
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-multi-typeahead-label").
				Aria("role", "listbox").
				Hidden(true).
				Body(
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("Alabama"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("Florida"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("New\n        \u00a0Jersey"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("New\n        \u00a0York"),
								),
						),
					app.Li().
						Aria("role", "presentation").
						Body(
							app.Button().
								Class("pf-c-select__menu-item").
								Aria("role", "option").
								Body(
									app.Text("North Carolina"),
								),
						),
				),
		)
}
