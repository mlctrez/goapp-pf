package formcontrol

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Input struct {
	app.Compo
}

func (c *Input) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Input().
				Class("pf-c-form-control").
				Type("text").
				Value("Standard").
				ID("input-standard").
				Aria("label", "Standard input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control").
				Type("text").
				Placeholder("false").
				ID("input-placeholder").
				Aria("label", "Placeholder input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control").
				ReadOnly(true).
				Type("text").
				Value("Readonly").
				ID("input-readonly").
				Aria("label", "Readonly input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-success").
				Type("text").
				Value("Success").
				ID("input-success").
				Aria("label", "Success state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-warning").
				Type("text").
				Value("Warning").
				ID("input-warning").
				Aria("label", "Warning state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control").
				Required(true).
				Type("text").
				Value("Error").
				ID("input-error").
				Aria("invalid", true).
				Aria("label", "Error state input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control").
				Disabled(true).
				Type("text").
				Value("Disabled").
				ID("input-disabled").
				Aria("label", "Disabled input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-expanded").
				Type("text").
				Value("Expanded").
				ID("input-expanded").
				Aria("label", "Expanded input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon pf-m-calendar").
				Type("text").
				Value("Calendar").
				ID("input-calendar").
				Name("input-calendar").
				Aria("label", "Calendar input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon pf-m-clock").
				Type("text").
				Value("Clock").
				ID("input-clock").
				Name("input-clock").
				Aria("label", "Clock input example"),
			app.Br(),
			app.Br(),
			app.Input().
				Class("pf-c-form-control pf-m-icon").
				Type("text").
				Value("Custom icon").
				ID("input-custom-icon").
				Name("custom-icon").
				Aria("label", "Custom icon input example"),
		)
}
