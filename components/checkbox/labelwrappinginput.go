package checkbox

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LabelWrappingInput struct {
	app.Compo
}

func (c *LabelWrappingInput) Render() app.UI {
	return app.Label().
		Class("pf-c-check").
		For("check-label-wrapping-input").
		Body(
			app.Input().
				Class("pf-c-check__input").
				Type("checkbox").
				ID("check-label-wrapping-input").
				Name("check-label-wrapping-input"),
			app.Span().
				Class("pf-c-check__label").
				Body(
					app.Text("Check label wraps input"),
				),
		)
}
