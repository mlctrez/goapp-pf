package masthead

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicWithMixedContent struct {
	app.Compo
}

func (c *BasicWithMixedContent) Render() app.UI {
	return app.Header().
		Class("pf-c-masthead").
		ID("basic-masthead-with-mixed-content").
		Body(
			app.Span().
				Class("pf-c-masthead__toggle").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Global navigation").
						Body(
							app.I().
								Class("fas fa-bars").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-masthead__main").
				Body(
					app.A().
						Class("pf-c-masthead__brand").
						Href("#").
						Body(
							app.Text("Logo"),
						),
				),
			app.Div().
				Class("pf-c-masthead__content").
				Body(
					app.Div().
						Class("pf-l-flex").
						Body(
							app.Span().
								Body(
									app.Text("Testing text color"),
								),
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("button").
								Body(
									app.Text("testing"),
								),
							app.Div().
								Class("pf-l-flex__item pf-m-align-flex-end").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("button").
										Body(
											app.Text("testing"),
										),
								),
						),
				),
		)
}
