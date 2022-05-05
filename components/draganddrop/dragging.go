package draganddrop

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Dragging struct {
	app.Compo
}

func (c *Dragging) Render() app.UI {
	return app.Div().
		Class("pf-c-droppable pf-m-dragging").
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
				Class("pf-c-draggable pf-m-dragging").
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
