package datepicker

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CustomWidthInput struct {
	app.Compo
}

func (c *CustomWidthInput) Render() app.UI {
	return app.Div().
		Class("pf-c-date-picker").
		Style("--pf-c-date-picker__input--c-form-control--Width", " 220px;").
		Body(
			app.Div().
				Class("pf-c-date-picker__input").
				Body(
					app.Div().
						Class("pf-c-input-group").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								Value("November 20, 2020").
								ID("basic-input").
								Name("basic-input").
								Aria("label", "Date picker"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Toggle date picker").
								Body(
									app.I().
										Class("fas fa-calendar-alt").
										Aria("hidden", true),
								),
						),
				),
		)
}
