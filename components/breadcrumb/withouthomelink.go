package breadcrumb

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithoutHomeLink struct {
	app.Compo
}

func (c *WithoutHomeLink) Render() app.UI {
	return app.Nav().
		Class("pf-c-breadcrumb").
		Aria("label", "breadcrumb").
		Body(
			app.Ol().
				Class("pf-c-breadcrumb__list").
				Body(
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.Text("Section home"),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section title"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section title"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section title"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link").
								Body(
									app.Text("Section title"),
								),
						),
					app.Li().
						Class("pf-c-breadcrumb__item").
						Body(
							app.Span().
								Class("pf-c-breadcrumb__item-divider").
								Body(
									app.I().
										Class("fas fa-angle-right").
										Aria("hidden", true),
								),
							app.A().
								Href("#").
								Class("pf-c-breadcrumb__link pf-m-current").
								Aria("current", "page").
								Body(
									app.Text("Section landing"),
								),
						),
				),
		)
}
