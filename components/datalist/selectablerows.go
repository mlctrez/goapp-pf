package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SelectableRows struct {
	app.Compo
}

func (c *SelectableRows) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Selectable rows data list example").
		ID("data-list-selectable-rows").
		Body(
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded pf-m-selectable pf-m-selected").
				Aria("labelledby", "data-list-selectable-rows-item-1").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-rows-item-1").
												Body(
													app.Text("Primary content"),
												),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-expanded pf-m-selectable pf-m-selected").
				Aria("labelledby", "data-list-selectable-rows-item-2").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-rows-item-2").
												Body(
													app.Text("Secondary content (selected)"),
												),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item pf-m-selectable").
				Aria("labelledby", "data-list-selectable-rows-item-3").
				TabIndex(0).
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-selectable-rows-item-3").
												Body(
													app.Text("Tertiary content"),
												),
										),
								),
						),
				),
		)
}
