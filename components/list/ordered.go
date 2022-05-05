package list

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Ordered struct {
	app.Compo
}

func (c *Ordered) Render() app.UI {
	return app.Ol().
		Class("pf-c-list").
		Body(
			app.Li().
				Body(
					app.Text("Donec blandit a lorem id convallis."),
				),
			app.Li().
				Body(
					app.Text("Cras gravida arcu at diam gravida gravida."),
				),
			app.Li().
				Body(
					app.Text("Integer in volutpat libero."),
				),
			app.Li().
				Body(
					app.Text("Donec a diam tellus."),
				),
			app.Li().
				Body(
					app.Text("Etiam auctor nisl et."), app.Ul().
						Class("pf-c-list").
						Body(
							app.Li().
								Body(
									app.Text("Donec blandit a lorem id convallis."),
								),
							app.Li().
								Body(
									app.Text("Cras gravida arcu at diam gravida gravida."),
								),
							app.Li().
								Body(
									app.Text("Integer in volutpat libero."), app.Ol().
										Class("pf-c-list").
										Body(
											app.Li().
												Body(
													app.Text("Donec blandit a lorem id convallis."),
												),
											app.Li().
												Body(
													app.Text("Cras gravida arcu at diam gravida gravida."),
												),
										),
								),
						),
				),
			app.Li().
				Body(
					app.Text("Aenean nec tortor orci."),
				),
			app.Li().
				Body(
					app.Text("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
				),
			app.Li().
				Body(
					app.Text("Vivamus maximus ultricies pulvinar."),
				),
		)
}
