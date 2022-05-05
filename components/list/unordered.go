package list

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Unordered struct {
	app.Compo
}

func (c *Unordered) Render() app.UI {
	return app.Ul().
		Class("pf-c-list").
		Body(
			app.Li().
				Body(
					app.Text("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
				),
			app.Li().
				Body(
					app.Text("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
				),
			app.Li().
				Body(
					app.Text("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."), app.Ul().
						Class("pf-c-list").
						Body(
							app.Li().
								Body(
									app.Text("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
								),
							app.Li().
								Body(
									app.Text("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
								),
							app.Li().
								Body(
									app.Text("Ut venenatis, nisl scelerisque."), app.Ol().
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
										),
								),
						),
				),
			app.Li().
				Body(
					app.Text("Ut non enim metus."),
				),
		)
}
