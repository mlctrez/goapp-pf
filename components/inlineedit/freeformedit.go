package inlineedit

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FreeFormEdit struct {
	app.Compo
}

func (c *FreeFormEdit) Render() app.UI {
	return app.Div().
		Class("pf-c-inline-edit").
		ID("free-form-edit-example").
		Body(
			app.Div().
				Class("pf-c-inline-edit__editable-text").
				Aria("role", "textbox").
				ID("free-form-edit-example-text").
				Aria("label", "Editable text").
				Body(
					app.Text("Free form text"),
				),
		)
}
