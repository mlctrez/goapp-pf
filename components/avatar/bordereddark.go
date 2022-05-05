package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BorderedDark struct {
	app.Compo
}

func (c *BorderedDark) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-dark").
		Src("/assets/images/img_avatar-dark.svg").
		Alt("Avatar image dark")
}
