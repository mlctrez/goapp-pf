package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ChipGroupWithCategoriesOverflow struct {
	app.Compo
}

func (c *ChipGroupWithCategoriesOverflow) Render() app.UI {
	return app.Div().
		Class("pf-c-chip-group pf-m-category").
		Body(
			app.Div().
				Class("pf-c-chip-group__main").
				Body(
					app.Span().
						Class("pf-c-chip-group__label").
						Aria("hidden", true).
						ID("chip-group-with-categories-overflow-label").
						Body(
							app.Text("Category one"),
						),
					app.Ul().
						Class("pf-c-chip-group__list").
						Aria("role", "list").
						Aria("labelledby", "chip-group-with-categories-overflow-label").
						Body(
							app.Li().
								Class("pf-c-chip-group__list-item").
								Body(
									app.Div().
										Class("pf-c-chip").
										Body(
											app.Span().
												Class("pf-c-chip__text").
												ID("chip-group-with-categories-overflowchip_one_toolbar_collapsed").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflowremove_chip_one_toolbar_collapsed chip-group-with-categories-overflowchip_one_toolbar_collapsed").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflowremove_chip_one_toolbar_collapsed").
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
												ID("chip-group-with-categories-overflowchip_two_toolbar_collapsed").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflowremove_chip_two_toolbar_collapsed chip-group-with-categories-overflowchip_two_toolbar_collapsed").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflowremove_chip_two_toolbar_collapsed").
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
												ID("chip-group-with-categories-overflowchip_three_toolbar_collapsed").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "chip-group-with-categories-overflowremove_chip_three_toolbar_collapsed chip-group-with-categories-overflowchip_three_toolbar_collapsed").
												Aria("label", "Remove").
												ID("chip-group-with-categories-overflowremove_chip_three_toolbar_collapsed").
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
													app.Text("2 more"),
												),
										),
								),
						),
				),
		)
}
