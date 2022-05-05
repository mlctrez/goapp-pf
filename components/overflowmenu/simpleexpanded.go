package overflowmenu

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SimpleExpanded struct {
	app.Compo
}

func (c *SimpleExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-overflow-menu").
		ID("overflow-menu-simple-expanded").
		Body(
			app.Div().
				Class("pf-c-overflow-menu__content").
				Body(
					app.Div().
						Class("pf-c-overflow-menu__item").
						Body(
							app.Text("Item 1"),
						),
					app.Div().
						Class("pf-c-overflow-menu__item").
						Body(
							app.Text("Item 2"),
						),
					app.Div().
						Class("pf-c-overflow-menu__group").
						Body(
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 3"),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 4"),
								),
							app.Div().
								Class("pf-c-overflow-menu__item").
								Body(
									app.Text("Item 5"),
								),
						),
				),
		)
}
