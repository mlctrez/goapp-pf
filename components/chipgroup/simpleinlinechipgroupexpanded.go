package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleInlineChipGroupExpanded struct {
	app.Compo
}

func (c *SimpleInlineChipGroupExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-chip-group").
		Body(
			app.Div().
				Class("pf-c-chip-group__main").
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
												ID("simple-inline-chip-group-expandedchip_one_select").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-expandedremove_chip_one_select simple-inline-chip-group-expandedchip_one_select").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-expandedremove_chip_one_select").
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
												ID("simple-inline-chip-group-expandedchip_two_select").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-expandedremove_chip_two_select simple-inline-chip-group-expandedchip_two_select").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-expandedremove_chip_two_select").
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
												ID("simple-inline-chip-group-expandedchip_three_select").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-expandedremove_chip_three_select simple-inline-chip-group-expandedchip_three_select").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-expandedremove_chip_three_select").
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
												ID("simple-inline-chip-group-expandedchip_four_select").
												Body(
													app.Text("Chip four"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-expandedremove_chip_four_select simple-inline-chip-group-expandedchip_four_select").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-expandedremove_chip_four_select").
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
												ID("simple-inline-chip-group-expandedchip_five_select").
												Body(
													app.Text("Chip five"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-expandedremove_chip_five_select simple-inline-chip-group-expandedchip_five_select").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-expandedremove_chip_five_select").
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
