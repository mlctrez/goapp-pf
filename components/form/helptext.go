package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HelpText struct {
	app.Compo
}

func (c *HelpText) Render() app.UI {
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
								For("form-help-text-name").
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
								Required(true).
								Type("text").
								ID("form-help-text-name").
								Name("form-help-text-name").
								Aria("describedby", "form-help-text-name-helper"),
							app.P().
								Class("pf-c-form__helper-text").
								ID("form-help-text-name-helper").
								Aria("live", "polite").
								Body(
									app.Text("This is helper text."),
								),
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
								For("form-help-text-email").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("E-mail"),
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
								Class("pf-c-form-control pf-m-warning").
								Required(true).
								Type("text").
								ID("form-help-text-email").
								Name("form-help-text-email").
								Aria("describedby", "form-help-text-email-helper"),
							app.P().
								Class("pf-c-form__helper-text pf-m-warning").
								ID("form-help-text-email-helper").
								Aria("live", "polite").
								Body(
									app.Text("This is helper text for a warning input."),
								),
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
								For("form-help-text-address").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Address"),
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
								Required(true).
								Type("text").
								ID("form-help-text-address").
								Name("form-help-text-address").
								Aria("invalid", true).
								Aria("describedby", "form-help-text-address-helper"),
							app.P().
								Class("pf-c-form__helper-text pf-m-error").
								ID("form-help-text-address-helper").
								Aria("live", "polite").
								Body(
									app.Text("This is helper text for an invalid input."),
								),
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
								For("form-help-text-comment").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Comment"),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control pf-m-success").
								Value("This is a valid comment").
								Type("text").
								ID("form-help-text-comment").
								Name("form-help-text-comment").
								Aria("describedby", "form-help-text-comment-helper"),
							app.P().
								Class("pf-c-form__helper-text pf-m-success").
								ID("form-help-text-comment-helper").
								Aria("live", "polite").
								Body(
									app.Text("This is helper text for success input."),
								),
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
								For("form-help-text-info").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Body(
											app.Text("Information"),
										),
								),
						),
					app.Textarea().
						Class("pf-c-form-control").
						ID("form-help-text-info").
						Name("form-help-text-info").
						Aria("invalid", true).
						Aria("describedby", "form-help-text-info-helper"),
					app.P().
						Class("pf-c-form__helper-text pf-m-error").
						ID("form-help-text-info-helper").
						Aria("live", "polite").
						Body(
							app.Span().
								Class("pf-c-form__helper-text-icon").
								Body(
									app.I().
										Class("fas fa-exclamation-circle").
										Aria("hidden", true),
								),
							app.Text("This is helper text with an icon."),
						),
				),
		)
}
