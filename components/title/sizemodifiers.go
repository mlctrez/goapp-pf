package title

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SizeModifiers struct {
	app.Compo
}

func (c *SizeModifiers) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.H1().
				Class("pf-c-title pf-m-4xl").
				Body(
					app.Text("4xl title"),
				),
			app.H1().
				Class("pf-c-title pf-m-3xl").
				Body(
					app.Text("3xl title"),
				),
			app.H1().
				Class("pf-c-title pf-m-2xl").
				Body(
					app.Text("2xl title"),
				),
			app.H1().
				Class("pf-c-title pf-m-xl").
				Body(
					app.Text("xl title"),
				),
			app.H1().
				Class("pf-c-title pf-m-lg").
				Body(
					app.Text("lg title"),
				),
			app.H1().
				Class("pf-c-title pf-m-md").
				Body(
					app.Text("md title"),
				),
		)
}
