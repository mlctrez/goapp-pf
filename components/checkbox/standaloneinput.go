package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type StandaloneInput struct {
	app.Compo
}

func (c *StandaloneInput) Render() app.UI {
	return app.Div().
		Class("pf-c-check pf-m-standalone").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-standalone-input").
				Name("check-standalone-input").
				Aria("label", "Standalone input"),
		)
}
