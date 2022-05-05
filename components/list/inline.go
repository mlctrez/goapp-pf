package list

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Inline struct {
	app.Compo
}

func (c *Inline) Render() app.UI {
	return app.Ul().
		Class("pf-c-list pf-m-inline").
		Body(
			app.Li().
				Body(
					app.Text("Inline list item 1"),
				),
			app.Li().
				Body(
					app.Text("Inline list item 2"),
				),
			app.Li().
				Body(
					app.Text("Inline list item 3"),
				),
		)
}
