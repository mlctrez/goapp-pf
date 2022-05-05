package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ChipGroupWithCategoriesRemovable struct {
	app.Compo
}

func (c *ChipGroupWithCategoriesRemovable) Render() app.UI {
	return app.Div().
		Class("pf-c-chip-group pf-m-category").
		Body(
			app.Div().
				Class("pf-c-chip-group__main").
				Body(
					app.Span().
						Class("pf-c-chip-group__label").
						Aria("hidden", true).
						ID("chip-group-with-categories-removable-label").
						Body(
							app.Text("Category one"),
						),
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("labelledby", "chip-group-with-categories-removable-label").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("chip-group-with-categories-removablechip_one_toolbar").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_one_toolbar chip-group-with-categories-removablechip_one_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_one_toolbar").
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
												ID("chip-group-with-categories-removablechip_two_toolbar").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_two_toolbar chip-group-with-categories-removablechip_two_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_two_toolbar").
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
												ID("chip-group-with-categories-removablechip_three_toolbar").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_three_toolbar chip-group-with-categories-removablechip_three_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_three_toolbar").
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
												ID("chip-group-with-categories-removablechip_four_toolbar").
												Body(
													app.Text("Chip four"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_four_toolbar chip-group-with-categories-removablechip_four_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_four_toolbar").
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
												ID("chip-group-with-categories-removablechip_five_toolbar").
												Body(
													app.Text("Chip five"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_five_toolbar chip-group-with-categories-removablechip_five_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_five_toolbar").
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
												ID("chip-group-with-categories-removablechip_six_toolbar").
												Body(
													app.Text("Chip six"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-removableremove_chip_six_toolbar chip-group-with-categories-removablechip_six_toolbar").
												Aria("label", "Remove").
												ID("chip-group-with-categories-removableremove_chip_six_toolbar").
												Body(
													app.I().
														Class("fas fa-times").
														Aria("hidden", true),
												),
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
						Aria("labelledby", "chip-group-with-categories-removable-button chip-group-with-categories-removable-label").
						Aria("label", "Close chip group").
						ID("chip-group-with-categories-removable-button").
						Body(
							app.I().
								Class("fas fa-times-circle").
								Aria("hidden", true),
						),
				),
		)
}
