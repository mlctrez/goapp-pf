package formselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Disabled struct {
	app.Compo
}

func (c *Disabled) Render() app.UI {
	return app.Select().
		Aria("label", "FormSelect Input").
		Class("pf-c-form-control").
		Aria("invalid", "false").
		DataSet("ouia-component-type", "PF4/FormSelect").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-FormSelect-default-7").
		Disabled(true).
		Body(
			app.Option().
				Class("").
				Value("please choose").
				Disabled(true).
				Body(
					app.Text("Select one"),
				),
			app.Option().
				Class("").
				Value("mr").
				Body(
					app.Text("Mr"),
				),
			app.Option().
				Class("").
				Value("miss").
				Body(
					app.Text("Miss"),
				),
			app.Option().
				Class("").
				Value("mrs").
				Body(
					app.Text("Mrs"),
				),
			app.Option().
				Class("").
				Value("ms").
				Body(
					app.Text("Ms"),
				),
			app.Option().
				Class("").
				Value("dr").
				Body(
					app.Text("Dr"),
				),
			app.Option().
				Class("").
				Value("other").
				Body(
					app.Text("Other"),
				),
		)
}
