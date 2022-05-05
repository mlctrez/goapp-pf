package textarea

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Validated struct {
	app.Compo
}

func (c *Validated) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form").
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("selection").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Comments:"),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Textarea().
								Class("pf-c-form-control").
								Aria("invalid", "false").
								Required(true).
								Aria("label", "invalid text area example"),
							app.Div().
								Class("pf-c-form__helper-text").
								ID("selection-helper").
								Aria("live", "polite").
								Body(
									app.Text("Share your thoughts."),
								),
						),
				),
		)
}
