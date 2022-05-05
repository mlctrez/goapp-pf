package spinner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type SizesLegacy struct {
	app.Compo
}

func (c *SizesLegacy) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Span().
				Class("pf-c-spinner pf-m-sm").
				Aria("role", "progressbar").
				Aria("label", "Loading...").
				Body(
					app.Span().
						Class("pf-c-spinner__clipper"),
					app.Span().
						Class("pf-c-spinner__lead-ball"),
					app.Span().
						Class("pf-c-spinner__tail-ball"),
				),
			app.Span().
				Class("pf-c-spinner pf-m-md").
				Aria("role", "progressbar").
				Aria("label", "Loading...").
				Body(
					app.Span().
						Class("pf-c-spinner__clipper"),
					app.Span().
						Class("pf-c-spinner__lead-ball"),
					app.Span().
						Class("pf-c-spinner__tail-ball"),
				),
			app.Span().
				Class("pf-c-spinner pf-m-lg").
				Aria("role", "progressbar").
				Aria("label", "Loading...").
				Body(
					app.Span().
						Class("pf-c-spinner__clipper"),
					app.Span().
						Class("pf-c-spinner__lead-ball"),
					app.Span().
						Class("pf-c-spinner__tail-ball"),
				),
			app.Span().
				Class("pf-c-spinner pf-m-xl").
				Aria("role", "progressbar").
				Aria("label", "Loading...").
				Body(
					app.Span().
						Class("pf-c-spinner__clipper"),
					app.Span().
						Class("pf-c-spinner__lead-ball"),
					app.Span().
						Class("pf-c-spinner__tail-ball"),
				),
		)
}
