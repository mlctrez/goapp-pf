package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type GridNone struct {
	app.Compo
}

func (c *GridNone) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list pf-m-grid-none").
		Aria("role", "list").
		Aria("label", "Grid none data list example").
		ID("data-list-grid-none").
		Body(
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-grid-none-item-1").
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
												ID("data-list-grid-none-item-1").
												Body(
													app.Text("Cell 1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 2"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 3"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 4"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 5"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 6"),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-grid-none-item-2").
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
												ID("data-list-grid-none-item-2").
												Body(
													app.Text("Cell 1"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 2"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 3"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 4"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 5"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Cell 6"),
										),
								),
						),
				),
		)
}
