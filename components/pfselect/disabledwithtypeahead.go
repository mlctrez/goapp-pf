package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DisabledWithTypeahead struct {
	app.Compo
}

func (c *DisabledWithTypeahead) Render() app.UI {
	return app.Div().
		Class("pf-c-select").
		Body(
			app.Span().
				ID("select-single-typeahead-disabled-label").
				Hidden(true).
				Body(
					app.Text("Choose a state"),
				),
			app.Div().
				Class("pf-c-select__toggle pf-m-typeahead pf-m-disabled").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Input().
								Class("pf-c-form-control pf-c-select__toggle-typeahead").
								Type("text").
								ID("select-single-typeahead-disabled-typeahead").
								Aria("label", "Type to filter").
								Placeholder("false").
								Disabled(true),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-button").
						Type("button").
						ID("select-single-typeahead-disabled-toggle").
						Aria("haspopup", true).
						Aria("expanded", "false").
						Aria("labelledby", "select-single-typeahead-disabled-label select-single-typeahead-disabled-toggle").
						Aria("label", "Select").
						Disabled(true).
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-single-typeahead-disabled-label").
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
								Class("pf-c-select__menu-item pf-m-selected").
								Aria("role", "option").
								Aria("selected", true).
								Body(
									app.Text("New\n        \u00a0Mexico"), app.Span().
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
