package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxWithCounts struct {
	app.Compo
}

func (c *CheckboxWithCounts) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-checkbox-counts-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-counts-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-checkbox-counts-label select-checkbox-counts-toggle").
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
								For("select-checkbox-counts-active").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-counts-active").
										Name("select-checkbox-counts-active").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-row").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-text").
														Body(
															app.Text("Active"),
														),
													app.Span().
														Class("pf-c-select__menu-item-count").
														Body(
															app.Text("3"),
														),
												),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-counts-canceled").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-counts-canceled").
										Name("select-checkbox-counts-canceled").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-row").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-text").
														Body(
															app.Text("Canceled"),
														),
													app.Span().
														Class("pf-c-select__menu-item-count").
														Body(
															app.Text("1"),
														),
												),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-counts-paused").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-counts-paused").
										Name("select-checkbox-counts-paused"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-row").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-text").
														Body(
															app.Text("Paused"),
														),
													app.Span().
														Class("pf-c-select__menu-item-count").
														Body(
															app.Text("15"),
														),
												),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-counts-warning").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-counts-warning").
										Name("select-checkbox-counts-warning").
										Checked(true),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-row").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-text").
														Body(
															app.Text("Warning"),
														),
													app.Span().
														Class("pf-c-select__menu-item-count").
														Body(
															app.Text("2"),
														),
												),
										),
								),
							app.Label().
								Class("pf-c-check pf-c-select__menu-item").
								For("select-checkbox-counts-restarted").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("select-checkbox-counts-restarted").
										Name("select-checkbox-counts-restarted"),
									app.Span().
										Class("pf-c-check__label").
										Body(
											app.Span().
												Class("pf-c-select__menu-item-row").
												Body(
													app.Span().
														Class("pf-c-select__menu-item-text").
														Body(
															app.Text("Restarted"),
														),
													app.Span().
														Class("pf-c-select__menu-item-count").
														Body(
															app.Text("8"),
														),
												),
										),
								),
						),
				),
		)
}
