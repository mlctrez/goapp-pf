package formcontrol

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Select struct {
	app.Compo
}

func (c *Select) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Select().
				Class("pf-c-form-control pf-m-placeholder").
				ID("select-standard").
				Name("select-standard").
				Aria("label", "Standard select example").
				Body(
					app.Option().
						Value(true).
						Selected(true).
						Disabled(true).
						Body(
							app.Text("Please choose"),
						),
					app.Option().
						Value("Mr").
						Body(
							app.Text("Mr"),
						),
					app.Option().
						Value("Miss").
						Body(
							app.Text("Miss"),
						),
					app.Option().
						Value("Mrs").
						Body(
							app.Text("Mrs"),
						),
					app.Option().
						Value("Ms").
						Body(
							app.Text("Ms"),
						),
					app.Option().
						Value("Dr").
						Body(
							app.Text("Dr"),
						),
					app.Option().
						Value("Other").
						Body(
							app.Text("Other"),
						),
				),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control pf-m-placeholder").
				ID("select-placeholder-enabled").
				Name("select-placeholder-enabled").
				Aria("label", "Placeholder enabled select example").
				Body(
					app.Option().
						Value(true).
						Selected(true).
						Body(
							app.Text("Please choose"),
						),
					app.Option().
						Value("Mr").
						Body(
							app.Text("Mr"),
						),
					app.Option().
						Value("Miss").
						Body(
							app.Text("Miss"),
						),
					app.Option().
						Value("Mrs").
						Body(
							app.Text("Mrs"),
						),
					app.Option().
						Value("Ms").
						Body(
							app.Text("Ms"),
						),
					app.Option().
						Value("Dr").
						Body(
							app.Text("Dr"),
						),
					app.Option().
						Value("Other").
						Body(
							app.Text("Other"),
						),
				),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control").
				ID("select-group").
				Name("select-group").
				Aria("label", "Select group example").
				Body(
					app.OptGroup().
						Label("Group 1").
						Body(
							app.Option().
								Value("Option 1").
								Body(
									app.Text("The first option"),
								),
							app.Option().
								Value("Option 2").
								Selected(true).
								Body(
									app.Text("The second option is selected by default"),
								),
						),
					app.OptGroup().
						Label("Group 2").
						Body(
							app.Option().
								Value("Option 3").
								Body(
									app.Text("The third option"),
								),
							app.Option().
								Value("Option 4").
								Body(
									app.Text("The fourth option"),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control pf-m-success").
				ID("select-group-success").
				Name("select-group-success").
				Aria("label", "Success state select group example").
				Body(
					app.Option().
						Value(true).
						Body(
							app.Text("Valid option"),
						),
					app.OptGroup().
						Label("Group 1").
						Body(
							app.Option().
								Value("Option 1").
								Body(
									app.Text("The first option"),
								),
							app.Option().
								Value("Option 2").
								Body(
									app.Text("The second option"),
								),
						),
					app.OptGroup().
						Label("Group 2").
						Body(
							app.Option().
								Value("Option 3").
								Body(
									app.Text("The third option"),
								),
							app.Option().
								Value("Option 4").
								Body(
									app.Text("The fourth option"),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control pf-m-warning").
				ID("select-group-warning").
				Name("select-group-warning").
				Aria("label", "Warning state select group example").
				Body(
					app.Option().
						Value(true).
						Body(
							app.Text("Warning option"),
						),
					app.OptGroup().
						Label("Group 1").
						Body(
							app.Option().
								Value("Option 1").
								Body(
									app.Text("The first option"),
								),
							app.Option().
								Value("Option 2").
								Body(
									app.Text("The second option"),
								),
						),
					app.OptGroup().
						Label("Group 2").
						Body(
							app.Option().
								Value("Option 3").
								Body(
									app.Text("The third option"),
								),
							app.Option().
								Value("Option 4").
								Body(
									app.Text("The fourth option"),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control").
				Required(true).
				Aria("invalid", true).
				ID("select-group-error").
				Name("select-group-error").
				Aria("label", "Error state select group example").
				Body(
					app.Option().
						Value(true).
						Body(
							app.Text("Invalid option"),
						),
					app.OptGroup().
						Label("Group 1").
						Body(
							app.Option().
								Value("Option 1").
								Body(
									app.Text("The first option"),
								),
							app.Option().
								Value("Option 2").
								Body(
									app.Text("The second option"),
								),
						),
					app.OptGroup().
						Label("Group 2").
						Body(
							app.Option().
								Value("Option 3").
								Body(
									app.Text("The third option"),
								),
							app.Option().
								Value("Option 4").
								Body(
									app.Text("The fourth option"),
								),
						),
				),
			app.Br(),
			app.Br(),
		)
}
