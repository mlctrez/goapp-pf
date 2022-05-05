package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectableExpandableRows struct {
	app.Compo
}

func (c *SelectableExpandableRows) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Selectable, expandable data list example").
		ID("data-list-selectable-expandable-rows").
		Body(
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded pf-m-selectable").
				Aria("labelledby", "data-list-selectable-expandable-rows-item-1").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-control").
								Body(
									app.Div().
										Class("pf-c-data-list__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "data-list-selectable-expandable-rows-toggle1 data-list-selectable-expandable-rows-item1").
												ID("data-list-selectable-expandable-rows-toggle1").
												Aria("label", "Toggle details for").
												Aria("expanded", true).
												Aria("controls", "data-list-selectable-expandable-rows-content1").
												Body(
													app.Div().
														Class("pf-c-data-list__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-expandable-rows-item-1").
												Body(
													app.Text("Primary content (selected, expanded)"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-data-list__expandable-content").
						ID("data-list-selectable-expandable-rows-content1").
						Aria("label", "Primary content details").
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body").
								Body(
									app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-selectable").
				Aria("labelledby", "data-list-selectable-expandable-rows-item-2").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-control").
								Body(
									app.Div().
										Class("pf-c-data-list__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "data-list-selectable-expandable-rows-toggle2 data-list-selectable-expandable-rows-item2").
												ID("data-list-selectable-expandable-rows-toggle2").
												Aria("label", "Toggle details for").
												Aria("expanded", "false").
												Aria("controls", "data-list-selectable-expandable-rows-content2").
												Body(
													app.Div().
														Class("pf-c-data-list__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-expandable-rows-item-2").
												Body(
													app.Text("Secondary content"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-data-list__expandable-content").
						ID("data-list-selectable-expandable-rows-content2").
						Aria("label", "Secondary content details").
						Hidden(true).
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body").
								Body(
									app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded pf-m-selectable").
				Aria("labelledby", "data-list-selectable-expandable-rows-item-3").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-control").
								Body(
									app.Div().
										Class("pf-c-data-list__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "data-list-selectable-expandable-rows-toggle3 data-list-selectable-expandable-rows-item3").
												ID("data-list-selectable-expandable-rows-toggle3").
												Aria("label", "Toggle details for").
												Aria("expanded", true).
												Aria("controls", "data-list-selectable-expandable-rows-content3").
												Body(
													app.Div().
														Class("pf-c-data-list__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-expandable-rows-item-3").
												Body(
													app.Text("Tertiary content (not selected, expanded)"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-data-list__expandable-content").
						ID("data-list-selectable-expandable-rows-content3").
						Aria("label", "Tertiary content details").
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body pf-m-no-padding").
								Body(
									app.Text("This expanded section has no padding."),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-selectable").
				Aria("labelledby", "data-list-selectable-expandable-rows-item-4").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-control").
								Body(
									app.Div().
										Class("pf-c-data-list__toggle").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "data-list-selectable-expandable-rows-toggle4 data-list-selectable-expandable-rows-item4").
												ID("data-list-selectable-expandable-rows-toggle4").
												Aria("label", "Toggle details for").
												Aria("expanded", "false").
												Aria("controls", "data-list-selectable-expandable-rows-content4").
												Body(
													app.Div().
														Class("pf-c-data-list__toggle-icon").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-expandable-rows-item-4").
												Body(
													app.Text("Quaternary content (selected)"),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-data-list__expandable-content").
						ID("data-list-selectable-expandable-rows-content4").
						Aria("label", "Quaternary content details").
						Hidden(true).
						Body(
							app.Div().
								Class("pf-c-data-list__expandable-content-body").
								Body(
									app.Text("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
								),
						),
				),
		)
}
