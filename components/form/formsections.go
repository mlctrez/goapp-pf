package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FormSections struct {
	app.Compo
}

func (c *FormSections) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form").
		Body(
			app.Section().
				Class("pf-c-form__section").
				Aria("role", "group").
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("form-section-example-form-section-1-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Body(
													app.Text("Form section 1 inputs"),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("text").
										ID("form-section-example-form-section-1-input").
										Name("form-section-example-form-section-1-input").
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
										For("form-section-example-form-section-1-input-2").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Body(
													app.Text("Form section 1 inputs"),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("text").
										ID("form-section-example-form-section-1-input-2").
										Name("form-section-example-form-section-1-input-2").
										Required(true),
								),
						),
				),
			app.Section().
				Class("pf-c-form__section").
				Aria("role", "group").
				Aria("labelledby", "form-section-example-section2-title").
				Body(
					app.Div().
						Class("pf-c-form__section-title").
						ID("form-section-example-section2-title").
						Aria("hidden", true).
						Body(
							app.Text("Section 2 title (optional)"),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("form-section-example-form-section-2-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Body(
													app.Text("Form section 2 inputs"),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("text").
										ID("form-section-example-form-section-2-input").
										Name("form-section-example-form-section-2-input").
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
										For("form-section-example-form-section-2-input-2").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Body(
													app.Text("Form section 2 inputs"),
												),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("text").
										ID("form-section-example-form-section-2-input-2").
										Name("form-section-example-form-section-2-input-2").
										Required(true),
								),
						),
				),
		)
}
