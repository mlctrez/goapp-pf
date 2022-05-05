package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LoadingMenuItem struct {
	app.Compo
}

func (c *LoadingMenuItem) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-single-loading-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-single-loading-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-single-loading-label select-single-loading-toggle").
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
			app.Ul().
				Class("pf-c-select__menu").
				Aria("role", "listbox").
				Aria("labelledby", "select-single-loading-label").
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
					app.Li().
						Aria("role", "presentation").
						Class("pf-c-select__list-item pf-m-loading").
						Body(
							app.Span().
								Class("pf-c-spinner pf-m-lg").
								Aria("role", "progressbar").
								Aria("label", "Loading items").
								Body(
									app.Span().
										Class("pf-c-spinner__clipper"),
									app.Span().
										Class("pf-c-spinner__lead-ball"),
									app.Span().
										Class("pf-c-spinner__tail-ball"),
								),
						),
				),
		)
}
