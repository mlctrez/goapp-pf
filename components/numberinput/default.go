package numberinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Default struct {
	app.Compo
}

func (c *Default) Render() app.UI {
	return app.Div().
		Class("pf-c-number-input").
		Body(
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Minus").
						Body(
							app.Span().
								Class("pf-c-number-input__icon").
								Body(
									app.I().
										Class("fas fa-minus").
										Aria("hidden", true),
								),
						),
					app.Input().
						Class("pf-c-form-control").
						Type("number").
						Value("90").
						Name("number-input-default-name").
						Aria("label", "Number input"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Plus").
						Body(
							app.Span().
								Class("pf-c-number-input__icon").
								Body(
									app.I().
										Class("fas fa-plus").
										Aria("hidden", true),
								),
						),
				),
		)
}
