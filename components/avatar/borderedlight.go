package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BorderedLight struct {
	app.Compo
}

func (c *BorderedLight) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-light").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image light")
}
