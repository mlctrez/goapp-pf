package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InvalidWithTypeahead struct {
	app.Compo
}

func (c *InvalidWithTypeahead) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-invalid").
		Body(
			app.Span().
				ID("select-single-typeahead-invalid-label").
				Hidden(true).
				Body(
					app.Text("Choose a state"),
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
								ID("select-single-typeahead-invalid-typeahead").
								Aria("invalid", true).
								Value("Invalid").
								Aria("label", "Type to filter").
								Placeholder("false"),
						),
					app.Span().
						Class("pf-c-select__toggle-status-icon").
						Body(
							app.I().
								Class("fas fa-exclamation-circle").
								Aria("hidden", true),
						),
					app.Button().
						Class("pf-c-button pf-m-plain pf-c-select__toggle-button").
						Type("button").
						ID("select-single-typeahead-invalid-toggle").
						Aria("haspopup", true).
						Aria("expanded", "false").
						Aria("labelledby", "select-single-typeahead-invalid-label select-single-typeahead-invalid-toggle").
						Aria("label", "Select").
						Body(
							app.I().
								Class("fas fa-caret-down pf-c-select__toggle-arrow").
								Aria("hidden", true),
						),
				),
			app.Ul().
				Class("pf-c-select__menu").
				Aria("labelledby", "select-single-typeahead-invalid-label").
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
