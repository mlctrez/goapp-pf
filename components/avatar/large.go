package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Large struct {
	app.Compo
}

func (c *Large) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-lg").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image large")
}
