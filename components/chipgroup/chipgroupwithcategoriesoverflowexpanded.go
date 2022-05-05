package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ChipGroupWithCategoriesOverflowExpanded struct {
	app.Compo
}

func (c *ChipGroupWithCategoriesOverflowExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-chip-group pf-m-category").
		Body(
			app.Div().
				Class("pf-c-chip-group__main").
				Body(
					app.Span().
						Class("pf-c-chip-group__label").
						Aria("hidden", true).
						ID("chip-group-with-categories-overflow-expanded-label").
						Body(
							app.Text("Category one"),
						),
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("labelledby", "chip-group-with-categories-overflow-expanded-label").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("chip-group-with-categories-overflow-expandedchip_one_toolbar").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflow-expandedremove_chip_one_toolbar chip-group-with-categories-overflow-expandedchip_one_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflow-expandedremove_chip_one_toolbar").
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
												ID("chip-group-with-categories-overflow-expandedchip_two_toolbar").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflow-expandedremove_chip_two_toolbar chip-group-with-categories-overflow-expandedchip_two_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflow-expandedremove_chip_two_toolbar").
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
												ID("chip-group-with-categories-overflow-expandedchip_three_toolbar").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflow-expandedremove_chip_three_toolbar chip-group-with-categories-overflow-expandedchip_three_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflow-expandedremove_chip_three_toolbar").
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
												ID("chip-group-with-categories-overflow-expandedchip_four_toolbar").
												Body(
													app.Text("Chip four"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflow-expandedremove_chip_four_toolbar chip-group-with-categories-overflow-expandedchip_four_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflow-expandedremove_chip_four_toolbar").
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
												ID("chip-group-with-categories-overflow-expandedchip_five_select").
												Body(
													app.Text("Chip five"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflow-expandedremove_chip_five_select chip-group-with-categories-overflow-expandedchip_five_select").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflow-expandedremove_chip_five_select").
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
									app.Button().
										Class("pf-c-chip pf-m-overflow").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												Body(
													app.Text("Show less"),
												),
										),
								),
						),
				),
		)
}
