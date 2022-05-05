package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxSelect struct {
	app.Compo
}

func (c *CheckboxSelect) Render() app.UI {
	return app.Div().
		Class("pf-c-select").
		Body(
			app.Span().
				ID("select-checkbox-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-toggle").
				Aria("haspopup", true).
				Aria("expanded", "false").
				Aria("labelledby", "select-checkbox-label select-checkbox-toggle").
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
				Hidden(true).
				Body(
					app.FieldSet().
						Class("pf-c-select__menu-fieldset").
						Aria("label", "Select input").
						Body(
							app.Label().
								Class("pf-c-check pf-c-select__menu-item pf-m-description").
								For("select-checkbox-active").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-active").
										Name("select-checkbox-active"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Active"),
										),
									app.Span().
										Class("pf-c-check__description").
										Body(
											app.Text("This is a description"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item pf-m-description").
								For("select-checkbox-canceled").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-canceled").
										Name("select-checkbox-canceled"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Canceled"),
										),
									app.Span().
										Class("pf-c-check__description").
										Body(
											app.Text("This is a really long description that describes the menu item. This is a really long description that describes the menu item."),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-paused").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-paused").
										Name("select-checkbox-paused"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Paused"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-warning").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-warning").
										Name("select-checkbox-warning"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Warning"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-restarted").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-restarted").
										Name("select-checkbox-restarted"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Restarted"),
										),
								),
						),
				),
		)
}
