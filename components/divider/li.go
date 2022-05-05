package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Li struct {
	app.Compo
}

func (c *Li) Render() app.UI {
	return app.Ul().
		Body(
			app.Li().
				Body(
					app.Text("List item one"),
				),
			app.Li().
				Class("pf-c-divider").
				Aria("role", "separator"),
			app.Li().
				Body(
					app.Text("List item two"),
				),
		)
}
