package brand

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Img().
		Class("pf-c-brand").
		Src("/assets/images/pf_logo.svg").
		Alt("PatternFly logo")
}
