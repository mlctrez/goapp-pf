package pfselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CheckboxExpandedAndSelectedWithGroupsAndFilter struct {
	app.Compo
}

func (c *CheckboxExpandedAndSelectedWithGroupsAndFilter) Render() app.UI {
	return app.Div().
		Class("pf-c-select pf-m-expanded").
		Body(
			app.Span().
				ID("select-checkbox-expanded-selected-filter-example-label").
				Hidden(true).
				Body(
					app.Text("Choose one"),
				),
			app.Button().
				Class("pf-c-select__toggle").
				Type("button").
				ID("select-checkbox-expanded-selected-filter-example-toggle").
				Aria("haspopup", true).
				Aria("expanded", true).
				Aria("labelledby", "select-checkbox-expanded-selected-filter-example-label select-checkbox-expanded-selected-filter-example-toggle").
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
				Body(
					app.Div().
						Class("pf-c-select__menu-search").
						Body(
							app.Div().
								Class("pf-c-search-input").
								Body(
									app.Div().
										Class("pf-c-search-input__bar").
										Body(
											app.Span().
												Class("pf-c-search-input__text").
												Body(
													app.Span().
														Class("pf-c-search-input__icon").
														Body(
															app.I().
																Class("fas fa-search fa-fw").
																Aria("hidden", true),
														),
													app.Input().
														Class("pf-c-search-input__text-input").
														Type("text").
														Placeholder("false").
														Aria("label", "Search"),
												),
										),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
					app.Div().
						Class("pf-c-select__menu-group").
						Body(
							app.Div().
								Class("pf-c-select__menu-group-title").
								Aria("hidden", true).
								ID("select-checkbox-expanded-selected-filter-example-group-status").
								Body(
									app.Text("Status"),
								),
							app.FieldSet().
								Class("pf-c-select__menu-fieldset").
								Aria("labelledby", "select-checkbox-expanded-selected-filter-example-group-status").
								Body(
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-running").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-running").
												Name("running"),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Running"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-stopped").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-stopped").
												Name("stopped").
												Checked(true),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Stopped"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-down").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-down").
												Name("down").
												Checked(true),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Down"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-degraded").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-degraded").
												Name("degraded"),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Degraded"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-needsMaintenance").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-needsMaintenance").
												Name("needsMaintenance").
												Checked(true),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Needs maintenance"),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-select__menu-group").
						Body(
							app.Div().
								Class("pf-c-select__menu-group-title").
								Aria("hidden", true).
								ID("select-checkbox-expanded-selected-filter-example-group-vendor").
								Body(
									app.Text("Vendor"),
								),
							app.FieldSet().
								Class("pf-c-select__menu-fieldset").
								Aria("labelledby", "select-checkbox-expanded-selected-filter-example-group-vendor").
								Body(
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-dell").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-dell").
												Name("dell"),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Dell"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item pf-m-disabled").
										For("select-checkbox-expanded-selected-filter-example-samsung").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-samsung").
												Name("samsung").
												Disabled(true),
											app.Span().
												Class("pf-c-check__label pf-m-disabled").
												Body(
													app.Text("Samsung"),
												),
										),
									app.Label().
										Class("pf-c-check pf-c-select__menu-item").
										For("select-checkbox-expanded-selected-filter-example-hp").
										Body(
											app.Input().
												Class("pf-c-check__input").
												Type("checkbox").
												ID("select-checkbox-expanded-selected-filter-example-hp").
												Name("hp"),
											app.Span().
												Class("pf-c-check__label").
												Body(
													app.Text("Hewlett-Packard"),
												),
										),
								),
						),
				),
		)
}
