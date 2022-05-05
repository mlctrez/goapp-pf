package avatar

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Small struct {
	app.Compo
}

func (c *Small) Render() app.UI {
	return app.Img().
		Class("pf-c-avatar pf-m-sm").
		Src("/assets/images/img_avatar-light.svg").
		Alt("Avatar image small")
}
