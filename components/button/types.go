package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Types struct {
	app.Compo
}

func (c *Types) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Body(
					app.Text("Submit"),
				),
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("reset").
				Body(
					app.Text("Reset"),
				),
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Body(
					app.Text("Default"),
				),
		)
}
