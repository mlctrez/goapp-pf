package list

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithHorizontalRules struct {
	app.Compo
}

func (c *WithHorizontalRules) Render() app.UI {
	return app.Ul().
		Class("pf-c-list pf-m-plain pf-m-bordered").
		Body(
			app.Li().
				Body(
					app.Text("Donec blandit a lorem id convallis."),
				),
			app.Li().
				Body(
					app.Text("Integer in volutpat libero."),
				),
			app.Li().
				Body(
					app.Text("Donec a diam tellus."),
				),
			app.Li().
				Body(
					app.Text("Aenean nec tortor orci."),
				),
			app.Li().
				Body(
					app.Text("Vivamus maximus ultricies pulvinar."),
				),
		)
}
