package jumplinks

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableResponsive struct {
	app.Compo
}

func (c *ExpandableResponsive) Render() app.UI {
	return app.Nav().
		Class("pf-c-jump-links pf-m-vertical pf-m-non-expandable-on-md pf-m-expandable-on-lg pf-m-non-expandable-on-xl pf-m-expandable").
		Aria("label", "Jump to section").
		Body(
			app.Div().
				Class("pf-c-jump-links__header").
				Body(
					app.Div().
						Class("pf-c-jump-links__toggle").
						Body(
							app.Button().
								Class("pf-c-button pf-m-plain").
								Type("button").
								Aria("label", "Toggle jump links").
								Aria("expanded", "false").
								Body(
									app.Span().
										Class("pf-c-jump-links__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-jump-links__toggle-text").
										Body(
											app.Text("Jump to section"),
										),
								),
						),
					app.Div().
						Class("pf-c-jump-links__label").
						Body(
							app.Text("Jump to section"),
						),
				),
			app.Ul().
				Class("pf-c-jump-links__list").
				Body(
					app.Li().
						Class("pf-c-jump-links__item").
						Body(
							app.A().
								Class("pf-c-jump-links__link").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-jump-links__link-text").
										Body(
											app.Text("Inactive section"),
										),
								),
						),
					app.Li().
						Class("pf-c-jump-links__item pf-m-current").
						Body(
							app.A().
								Class("pf-c-jump-links__link").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-jump-links__link-text").
										Body(
											app.Text("Active section"),
										),
								),
						),
					app.Li().
						Class("pf-c-jump-links__item").
						Body(
							app.A().
								Class("pf-c-jump-links__link").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-jump-links__link-text").
										Body(
											app.Text("Inactive section"),
										),
								),
						),
					app.Li().
						Class("pf-c-jump-links__item").
						Body(
							app.A().
								Class("pf-c-jump-links__link").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-jump-links__link-text").
										Body(
											app.Text("Inactive section"),
										),
								),
						),
					app.Li().
						Class("pf-c-jump-links__item").
						Body(
							app.A().
								Class("pf-c-jump-links__link").
								Href("#").
								Body(
									app.Span().
										Class("pf-c-jump-links__link-text").
										Body(
											app.Text("Inactive section"),
										),
								),
						),
				),
		)
}
