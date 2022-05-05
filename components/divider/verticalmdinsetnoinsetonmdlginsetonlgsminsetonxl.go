package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticalMdInsetNoInsetOnMdLgInsetOnLgSmInsetOnXl struct {
	app.Compo
}

func (c *VerticalMdInsetNoInsetOnMdLgInsetOnLgSmInsetOnXl) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-vertical pf-m-inset-md pf-m-inset-none-on-md pf-m-inset-lg-on-lg pf-m-inset-sm-on-xl").
		Aria("role", "separator")
}
