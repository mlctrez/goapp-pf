package accordion

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Fluid struct {
	app.Compo
}

func (c *Fluid) Render() app.UI {
	return app.Div().
		Class("pf-c-accordion").
		Body(
			app.H3().
				Body(
					app.Button().
						Class("pf-c-accordion__toggle").
						Type("button").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-accordion__toggle-text").
								Body(
									app.Text("Item one"),
								),
							app.Span().
								Class("pf-c-accordion__toggle-icon").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-accordion__expanded-content").
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-accordion__expanded-content-body").
						Body(
							app.Text("This text is hidden"),
						),
				),
			app.H3().
				Body(
					app.Button().
						Class("pf-c-accordion__toggle").
						Type("button").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-accordion__toggle-text").
								Body(
									app.Text("Item two"),
								),
							app.Span().
								Class("pf-c-accordion__toggle-icon").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-accordion__expanded-content").
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-accordion__expanded-content-body").
						Body(
							app.Text("This text is hidden"),
						),
				),
			app.H3().
				Body(
					app.Button().
						Class("pf-c-accordion__toggle").
						Type("button").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-accordion__toggle-text").
								Body(
									app.Text("Item three"),
								),
							app.Span().
								Class("pf-c-accordion__toggle-icon").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-accordion__expanded-content").
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-accordion__expanded-content-body").
						Body(
							app.Text("This text is hidden"),
						),
				),
			app.H3().
				Body(
					app.Button().
						Class("pf-c-accordion__toggle pf-m-expanded").
						Type("button").
						Aria("expanded", true).
						Body(
							app.Span().
								Class("pf-c-accordion__toggle-text").
								Body(
									app.Text("Item four"),
								),
							app.Span().
								Class("pf-c-accordion__toggle-icon").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-accordion__expanded-content pf-m-expanded").
				Body(
					app.Div().
						Class("pf-c-accordion__expanded-content-body").
						Body(
							app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis molestie lorem lacinia dolor aliquet faucibus. Suspendisse gravida imperdiet accumsan. Aenean auctor lorem justo, vitae tincidunt enim blandit vel. Aenean quis tempus dolor. Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
						),
				),
			app.H3().
				Body(
					app.Button().
						Class("pf-c-accordion__toggle").
						Type("button").
						Aria("expanded", "false").
						Body(
							app.Span().
								Class("pf-c-accordion__toggle-text").
								Body(
									app.Text("Item five"),
								),
							app.Span().
								Class("pf-c-accordion__toggle-icon").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
						),
				),
			app.Div().
				Class("pf-c-accordion__expanded-content").
				Hidden(true).
				Body(
					app.Div().
						Class("pf-c-accordion__expanded-content-body").
						Body(
							app.Text("This text is hidden"),
						),
				),
		)
}
