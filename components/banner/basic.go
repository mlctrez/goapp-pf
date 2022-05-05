package banner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-banner").
				Body(
					app.Text("Default banner"),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-info").
				Body(
					app.Text("Info banner"),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-danger").
				Body(
					app.Text("Danger banner"),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-success").
				Body(
					app.Text("Success banner"),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-warning").
				Body(
					app.Text("Warning banner"),
				),
		)
}
