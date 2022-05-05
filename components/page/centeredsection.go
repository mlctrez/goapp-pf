package page

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CenteredSection struct {
	app.Compo
}

func (c *CenteredSection) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.Div().
						Class("pf-c-page__header-brand").
						Body(
							app.Div().
								Class("pf-c-page__header-brand-toggle").
								Body(
									app.Text("toggle"),
								),
							app.A().
								Href("#").
								Class("pf-c-page__header-brand-link").
								Body(
									app.Text("Logo"),
								),
						),
					app.Div().
						Class("pf-c-page__header-tools").
						Body(
							app.Text("header-tools"),
						),
				),
			app.Main().
				Class("pf-c-page__main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-limit-width pf-m-align-center").
						Body(
							app.Div().
								Class("pf-c-page__main-body").
								Body(
									app.Div().
										Class("pf-c-card").
										Body(
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("When a width limited page section is wider than the value of"), app.Code().
														Body(
															app.Text("--pf-c-page--section--m-limit-width--MaxWidth"),
														),
													app.Text(", the section will be centered in the main section."),
													app.Br(),
													app.Br(),
													app.Text("The content in this example is placed in a card to better illustrate how the section behaves when it is centered. A card is not required to center a page section."),
												),
										),
								),
						),
				),
		)
}
