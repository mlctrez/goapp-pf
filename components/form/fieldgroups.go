package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FieldGroups struct {
	app.Compo
}

func (c *FieldGroups) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form").
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-expandable-field-groups-label1").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Label 1"),
										),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Body(
											app.Text("*"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for label 1 field").
								Aria("describedby", "form-expandable-field-groups-label1").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-expandable-field-groups-label1").
								Name("form-expandable-field-groups-label1").
								Required(true),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-expandable-field-groups-label2").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Label 2"),
										),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Body(
											app.Text("*"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for label 2 field").
								Aria("describedby", "form-expandable-field-groups-label2").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-expandable-field-groups-label2").
								Name("form-expandable-field-groups-label2").
								Required(true),
						),
				),
			app.Div().
				Class("pf-c-form__field-group pf-m-expanded").
				Aria("role", "group").
				Aria("labelledby", "form-expandable-field-groups-field-group1-title").
				Body(
					app.Div().
						Class("pf-c-form__field-group-toggle").
						Body(
							app.Div().
								Class("pf-c-form__field-group-toggle-button").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("expanded", true).
										Aria("label", "Details").
										ID("form-expandable-field-groups-field-group1-toggle").
										Aria("labelledby", "form-expandable-field-groups-field-group1-title form-expandable-field-groups-field-group1-toggle").
										Body(
											app.Span().
												Class("pf-c-form__field-group-toggle-icon").
												Body(
													app.I().
														Class("fas fa-angle-right").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-header").
						Body(
							app.Div().
								Class("pf-c-form__field-group-header-main").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header-title").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-title-text").
												ID("form-expandable-field-groups-field-group1-title").
												Body(
													app.Text("Field group 1"),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header-description").
										Body(
											app.Text("Field group 1 description text."),
										),
								),
							app.Div().
								Class("pf-c-form__field-group-header-actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Body(
											app.Text("Delete all"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Add parameter"),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-body").
						Body(
							app.Div().
								Class("pf-c-form__field-group pf-m-expanded").
								Aria("role", "group").
								Aria("labelledby", "form-expandable-field-groups-field-group2-title").
								Body(
									app.Div().
										Class("pf-c-form__field-group-toggle").
										Body(
											app.Div().
												Class("pf-c-form__field-group-toggle-button").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("expanded", true).
														Aria("label", "Details").
														ID("form-expandable-field-groups-field-group2-toggle").
														Aria("labelledby", "form-expandable-field-groups-field-group2-title form-expandable-field-groups-field-group2-toggle").
														Body(
															app.Span().
																Class("pf-c-form__field-group-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-main").
												Body(
													app.Div().
														Class("pf-c-form__field-group-header-title").
														Body(
															app.Div().
																Class("pf-c-form__field-group-header-title-text").
																ID("form-expandable-field-groups-field-group2-title").
																Body(
																	app.Text("Nested field group 1"),
																),
														),
													app.Div().
														Class("pf-c-form__field-group-header-description").
														Body(
															app.Text("Nested field group 1 description text."),
														),
												),
											app.Div().
												Class("pf-c-form__field-group-header-actions").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Remove").
														Body(
															app.I().
																Class("fas fa-trash"),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-body").
										Body(
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group2-label1").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 1"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 1 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group2-label1").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group2-label1").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group2-label1").
																Required(true),
														),
												),
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group2-label2").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 2"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 2 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group2-label2").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group2-label2").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group2-label2").
																Required(true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-form__field-group").
								Aria("role", "group").
								Aria("labelledby", "form-expandable-field-groups-field-group3-title").
								Body(
									app.Div().
										Class("pf-c-form__field-group-toggle").
										Body(
											app.Div().
												Class("pf-c-form__field-group-toggle-button").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("expanded", "false").
														Aria("label", "Details").
														ID("form-expandable-field-groups-field-group3-toggle").
														Aria("labelledby", "form-expandable-field-groups-field-group3-title form-expandable-field-groups-field-group3-toggle").
														Body(
															app.Span().
																Class("pf-c-form__field-group-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-main").
												Body(
													app.Div().
														Class("pf-c-form__field-group-header-title").
														Body(
															app.Div().
																Class("pf-c-form__field-group-header-title-text").
																ID("form-expandable-field-groups-field-group3-title").
																Body(
																	app.Text("Nested field group 2"),
																),
														),
												),
											app.Div().
												Class("pf-c-form__field-group-header-actions").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Remove").
														Body(
															app.I().
																Class("fas fa-trash"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-form__field-group").
								Aria("role", "group").
								Aria("labelledby", "form-expandable-field-groups-field-group4-title").
								Body(
									app.Div().
										Class("pf-c-form__field-group-toggle").
										Body(
											app.Div().
												Class("pf-c-form__field-group-toggle-button").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("expanded", "false").
														Aria("label", "Details").
														ID("form-expandable-field-groups-field-group4-toggle").
														Aria("labelledby", "form-expandable-field-groups-field-group4-title form-expandable-field-groups-field-group4-toggle").
														Body(
															app.Span().
																Class("pf-c-form__field-group-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-main").
												Body(
													app.Div().
														Class("pf-c-form__field-group-header-title").
														Body(
															app.Div().
																Class("pf-c-form__field-group-header-title-text").
																ID("form-expandable-field-groups-field-group4-title").
																Body(
																	app.Text("Nested field group 3"),
																),
														),
													app.Div().
														Class("pf-c-form__field-group-header-description").
														Body(
															app.Text("Nested field group 3 description text."),
														),
												),
											app.Div().
												Class("pf-c-form__field-group-header-actions").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Remove").
														Body(
															app.I().
																Class("fas fa-trash"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group1-label1").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 1"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 1 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group1-label1").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group1-label1").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group1-label1").
												Required(true),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group1-label2").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 2"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 2 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group1-label2").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group1-label2").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group1-label2").
												Required(true),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-form__field-group").
				Aria("role", "group").
				Aria("labelledby", "form-expandable-field-groups-field-group5-title").
				Body(
					app.Div().
						Class("pf-c-form__field-group-toggle").
						Body(
							app.Div().
								Class("pf-c-form__field-group-toggle-button").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("expanded", "false").
										Aria("label", "Details").
										ID("form-expandable-field-groups-field-group5-toggle").
										Aria("labelledby", "form-expandable-field-groups-field-group5-title form-expandable-field-groups-field-group5-toggle").
										Body(
											app.Span().
												Class("pf-c-form__field-group-toggle-icon").
												Body(
													app.I().
														Class("fas fa-angle-right").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-header").
						Body(
							app.Div().
								Class("pf-c-form__field-group-header-main").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header-title").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-title-text").
												ID("form-expandable-field-groups-field-group5-title").
												Body(
													app.Text("Field group 2"),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header-description").
										Body(
											app.Text("Field group 1 description text."),
										),
								),
							app.Div().
								Class("pf-c-form__field-group-header-actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Body(
											app.Text("Delete all"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Add parameter"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-form__field-group pf-m-expanded").
				Aria("role", "group").
				Aria("labelledby", "form-expandable-field-groups-field-group6-title").
				Body(
					app.Div().
						Class("pf-c-form__field-group-toggle").
						Body(
							app.Div().
								Class("pf-c-form__field-group-toggle-button").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("expanded", true).
										Aria("label", "Details").
										ID("form-expandable-field-groups-field-group6-toggle").
										Aria("labelledby", "form-expandable-field-groups-field-group6-title form-expandable-field-groups-field-group6-toggle").
										Body(
											app.Span().
												Class("pf-c-form__field-group-toggle-icon").
												Body(
													app.I().
														Class("fas fa-angle-right").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-header").
						Body(
							app.Div().
								Class("pf-c-form__field-group-header-main").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header-title").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-title-text").
												ID("form-expandable-field-groups-field-group6-title").
												Body(
													app.Text("Field group 3"),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header-description").
										Body(
											app.Text("Field group 1 description text."),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-body").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group6-label1").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 1"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 1 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group6-label1").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group6-label1").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group6-label1").
												Required(true),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group6-label2").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 2"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 2 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group6-label2").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group6-label2").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group6-label2").
												Required(true),
										),
								),
							app.Div().
								Class("pf-c-form__field-group pf-m-expanded").
								Aria("role", "group").
								Aria("labelledby", "form-expandable-field-groups-field-group7-title").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-main").
												Body(
													app.Div().
														Class("pf-c-form__field-group-header-title").
														Body(
															app.Div().
																Class("pf-c-form__field-group-header-title-text").
																ID("form-expandable-field-groups-field-group7-title").
																Body(
																	app.Text("Nested field group 1 (non-expandable)"),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-body").
										Body(
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group7-label1").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 1"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 1 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group7-label1").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group7-label1").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group7-label1").
																Required(true),
														),
												),
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group7-label2").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 2"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 2 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group7-label2").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group7-label2").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group7-label2").
																Required(true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-form__field-group pf-m-expanded").
								Aria("role", "group").
								Aria("labelledby", "form-expandable-field-groups-field-group8-title").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-main").
												Body(
													app.Div().
														Class("pf-c-form__field-group-header-title").
														Body(
															app.Div().
																Class("pf-c-form__field-group-header-title-text").
																ID("form-expandable-field-groups-field-group8-title").
																Body(
																	app.Text("Nested field group 2 (non-expandable)"),
																),
														),
													app.Div().
														Class("pf-c-form__field-group-header-description").
														Body(
															app.Text("Field group 1 description text."),
														),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-body").
										Body(
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group8-label1").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 1"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 1 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group8-label1").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group8-label1").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group8-label1").
																Required(true),
														),
												),
											app.Div().
												Class("pf-c-form__group").
												Body(
													app.Div().
														Class("pf-c-form__group-label").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("form-expandable-field-groupsform-expandable-field-groups-field-group8-label2").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Label 2"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Button().
																Class("pf-c-form__group-label-help").
																Aria("label", "More information for label 2 field").
																Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group8-label2").
																Body(
																	app.I().
																		Class("pficon pf-icon-help").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-form__group-control").
														Body(
															app.Input().
																Class("pf-c-form-control").
																Type("text").
																ID("form-expandable-field-groupsform-expandable-field-groups-field-group8-label2").
																Name("form-expandable-field-groupsform-expandable-field-groups-field-group8-label2").
																Required(true),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-form__field-group").
				Aria("role", "group").
				Aria("labelledby", "form-expandable-field-groups-field-group9-title").
				Body(
					app.Div().
						Class("pf-c-form__field-group-header").
						Body(
							app.Div().
								Class("pf-c-form__field-group-header-main").
								Body(
									app.Div().
										Class("pf-c-form__field-group-header-title").
										Body(
											app.Div().
												Class("pf-c-form__field-group-header-title-text").
												ID("form-expandable-field-groups-field-group9-title").
												Body(
													app.Text("Field group 4 (non-expandable)"),
												),
										),
									app.Div().
										Class("pf-c-form__field-group-header-description").
										Body(
											app.Text("Field group 1 description text."),
										),
								),
							app.Div().
								Class("pf-c-form__field-group-header-actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Body(
											app.Text("Delete all"),
										),
									app.Button().
										Class("pf-c-button pf-m-secondary").
										Type("button").
										Body(
											app.Text("Add parameter"),
										),
								),
						),
					app.Div().
						Class("pf-c-form__field-group-body").
						Hidden(true).
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group9-label1").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 1"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 1 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group9-label1").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group9-label1").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group9-label1").
												Required(true),
										),
								),
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("form-expandable-field-groupsform-expandable-field-groups-field-group9-label2").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Body(
															app.Text("Label 2"),
														),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Body(
															app.Text("*"),
														),
												),
											app.Button().
												Class("pf-c-form__group-label-help").
												Aria("label", "More information for label 2 field").
												Aria("describedby", "form-expandable-field-groupsform-expandable-field-groups-field-group9-label2").
												Body(
													app.I().
														Class("pficon pf-icon-help").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("text").
												ID("form-expandable-field-groupsform-expandable-field-groups-field-group9-label2").
												Name("form-expandable-field-groupsform-expandable-field-groups-field-group9-label2").
												Required(true),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-expandable-field-groups-label3").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Label 3"),
										),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Body(
											app.Text("*"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for label 3 field").
								Aria("describedby", "form-expandable-field-groups-label3").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-expandable-field-groups-label3").
								Name("form-expandable-field-groups-label3").
								Required(true),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-expandable-field-groups-label4").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Label 4"),
										),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Body(
											app.Text("*"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for label 4 field").
								Aria("describedby", "form-expandable-field-groups-label4").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-expandable-field-groups-label4").
								Name("form-expandable-field-groups-label4").
								Required(true),
						),
				),
		)
}
