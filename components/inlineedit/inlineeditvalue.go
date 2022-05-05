package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type InlineEditValue struct {
	app.Compo
}

func (c *InlineEditValue) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit").
		ID("inline-edit-value-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__value").
				Body(
					app.Text("Static value"),
				),
		)
}
