package togglegroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithIcon struct {
	app.Compo
}

func (c *WithIcon) Render() app.UI {
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
								Aria("label", "Copy button").
								Body(
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
								Aria("label", "Undo button").
								Body(
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
								Class("pf-c-toggle-group__button").
								Type("button").
								Aria("label", "Share button").
								Body(
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
			app.Br(),
			app.Div().
				Class("pf-c-toggle-group").
				Body(
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button pf-m-selected").
								Type("button").
								Aria("label", "Copy button").
								Body(
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
								Aria("label", "Undo button").
								Body(
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
								Class("pf-c-toggle-group__button").
								Type("button").
								Aria("label", "Share button").
								Body(
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
								Aria("label", "Copy button").
								Body(
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
								Aria("label", "Undo button").
								Body(
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
								Class("pf-c-toggle-group__button").
								Type("button").
								Aria("label", "Share button").
								Disabled(true).
								Body(
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
