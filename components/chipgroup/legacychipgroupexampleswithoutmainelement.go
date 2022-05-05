package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LegacyChipGroupExamplesWithoutMainElement struct {
	app.Compo
}

func (c *LegacyChipGroupExamplesWithoutMainElement) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-chip-group").
				Body(
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("label", "Chip group list").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simplechip_one_select_collapsed").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simpleremove_chip_one_select_collapsed legacy-simplechip_one_select_collapsed").
												Aria("label", "Remove").
												ID("legacy-simpleremove_chip_one_select_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simplechip_two_select_collapsed").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simpleremove_chip_two_select_collapsed legacy-simplechip_two_select_collapsed").
												Aria("label", "Remove").
												ID("legacy-simpleremove_chip_two_select_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simplechip_three_select_collapsed").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simpleremove_chip_three_select_collapsed legacy-simplechip_three_select_collapsed").
												Aria("label", "Remove").
												ID("legacy-simpleremove_chip_three_select_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip-group").
				Body(
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("label", "Chip group list").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_one_toolbar").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_one_toolbar legacy-simple-removablechip_one_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_one_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_two_toolbar").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_two_toolbar legacy-simple-removablechip_two_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_two_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_three_toolbar").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_three_toolbar legacy-simple-removablechip_three_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_three_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_four_toolbar").
												Body(
													app.Text("Chip four"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_four_toolbar legacy-simple-removablechip_four_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_four_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_five_toolbar").
												Body(
													app.Text("Chip five"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_five_toolbar legacy-simple-removablechip_five_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_five_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-simple-removablechip_six_toolbar").
												Body(
													app.Text("Chip six"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-simple-removableremove_chip_six_toolbar legacy-simple-removablechip_six_toolbar").
												Aria("label", "Remove").
												ID("legacy-simple-removableremove_chip_six_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-chip-group__close").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("labelledby", "legacy-simple-removable-button legacy-simple-removable-label").
								Aria("label", "Close chip group").
								ID("legacy-simple-removable-button").
								Body(
									app.I().
										Class("fas fa-times-circle").
										Aria("hidden", true),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip-group pf-m-category").
				Body(
					app.Span().
						Class("pf-c-chip-group__label").
						Aria("hidden", true).
						ID("legacy-category-label").
						Body(
							app.Text("Category one"),
						),
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("labelledby", "legacy-category-label").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-categorychip_one_toolbar_collapsed").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-categoryremove_chip_one_toolbar_collapsed legacy-categorychip_one_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-categoryremove_chip_one_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-categorychip_two_toolbar_collapsed").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-categoryremove_chip_two_toolbar_collapsed legacy-categorychip_two_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-categoryremove_chip_two_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-categorychip_three_toolbar_collapsed").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-categoryremove_chip_three_toolbar_collapsed legacy-categorychip_three_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-categoryremove_chip_three_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-chip-group pf-m-category").
				Body(
					app.Span().
						Class("pf-c-chip-group__label").
						Aria("hidden", true).
						ID("legacy-category-removable-label").
						Body(
							app.Text("Category one"),
						),
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("labelledby", "legacy-category-removable-label").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-category-removablechip_one_toolbar_collapsed").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-category-removableremove_chip_one_toolbar_collapsed legacy-category-removablechip_one_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-category-removableremove_chip_one_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-category-removablechip_two_toolbar_collapsed").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-category-removableremove_chip_two_toolbar_collapsed legacy-category-removablechip_two_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-category-removableremove_chip_two_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("legacy-category-removablechip_three_toolbar_collapsed").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "legacy-category-removableremove_chip_three_toolbar_collapsed legacy-category-removablechip_three_toolbar_collapsed").
												Aria("label", "Remove").
												ID("legacy-category-removableremove_chip_three_toolbar_collapsed").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-chip-group__close").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("labelledby", "legacy-category-removable-button legacy-category-removable-label").
								Aria("label", "Close chip group").
								ID("legacy-category-removable-button").
								Body(
									app.I().
										Class("fas fa-times-circle").
										Aria("hidden", true),
								),
						),
				),
		)
}
