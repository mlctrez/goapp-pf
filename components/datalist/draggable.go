package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Draggable struct {
	app.Compo
}

func (c *Draggable) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				ID("draggable-help").
				Body(
					app.Text("Activate the reorder button and use the arrow keys to reorder the list or use your mouse to drag/reorder. Press escape to cancel the reordering."),
				),
			app.Ul().
				Class("pf-c-data-list pf-m-compact").
				Aria("role", "list").
				Aria("label", "Draggable data list rows").
				ID("data-list-draggable").
				Body(
					app.Li().
						Class("pf-c-data-list__item").
						Aria("labelledby", "data-list-draggable-item-1").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Button().
												Class("pf-c-data-list__item-draggable-button pf-m-disabled").
												Type("button").
												Aria("label", "Reorder").
												Aria("pressed", "false").
												ID("data-list-draggable-draggable-button-1").
												Aria("describedby", "draggable-help").
												Aria("labelledby", "data-list-draggable-draggable-button-1 data-list-draggable-item-1").
												Disabled(true).
												Body(
													app.Span().
														Class("pf-c-data-list__item-draggable-icon").
														Body(
															app.I().
																Class("fas fa-grip-vertical"),
														),
												),
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-draggable-item-1-checkbox").
														Aria("labelledby", "data-list-draggable-item-1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Span().
														Class("pf-c-data-list__cell-text").
														ID("data-list-draggable-item-1").
														Body(
															app.Text("Draggable icon disabled"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-data-list__item").
						Aria("labelledby", "data-list-draggable-item-2").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Button().
												Class("pf-c-data-list__item-draggable-button").
												Type("button").
												Aria("label", "Reorder").
												Aria("pressed", "false").
												ID("data-list-draggable-draggable-button-2").
												Aria("describedby", "draggable-help").
												Aria("labelledby", "data-list-draggable-draggable-button-2 data-list-draggable-item-2").
												Body(
													app.Span().
														Class("pf-c-data-list__item-draggable-icon").
														Body(
															app.I().
																Class("fas fa-grip-vertical"),
														),
												),
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-draggable-item-2-checkbox").
														Aria("labelledby", "data-list-draggable-item-2"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Span().
														Class("pf-c-data-list__cell-text").
														ID("data-list-draggable-item-2").
														Body(
															app.Text("List item"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-data-list__item pf-m-ghost-row").
						Aria("labelledby", "data-list-draggable-item-3").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Button().
												Class("pf-c-data-list__item-draggable-button").
												Type("button").
												Aria("label", "Reorder").
												Aria("pressed", "false").
												ID("data-list-draggable-draggable-button-3").
												Aria("describedby", "draggable-help").
												Aria("labelledby", "data-list-draggable-draggable-button-3 data-list-draggable-item-3").
												Disabled(true).
												Body(
													app.Span().
														Class("pf-c-data-list__item-draggable-icon").
														Body(
															app.I().
																Class("fas fa-grip-vertical"),
														),
												),
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-draggable-item-3-checkbox").
														Aria("labelledby", "data-list-draggable-item-3"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Span().
														Class("pf-c-data-list__cell-text").
														ID("data-list-draggable-item-3").
														Body(
															app.Text("Ghost row"),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-data-list__item").
						Aria("labelledby", "data-list-draggable-item-4").
						Body(
							app.Div().
								Class("pf-c-data-list__item-row").
								Body(
									app.Div().
										Class("pf-c-data-list__item-control").
										Body(
											app.Button().
												Class("pf-c-data-list__item-draggable-button").
												Type("button").
												Aria("label", "Reorder").
												Aria("pressed", "false").
												ID("data-list-draggable-draggable-button-4").
												Aria("describedby", "draggable-help").
												Aria("labelledby", "data-list-draggable-draggable-button-4 data-list-draggable-item-4").
												Body(
													app.Span().
														Class("pf-c-data-list__item-draggable-icon").
														Body(
															app.I().
																Class("fas fa-grip-vertical"),
														),
												),
											app.Div().
												Class("pf-c-data-list__check").
												Body(
													app.Input().
														Type("checkbox").
														Name("data-list-draggable-item-4-checkbox").
														Aria("labelledby", "data-list-draggable-item-4"),
												),
										),
									app.Div().
										Class("pf-c-data-list__item-content").
										Body(
											app.Div().
												Class("pf-c-data-list__cell").
												Body(
													app.Span().
														Class("pf-c-data-list__cell-text").
														ID("data-list-draggable-item-4").
														Body(
															app.Text("List item"),
														),
												),
										),
								),
						),
				),
			app.Div().
				Class("pf-screen-reader").
				Aria("live", "assertive").
				Body(
					app.Text("This is the aria-live section that provides real-time feedback to the user."),
				),
		)
}
