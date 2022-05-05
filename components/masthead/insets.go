package masthead

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Insets struct {
	app.Compo
}

func (c *Insets) Render() app.UI {
	return app.Header().
		Class("pf-c-masthead pf-m-inset-sm").
		ID("inset-masthead").
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
					app.Span().
						Body(
							app.Text("Content"),
						),
				),
		)
}
