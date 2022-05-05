package formselect

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
											app.Text("Selection:"),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Select().
								ID("selection").
								Aria("label", "FormSelect Input").
								Class("pf-c-form-control pf-m-placeholder").
								Aria("invalid", "false").
								DataSet("ouia-component-type", "PF4/FormSelect").
								DataSet("ouia-safe", true).
								DataSet("ouia-component-id", "OUIA-Generated-FormSelect-default-7").
								Body(
									app.Option().
										Class("").
										Value(true).
										Body(
											app.Text("Select a number"),
										),
									app.Option().
										Class("").
										Value("1").
										Body(
											app.Text("One"),
										),
									app.Option().
										Class("").
										Value("2").
										Body(
											app.Text("Two"),
										),
									app.Option().
										Class("").
										Value("3").
										Body(
											app.Text("Three - the only valid option"),
										),
								),
						),
				),
		)
}
