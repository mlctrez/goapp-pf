package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExtraLarge struct {
	app.Compo
}

func (c *ExtraLarge) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-xl").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image extra large")
}
