package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Invalid struct {
	app.Compo
}

func (c *Invalid) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-invalid").
		Body(
			app.Span().
				ID("select-invalid-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-invalid-toggle").
				Aria("haspopup", true).
				Aria("expanded", "false").
				Aria("labelledby", "select-invalid-label select-invalid-toggle").
				Aria("invalid", true).
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
						Class("pf-c-select__toggle-status-icon").
						Body(
							app.I().
								Class("fas fa-exclamation-circle").
								Aria("hidden", true),
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
				Aria("labelledby", "select-invalid-label").
				Hidden(true).
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
		)
}
