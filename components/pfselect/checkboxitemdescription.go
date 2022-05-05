package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxItemDescription struct {
	app.Compo
}

func (c *CheckboxItemDescription) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-checkbox-description-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-description-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-checkbox-description-label select-checkbox-description-toggle").
				Body(
					app.Div().
						Class("pf-c-select__toggle-wrapper").
						Body(
							app.Span().
								Class("pf-c-select__toggle-text").
								Body(
									app.Text("Filter"),
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
					app.FieldSet().
						Class("pf-c-select__menu-fieldset").
						Aria("label", "Select input").
						Body(
							app.Label().
								Class("pf-c-check pf-c-select__menu-item pf-m-description").
								For("select-checkbox-description-active").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-description-active").
										Name("select-checkbox-description-active"),
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
								For("select-checkbox-description-canceled").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-description-canceled").
										Name("select-checkbox-description-canceled"),
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
								For("select-checkbox-description-paused").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-description-paused").
										Name("select-checkbox-description-paused"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Paused"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-description-warning").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-description-warning").
										Name("select-checkbox-description-warning"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Warning"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-description-restarted").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-description-restarted").
										Name("select-checkbox-description-restarted"),
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
