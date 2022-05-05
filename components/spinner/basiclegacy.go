package spinner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BasicLegacy struct {
	app.Compo
}

func (c *BasicLegacy) Render() app.UI {
	return app.Span().
		Class("pf-c-spinner").
		Aria("role", "progressbar").
		Aria("label", "Loading...").
		Body(
			app.Span().
				Class("pf-c-spinner__clipper"),
			app.Span().
				Class("pf-c-spinner__lead-ball"),
			app.Span().
				Class("pf-c-spinner__tail-ball"),
		)
}
