package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hr struct {
	app.Compo
}

func (c *Hr) Render() app.UI {
	return app.Hr().
		Class("pf-c-divider")
}
