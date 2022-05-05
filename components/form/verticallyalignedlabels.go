package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type VerticallyAlignedLabels struct {
	app.Compo
}

func (c *VerticallyAlignedLabels) Render() app.UI {
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
								For("form-vertical-name").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Name"),
										),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Body(
											app.Text("*"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for name field").
								Aria("describedby", "form-vertical-name").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-vertical-name").
								Name("form-vertical-name").
								Required(true),
						),
				),
		)
}
