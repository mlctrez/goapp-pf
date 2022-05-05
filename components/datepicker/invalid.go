package datepicker

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Invalid struct {
	app.Compo
}

func (c *Invalid) Render() app.UI {
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
								Aria("invalid", true).
								Type("text").
								Value("2020-03-05").
								ID("invalid-input").
								Name("invalid-input").
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
				Class("pf-c-date-picker__helper-text pf-m-error").
				Body(
					app.Text("Invalid date"),
				),
		)
}
