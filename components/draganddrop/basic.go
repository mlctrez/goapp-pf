package draganddrop

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-droppable").
		Body(
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
			app.Div().
				Class("pf-c-draggable").
				Body(
					app.Text("Item"),
				),
		)
}
