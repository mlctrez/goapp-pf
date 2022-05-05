package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image")
}
