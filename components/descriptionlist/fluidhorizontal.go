package descriptionlist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FluidHorizontal struct {
	app.Compo
}

func (c *FluidHorizontal) Render() app.UI {
	return app.Dl().
		Class("pf-c-description-list pf-m-horizontal pf-m-fluid pf-m-2-col").
		Body(
			app.Div().
				Class("pf-c-description-list__group").
				Body(
					app.Dt().
						Class("pf-c-description-list__term").
						Body(
							app.Span().
								Class("pf-c-description-list__text").
								Body(
									app.Text("Name"),
								),
						),
					app.Dd().
						Class("pf-c-description-list__description").
						Body(
							app.Div().
								Class("pf-c-description-list__text").
								Body(
									app.Text("example"),
								),
						),
				),
			app.Div().
				Class("pf-c-description-list__group").
				Body(
					app.Dt().
						Class("pf-c-description-list__term").
						Body(
							app.Span().
								Class("pf-c-description-list__text").
								Body(
									app.Text("Namespace"),
								),
						),
					app.Dd().
						Class("pf-c-description-list__description").
						Body(
							app.Div().
								Class("pf-c-description-list__text").
								Body(
									app.A().
										Href("#").
										Body(
											app.Text("mary-test"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-description-list__group").
				Body(
					app.Dt().
						Class("pf-c-description-list__term").
						Body(
							app.Span().
								Class("pf-c-description-list__text").
								Body(
									app.Text("Labels"),
								),
						),
					app.Dd().
						Class("pf-c-description-list__description").
						Body(
							app.Div().
								Class("pf-c-description-list__text").
								Body(
									app.Text("example"),
								),
						),
				),
			app.Div().
				Class("pf-c-description-list__group").
				Body(
					app.Dt().
						Class("pf-c-description-list__term").
						Body(
							app.Span().
								Class("pf-c-description-list__text").
								Body(
									app.Text("Pod selector"),
								),
						),
					app.Dd().
						Class("pf-c-description-list__description").
						Body(
							app.Div().
								Class("pf-c-description-list__text").
								Body(
									app.Button().
										Class("pf-c-button pf-m-link pf-m-inline").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-button__icon pf-m-start").
												Body(
													app.I().
														Class("fas fa-plus-circle").
														Aria("hidden", true),
												),
											app.Text("app=MyApp"),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-description-list__group").
				Body(
					app.Dt().
						Class("pf-c-description-list__term").
						Body(
							app.Span().
								Class("pf-c-description-list__text").
								Body(
									app.Text("Annotation"),
								),
						),
					app.Dd().
						Class("pf-c-description-list__description").
						Body(
							app.Div().
								Class("pf-c-description-list__text").
								Body(
									app.Text("2 Annotations"),
								),
						),
				),
		)
}
