package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type IconSpriteVariants struct {
	app.Compo
}

func (c *IconSpriteVariants) Render() app.UI {
	return app.Div().
		ID("ws-react-c-text-area-icon-sprite-variants").
		Class("ws-react-c-text-area pf-u-h-100").
		Body(
			app.Textarea().
				Class("pf-c-form-control pf-m-icon-sprite pf-m-success").
				Aria("invalid", "false").
				Aria("label", "success icon sprite text area example"),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-icon-sprite pf-m-warning").
				Aria("invalid", "false").
				Aria("label", "warning icon sprite text input example"),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-icon-sprite").
				Aria("invalid", true).
				Aria("label", "error icon sprite text area example"),
		)
}
