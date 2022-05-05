package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Basic data list example").
		ID("data-list-basic").
		Body(
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-basic-item-1").
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
												ID("data-list-basic-item-1").
												Body(
													app.Text("Primary content"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("Secondary content"),
										),
								),
						),
				),
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-basic-item-2").
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell pf-m-no-fill").
										Body(
											app.Span().
												ID("data-list-basic-item-2").
												Body(
													app.Text("Secondary content (pf-m-no-fill)"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell pf-m-no-fill pf-m-align-right").
										Body(
											app.Text("Secondary content (pf-m-align-right pf-m-no-fill)"),
										),
								),
						),
				),
		)
}
