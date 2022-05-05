package timepicker

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithSeconds struct {
	app.Compo
}

func (c *WithSeconds) Render() app.UI {
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
								ID("time-picker-1651760014346pu1t0qcl0bp").
								Body(
									app.Div().
										Style("padding-left", " 0px;").
										Body(
											app.Input().
												Aria("haspopup", "menu").
												ID("time-picker-1651760014346pu1t0qcl0bp-input").
												Aria("label", "Time picker").
												Placeholder("false").
												AutoComplete(false).
												Class("pf-c-form-control pf-m-icon pf-m-clock pf-c-form-control").
												Type("text").
												Aria("invalid", "false").
												DataSet("ouia-component-type", "PF4/TextInput").
												DataSet("ouia-safe", true).
												DataSet("ouia-component-id", "OUIA-Generated-TextInputBase-7").
												Value("3:35:20 PM"),
										),
								),
						),
				),
		)
}
