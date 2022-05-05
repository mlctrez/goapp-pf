package dropdown

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DropdownWithImageAndText struct {
	app.Compo
}

func (c *DropdownWithImageAndText) Render() app.UI {
	return app.Div().
		Class("pf-c-dropdown pf-m-expanded").
		Body(
			app.Button().
				Class("pf-c-dropdown__toggle").
				ID("dropdown-with-image-and-text-example-button").
				Aria("expanded", true).
				Type("button").
				Body(
					app.Span().
						Class("pf-c-dropdown__toggle-image").
						Body(
							app.Img().
								Class("pf-c-avatar").
								Src("/assets/images/img_avatar.svg").
								Alt("Avatar image"),
						),
					app.Span().
						Class("pf-c-dropdown__toggle-text").
						Body(
							app.Text("Ned Username"),
						),
					app.Span().
						Class("pf-c-dropdown__toggle-icon").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-dropdown__menu").
				Body(
					app.Section().
						Class("pf-c-dropdown__group").
						Body(
							app.Div().
								Class("pf-c-dropdown__menu-item pf-m-text").
								Body(
									app.Text("Text"),
								),
							app.Div().
								Class("pf-c-dropdown__menu-item pf-m-text").
								Body(
									app.Text("More text"),
								),
						),
					app.Hr().
						Class("pf-c-divider"),
					app.Section().
						Class("pf-c-dropdown__group").
						Body(
							app.Ul().
								Body(
									app.Li().
										Body(
											app.A().
												Class("pf-c-dropdown__menu-item").
												Href("#").
												Body(
													app.Text("My profile"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-dropdown__menu-item").
												Href("#").
												Body(
													app.Text("User management"),
												),
										),
									app.Li().
										Body(
											app.A().
												Class("pf-c-dropdown__menu-item").
												Href("#").
												Body(
													app.Text("Logout"),
												),
										),
								),
						),
				),
		)
}
