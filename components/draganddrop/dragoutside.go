package draganddrop

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DragOutside struct {
	app.Compo
}

func (c *DragOutside) Render() app.UI {
	return app.Div().
		Class("pf-c-droppable pf-m-dragging pf-m-drag-outside").
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
				Class("pf-c-draggable pf-m-dragging pf-m-drag-outside").
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
