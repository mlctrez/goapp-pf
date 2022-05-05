package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxExpandedWithoutBadge struct {
	app.Compo
}

func (c *CheckboxExpandedWithoutBadge) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-checkbox-without-badge-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-without-badge-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-checkbox-without-badge-label select-checkbox-without-badge-toggle").
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
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-without-badge-active").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-without-badge-active").
										Name("select-checkbox-without-badge-active"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Active"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-without-badge-canceled").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-without-badge-canceled").
										Name("select-checkbox-without-badge-canceled").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Canceled"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-without-badge-paused").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-without-badge-paused").
										Name("select-checkbox-without-badge-paused").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Paused"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-without-badge-warning").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-without-badge-warning").
										Name("select-checkbox-without-badge-warning"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Warning"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-without-badge-restarted").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-without-badge-restarted").
										Name("select-checkbox-without-badge-restarted").
										Checked(true),
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
