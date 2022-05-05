package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Uncontrolled struct {
	app.Compo
}

func (c *Uncontrolled) Render() app.UI {
	return app.Textarea().
		Class("pf-c-form-control").
		Aria("invalid", "false").
		Aria("label", "uncontrolled text area example").
		Body(
			app.Text("default value"),
		)
}
