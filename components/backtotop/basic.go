package backtotop

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-back-to-top").
		Body(
			app.A().
				Class("pf-c-button pf-m-primary").
				Href("#").
				Body(
					app.Text("Back to top"), app.Span().
						Class("pf-c-button__icon pf-m-end").
						Body(
							app.I().
								Class("fas fa-angle-up").
								Aria("hidden", true),
						),
				),
		)
}
