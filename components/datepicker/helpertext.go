package datepicker

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type HelperText struct {
	app.Compo
}

func (c *HelperText) Render() app.UI {
	return app.Div().
		Class("pf-c-date-picker").
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
								Value("2020-03-05").
								ID("helper-text-input").
								Name("helper-text-input").
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
			app.Div().
				Class("pf-c-date-picker__helper-text").
				Body(
					app.Text("Select a date."),
				),
		)
}
