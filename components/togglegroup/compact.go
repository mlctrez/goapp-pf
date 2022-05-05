package togglegroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-toggle-group pf-m-compact").
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
											app.Text("Option 1"),
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
											app.Text("Option 2"),
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
											app.Text("Option 3"),
										),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-toggle-group pf-m-compact").
				Body(
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
											app.Text("Option 1"),
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
											app.Text("Option 2"),
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
											app.Text("Option 3"),
										),
								),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-toggle-group pf-m-compact").
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
											app.Text("Option 1"),
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
											app.Text("Option 2"),
										),
								),
						),
					app.Div().
						Class("pf-c-toggle-group__item").
						Body(
							app.Button().
								Class("pf-c-toggle-group__button").
								Type("button").
								Disabled(true).
								Body(
									app.Span().
										Class("pf-c-toggle-group__text").
										Body(
											app.Text("Option 3"),
										),
								),
						),
				),
		)
}
