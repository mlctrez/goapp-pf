package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineEditTableRow struct {
	app.Compo
}

func (c *InlineEditTableRow) Render() app.UI {
	return app.Form().
		Class("pf-c-inline-edit").
		ID("bulk-edit-table-example").
		Body(
			app.Table().
				Class("pf-c-table pf-m-grid-lg").
				Aria("role", "grid").
				Aria("label", "Inline edit table row example").
				ID("inline-edit-table-row-example").
				Body(
					app.Caption().
						Body(
							app.Text("This is the table caption"),
						),
					app.THead().
						Body(
							app.Tr().
								Aria("role", "row").
								Body(
									app.Th().
										Aria("role", "columnheader").
										Body(
											app.Text("Text input"),
										),
									app.Th().
										Aria("role", "columnheader").
										Body(
											app.Text("Disabled text input"),
										),
									app.Th().
										Aria("role", "columnheader").
										Body(
											app.Text("Checkboxes"),
										),
									app.Th().
										Aria("role", "columnheader").
										Body(
											app.Text("Radios"),
										),
									app.Th().
										Aria("role", "columnheader").
										Body(
											app.Text("Number"),
										),
									app.Td(),
									app.Td(),
								),
						),
					app.TBody().
						Aria("role", "rowgroup").
						Body(
							app.Tr().
								Class("pf-m-inline-editable").
								Aria("role", "row").
								Body(
									app.Th().
										Aria("role", "columnheader").
										DataSet("label", "Text input").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Text input description content"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("Text input description content").
														ID("bulk-edit-table-example-row-1-text-input").
														Aria("label", "Text input"),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Disabled text input").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Text input disabled, description content"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("Text input disabled, description content").
														ID("bulk-edit-table-example-row-1-text-input-disabled").
														Aria("label", "Disabled text input").
														Disabled(true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Checkboxes").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Check 1, Check 2"),
												),
											app.Div().
												Class("pf-c-inline-edit__group pf-m-column").
												Body(
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-check").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("bulk-edit-table-example-row-1-check-1").
																		Name("bulk-edit-table-example-row-1-example-check"),
																	app.Label().
																		Class("pf-c-check__label").
																		For("bulk-edit-table-example-row-1-check-1").
																		Body(
																			app.Text("Check 1"),
																		),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-check").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("bulk-edit-table-example-row-1-check-2").
																		Name("bulk-edit-table-example-row-1-example-check-2"),
																	app.Label().
																		Class("pf-c-check__label").
																		For("bulk-edit-table-example-row-1-check-2").
																		Body(
																			app.Text("Check 2"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Radios").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Radio 1, Radio 2"),
												),
											app.Div().
												Class("pf-c-inline-edit__group pf-m-column").
												Aria("role", "radiogroup").
												Aria("label", "Radio group example").
												Body(
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-radio").
																Body(
																	app.Input().
																		Class("pf-c-radio__input").
																		Type("radio").
																		ID("bulk-edit-table-example-row-1-radio-1").
																		Name("bulk-edit-table-example-row-1-example-radio"),
																	app.Label().
																		Class("pf-c-radio__label").
																		For("bulk-edit-table-example-row-1-radio-1").
																		Body(
																			app.Text("Radio 1"),
																		),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-radio").
																Body(
																	app.Input().
																		Class("pf-c-radio__input").
																		Type("radio").
																		ID("bulk-edit-table-example-row-1-radio-2").
																		Name("bulk-edit-table-example-row-1-example-radio"),
																	app.Label().
																		Class("pf-c-radio__label").
																		For("bulk-edit-table-example-row-1-radio-2").
																		Body(
																			app.Text("Radio 2"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Number").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("2"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("number").
														Value("2").
														ID("bulk-edit-table-example-row-1-number-input").
														Aria("label", "Number input"),
												),
										),
									app.Td().
										Class("pf-c-table__inline-edit-action").
										Aria("role", "cell").
										Body(
											app.Div().
												Class("pf-c-inline-edit__group pf-m-action-group pf-m-icon-group").
												Body(
													app.Div().
														Class("pf-c-inline-edit__action pf-m-valid").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Save edits").
																Body(
																	app.I().
																		Class("fas fa-check").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__action").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Cancel edits").
																Body(
																	app.I().
																		Class("fas fa-times").
																		Aria("hidden", true),
																),
														),
												),
											app.Div().
												Class("pf-c-inline-edit__action pf-m-enable-editable").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														ID("bulk-edit-table-example-row-1-edit-button").
														Aria("label", "Edit").
														Aria("labelledby", "bulk-edit-table-example-label bulk-edit-table-example-row-1-edit-button").
														Body(
															app.I().
																Class("fas fa-pencil-alt").
																Aria("hidden", true),
														),
												),
										),
									app.Td().
										Class("pf-c-table__action").
										Aria("role", "cell").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("inline-edit-table-row-example-row-1--dropdown-kebab-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "inline-edit-table-row-example-row-1--dropdown-kebab-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
																		Href("#").
																		Aria("disabled", true).
																		TabIndex(-1).
																		Body(
																			app.Text("Disabled link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
								),
							app.Tr().
								Aria("role", "row").
								Body(
									app.Th().
										Aria("role", "columnheader").
										DataSet("label", "Text input").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Text input description content"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("Text input description content").
														ID("bulk-edit-table-example-row-2-text-input").
														Aria("label", "Text input"),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Disabled text input").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Text input disabled, description content"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("Text input disabled, description content").
														ID("bulk-edit-table-example-row-2-text-input-disabled").
														Aria("label", "Disabled text input").
														Disabled(true),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Checkboxes").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Check 1, Check 2"),
												),
											app.Div().
												Class("pf-c-inline-edit__group pf-m-column").
												Body(
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-check").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("bulk-edit-table-example-row-2-check-1").
																		Name("bulk-edit-table-example-row-2-example-check"),
																	app.Label().
																		Class("pf-c-check__label").
																		For("bulk-edit-table-example-row-2-check-1").
																		Body(
																			app.Text("Check 1"),
																		),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-check").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("bulk-edit-table-example-row-2-check-2").
																		Name("bulk-edit-table-example-row-2-example-check-2"),
																	app.Label().
																		Class("pf-c-check__label").
																		For("bulk-edit-table-example-row-2-check-2").
																		Body(
																			app.Text("Check 2"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Radios").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("Radio 1, Radio 2"),
												),
											app.Div().
												Class("pf-c-inline-edit__group pf-m-column").
												Aria("role", "radiogroup").
												Aria("label", "Radio group example").
												Body(
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-radio").
																Body(
																	app.Input().
																		Class("pf-c-radio__input").
																		Type("radio").
																		ID("bulk-edit-table-example-row-2-radio-1").
																		Name("bulk-edit-table-example-row-2-example-radio-1"),
																	app.Label().
																		Class("pf-c-radio__label").
																		For("bulk-edit-table-example-row-2-radio-1").
																		Body(
																			app.Text("Radio 1"),
																		),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__input").
														Body(
															app.Div().
																Class("pf-c-radio").
																Body(
																	app.Input().
																		Class("pf-c-radio__input").
																		Type("radio").
																		ID("bulk-edit-table-example-row-2-radio-2").
																		Name("bulk-edit-table-example-row-2-example-radio-2"),
																	app.Label().
																		Class("pf-c-radio__label").
																		For("bulk-edit-table-example-row-2-radio-2").
																		Body(
																			app.Text("Radio 2"),
																		),
																),
														),
												),
										),
									app.Td().
										Aria("role", "cell").
										DataSet("label", "Number").
										Body(
											app.Div().
												Class("pf-c-inline-edit__value").
												Body(
													app.Text("2"),
												),
											app.Div().
												Class("pf-c-inline-edit__input").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("number").
														Value("2").
														ID("bulk-edit-table-example-row-2-number-input").
														Aria("label", "Number input"),
												),
										),
									app.Td().
										Class("pf-c-table__inline-edit-action").
										Aria("role", "cell").
										Body(
											app.Div().
												Class("pf-c-inline-edit__group pf-m-action-group pf-m-icon-group").
												Body(
													app.Div().
														Class("pf-c-inline-edit__action pf-m-valid").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Save edits").
																Body(
																	app.I().
																		Class("fas fa-check").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-inline-edit__action").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Cancel edits").
																Body(
																	app.I().
																		Class("fas fa-times").
																		Aria("hidden", true),
																),
														),
												),
											app.Div().
												Class("pf-c-inline-edit__action pf-m-enable-editable").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														ID("bulk-edit-table-example-row-2-edit-button").
														Aria("label", "Edit").
														Aria("labelledby", "bulk-edit-table-example-label bulk-edit-table-example-row-2-edit-button").
														Body(
															app.I().
																Class("fas fa-pencil-alt").
																Aria("hidden", true),
														),
												),
										),
									app.Td().
										Class("pf-c-table__action").
										Aria("role", "cell").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("inline-edit-table-row-example-row-2--dropdown-kebab-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "inline-edit-table-row-example-row-2--dropdown-kebab-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
																		Href("#").
																		Aria("disabled", true).
																		TabIndex(-1).
																		Body(
																			app.Text("Disabled link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
								),
						),
				),
		)
}
