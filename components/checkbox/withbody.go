package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithBody struct {
	app.Compo
}

func (c *WithBody) Render() app.UI {
	return app.Div().
		Class("pf-c-check").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-with-body").
				Name("check-with-body"),
			app.Label().
				Class("pf-c-check__label").
				For("check-with-body").
				Body(
					app.Text("Check with body"),
				),
			app.Span().
				Class("pf-c-check__body").
				Body(
					app.Text("This is where custom content goes."),
				),
		)
}
