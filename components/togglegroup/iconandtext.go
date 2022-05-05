package togglegroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type IconAndText struct {
	app.Compo
}

func (c *IconAndText) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-toggle-group").
				Body(
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-copy").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Copy"),
										),
								),
						),
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button").
								Type("button").
								Aria("label", "Undo button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-undo").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Undo"),
										),
								),
						),
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button pf-m-selected").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-share-square").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Share"),
										),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-toggle-group").
				Body(
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Copy"),
										),
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-copy").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Undo"),
										),
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-undo").
												Aria("hidden", true),
										),
								),
						),
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button pf-m-selected").
								Type("button").
								Body(
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Share"),
										),
									app.Span().
										Class("pf-c-toggle-group__icon").
										Body(
											app.I().
												Class("fas fa-share-square").
												Aria("hidden", true),
										),
								),
						),
				),
		)
}
