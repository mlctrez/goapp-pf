package sidebar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Stack struct {
	app.Compo
}

func (c *Stack) Render() app.UI {
	return app.Div().
		Class("pf-c-sidebar pf-m-stack").
		Body(
			app.Div().
				Class("pf-c-sidebar__main").
				Body(
					app.Div().
						Class("pf-c-sidebar__panel").
						Body(
							app.Text("Sidebar panel"),
						),
					app.Div().
						Class("pf-c-sidebar__content").
						Body(
							app.Div().
								Class("pf-c-content").
								Body(
									app.P().
										Body(
											app.Text("Forces a stacked layout."),
										),
									app.P().
										Body(
											app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse dapibus nulla id augue dictum commodo. Donec mollis arcu massa, sollicitudin venenatis est rutrum vitae. Integer pulvinar ligula at augue mollis, ac pulvinar arcu semper. Maecenas nisi lorem, malesuada ac lectus nec, porta pretium neque. Ut convallis libero sit amet metus mattis, vel facilisis lorem malesuada. Duis consectetur ante sit amet magna efficitur, a interdum leo vulputate."),
										),
									app.P().
										Body(
											app.Text("Praesent at odio nec sapien ultrices tincidunt in non mauris. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Duis consectetur nisl quis facilisis faucibus. Sed eu bibendum risus. Suspendisse porta euismod tortor, at elementum odio suscipit sed. Cras eget ultrices urna, ac feugiat lectus. Integer a pharetra velit, in imperdiet mi. Phasellus vel hendrerit velit. Vestibulum ut augue vitae erat vulputate bibendum a ut magna."),
										),
								),
						),
				),
		)
}
