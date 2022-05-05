package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Plain struct {
	app.Compo
}

func (c *Plain) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-menu-toggle pf-m-plain").
				Type("button").
				Aria("expanded", "false").
				Aria("label", "Actions").
				Body(
					app.I().
						Class("fas fa-ellipsis-v").
						Aria("hidden", true),
				),
			app.Button().
				Class("pf-c-menu-toggle pf-m-plain pf-m-expanded").
				Type("button").
				Aria("expanded", true).
				Aria("label", "Actions").
				Body(
					app.I().
						Class("fas fa-ellipsis-v").
						Aria("hidden", true),
				),
			app.Button().
				Class("pf-c-menu-toggle pf-m-plain").
				Type("button").
				Aria("expanded", "false").
				Disabled(true).
				Aria("label", "Actions").
				Body(
					app.I().
						Class("fas fa-ellipsis-v").
						Aria("hidden", true),
				),
		)
}
