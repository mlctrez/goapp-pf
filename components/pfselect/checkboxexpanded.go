package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxExpanded struct {
	app.Compo
}

func (c *CheckboxExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-checkbox-expanded-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-expanded-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-checkbox-expanded-label select-checkbox-expanded-toggle").
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
								For("select-checkbox-expanded-active").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-expanded-active").
										Name("select-checkbox-expanded-active"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Active"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-expanded-canceled").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-expanded-canceled").
										Name("select-checkbox-expanded-canceled").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Canceled"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-expanded-paused").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-expanded-paused").
										Name("select-checkbox-expanded-paused").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Paused"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-expanded-warning").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-expanded-warning").
										Name("select-checkbox-expanded-warning"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Text("Warning"),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-expanded-restarted").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-expanded-restarted").
										Name("select-checkbox-expanded-restarted").
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
