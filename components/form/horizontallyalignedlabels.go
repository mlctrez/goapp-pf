package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HorizontallyAlignedLabels struct {
	app.Compo
}

func (c *HorizontallyAlignedLabels) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form pf-m-horizontal").
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-horizontal-name").
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
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								ID("form-horizontal-name").
								Name("form-horizontal-name").
								Required(true),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("form-horizontal-info").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Information"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for information field").
								Aria("describedby", "form-horizontal-info").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Textarea().
								Class("pf-c-form-control").
								ID("form-horizontal-info").
								Name("form-horizontal-info").
								Aria("label", "Textarea example"),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Aria("role", "group").
				Aria("labelledby", "form-horizontal-checkbox-legend").
				Body(
					app.Div().
						Class("pf-c-form__group-label pf-m-no-padding-top").
						ID("form-horizontal-checkbox-legend").
						Body(
							app.Span().
								Class("pf-c-form__label").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Label (no top padding)"),
										),
								),
							app.Button().
								Class("pf-c-form__group-label-help").
								Aria("label", "More information for label field").
								Aria("describedby", "form-horizontal-checkbox-legend").
								Body(
									app.I().
										Class("pficon pf-icon-help").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-form__group-control pf-m-stack").
						Body(
							app.Div().
								Class("pf-c-check").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("form-horizontal-checkbox").
										Name("form-horizontal-checkbox"),
									app.Label().
										Class("pf-c-check__label").
										For("form-horizontal-checkbox").
										Body(
											app.Text("Option 1"),
										),
								),
							app.Div().
								Class("pf-c-check").
								Body(
									app.Input().
										Class("pf-c-check__input").
										Type("checkbox").
										ID("form-horizontal-checkbox2").
										Name("form-horizontal-checkbox2"),
									app.Label().
										Class("pf-c-check__label").
										For("form-horizontal-checkbox2").
										Body(
											app.Text("Option 2"),
										),
								),
						),
				),
		)
}
