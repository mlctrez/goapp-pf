package timepicker

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic24Hour struct {
	app.Compo
}

func (c *Basic24Hour) Render() app.UI {
	return app.Div().
		Class("pf-c-date-picker").
		Body(
			app.Div().
				Class("pf-c-date-picker__input").
				Style("--pf-c-date-picker__input--c-form-control--Width", "150px;").
				Body(
					app.Div().
						Class("pf-c-input-group").
						Body(
							app.Div().
								ID("time-picker-1651760007251i3wbp4538rj").
								Body(
									app.Div().
										Style("padding-left", " 0px;").
										Body(
											app.Input().
												Aria("haspopup", "menu").
												ID("time-picker-1651760007251i3wbp4538rj-input").
												Aria("label", "Time picker").
												Placeholder("false").
												AutoComplete(false).
												Class("pf-c-form-control pf-m-icon pf-m-clock pf-c-form-control").
												Type("text").
												Aria("invalid", "false").
												DataSet("ouia-component-type", "PF4/TextInput").
												DataSet("ouia-safe", true).
												DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-7").
												Value("13:06"),
										),
								),
						),
				),
		)
}
