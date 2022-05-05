package popover

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutHeaderfooter struct {
	app.Compo
}

func (c *WithoutHeaderfooter) Render() app.UI {
	return app.Div().
		Class("pf-c-popover pf-m-right").
		Aria("modal", true).
		Aria("label", "Popover with no header example").
		Aria("describedby", "popover-no-header-body").
		Body(
			app.Div().
				Class("pf-c-popover__arrow"),
			app.Div().
				Class("pf-c-popover__content").
				Body(
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Close").
						Body(
							app.I().
								Class("fas fa-times").
								Aria("hidden", true),
						),
					app.Div().
						Class("pf-c-popover__body").
						ID("popover-no-header-body").
						Body(
							app.Text("Popovers are triggered by click rather than hover. Click again to close."),
						),
				),
		)
}
