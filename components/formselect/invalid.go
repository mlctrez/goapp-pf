package formselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Invalid struct {
	app.Compo
}

func (c *Invalid) Render() app.UI {
	return app.Select().
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
					app.Text("Three"),
				),
		)
}
