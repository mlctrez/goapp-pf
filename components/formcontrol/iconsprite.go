package formcontrol

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type IconSprite struct {
	app.Compo
}

func (c *IconSprite) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Input().
				Class("pf-c-form-control pf-m-success pf-m-icon-sprite").
				Type("text").
				Value("Success").
				ID("input-success").
				Aria("label", "Success state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-warning pf-m-icon-sprite").
				Type("text").
				Value("Warning").
				ID("input-warning").
				Aria("label", "Warning state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon-sprite").
				Required(true).
				Type("text").
				Value("Error").
				ID("input-error").
				Aria("invalid", true).
				Aria("label", "Error state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-search pf-m-icon-sprite").
				Type("search").
				Value("Search").
				ID("input-search").
				Name("search-input").
				Aria("label", "Search input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon pf-m-calendar pf-m-icon-sprite").
				Type("text").
				Value("Calendar").
				ID("input-calendar").
				Name("input-calendar").
				Aria("label", "Calendar input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon pf-m-clock pf-m-icon-sprite").
				Type("text").
				Value("Clock").
				ID("input-clock").
				Name("input-clock").
				Aria("label", "Clock input example"),
			app.Br(),
			app.Br(),
			app.Select().
				Class("pf-c-form-control pf-m-success pf-m-icon-sprite").
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
				Class("pf-c-form-control pf-m-warning pf-m-icon-sprite").
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
				Class("pf-c-form-control pf-m-icon-sprite").
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
			app.Textarea().
				Class("pf-c-form-control pf-m-success pf-m-icon-sprite").
				Name("textarea-success").
				ID("textarea-success").
				Aria("label", "Success state textarea example").
				Body(
					app.Text("Success"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-warning pf-m-icon-sprite").
				Name("textarea-warning").
				ID("textarea-warning").
				Aria("label", "Warning state textarea example").
				Body(
					app.Text("Warning"),
				),
			app.Br(),
			app.Br(),
			app.Textarea().
				Class("pf-c-form-control pf-m-icon-sprite").
				Required(true).
				Name("textarea-error").
				ID("textarea-error").
				Aria("label", "Error state textarea example").
				Aria("invalid", true).
				Body(
					app.Text("Error"),
				),
		)
}
