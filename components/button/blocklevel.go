package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BlockLevel struct {
	app.Compo
}

func (c *BlockLevel) Render() app.UI {
	return app.Button().
		Class("pf-c-button pf-m-primary pf-m-block").
		Type("button").
		Body(
			app.Text("Block level button"),
		)
}
