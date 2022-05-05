package formselect

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Grouped struct {
	app.Compo
}

func (c *Grouped) Render() app.UI {
	return app.Select().
		Aria("label", "FormSelect Input").
		Class("pf-c-form-control").
		Aria("invalid", "false").
		DataSet("ouia-component-type", "PF4/FormSelect").
		DataSet("ouia-safe", true).
		DataSet("ouia-component-id", "OUIA-Generated-FormSelect-default-7").
		Body(
			app.OptGroup().
				Class("").
				Label("Group1").
				Body(
					app.Option().
						Class("").
						Value("1").
						Body(
							app.Text("The first option"),
						),
					app.Option().
						Class("").
						Value("2").
						Body(
							app.Text("Second option is selected by default"),
						),
				),
			app.OptGroup().
				Class("").
				Label("Group2").
				Body(
					app.Option().
						Class("").
						Value("3").
						Body(
							app.Text("The third option"),
						),
					app.Option().
						Class("").
						Value("4").
						Body(
							app.Text("The fourth option"),
						),
				),
			app.OptGroup().
				Disabled(true).
				Class("").
				Label("Group3").
				Body(
					app.Option().
						Class("").
						Value("5").
						Body(
							app.Text("The fifth option"),
						),
					app.Option().
						Class("").
						Value("6").
						Body(
							app.Text("The sixth option"),
						),
				),
		)
}
