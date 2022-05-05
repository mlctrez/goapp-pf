package divider

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MdInsetNoInsetOnMd3xlInsetOnLgLgInsetOnXl struct {
	app.Compo
}

func (c *MdInsetNoInsetOnMd3xlInsetOnLgLgInsetOnXl) Render() app.UI {
	return app.Div().
		Class("pf-c-divider pf-m-inset-md pf-m-inset-none-on-md pf-m-inset-3xl-on-lg pf-m-inset-lg-on-xl").
		Aria("role", "separator")
}
