package chipgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleInlineChipGroupOverflow struct {
	app.Compo
}

func (c *SimpleInlineChipGroupOverflow) Render() app.UI {
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
												ID("simple-inline-chip-group-overflowchip_one_select_collapsed").
												Body(
													app.Text("Chip one"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-overflowremove_chip_one_select_collapsed simple-inline-chip-group-overflowchip_one_select_collapsed").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-overflowremove_chip_one_select_collapsed").
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
												ID("simple-inline-chip-group-overflowchip_two_select_collapsed").
												Body(
													app.Text("Chip two"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-overflowremove_chip_two_select_collapsed simple-inline-chip-group-overflowchip_two_select_collapsed").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-overflowremove_chip_two_select_collapsed").
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
												ID("simple-inline-chip-group-overflowchip_three_select_collapsed").
												Body(
													app.Text("Chip three"),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("labelledby", "simple-inline-chip-group-overflowremove_chip_three_select_collapsed simple-inline-chip-group-overflowchip_three_select_collapsed").
												Aria("label", "Remove").
												ID("simple-inline-chip-group-overflowremove_chip_three_select_collapsed").
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
