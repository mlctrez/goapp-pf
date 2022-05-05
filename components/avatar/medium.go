package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Medium struct {
	app.Compo
}

func (c *Medium) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-md").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image medium")
}
