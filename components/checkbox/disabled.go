package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-check").
				Body(
					app.Input().
						Class("pf-c-check__input").
						Type("checkbox").
						ID("check-disabled").
						Name("check-disabled").
						Disabled(true),
					app.Label().
						Class("pf-c-check__label pf-m-disabled").
						For("check-disabled").
						Body(
							app.Text("Check disabled"),
						),
				),
			app.Div().
				Class("pf-c-check").
				Body(
					app.Input().
						Class("pf-c-check__input").
						Type("checkbox").
						ID("check-disabled-2").
						Name("check-disabled-2").
						Checked(true).
						Disabled(true),
					app.Label().
						Class("pf-c-check__label pf-m-disabled").
						For("check-disabled-2").
						Body(
							app.Text("Check disabled checked"),
						),
				),
		)
}
