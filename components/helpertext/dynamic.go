package helpertext

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Dynamic struct {
	app.Compo
}

func (c *Dynamic) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-helper-text").
				Body(
					app.Div().
						Class("pf-c-helper-text__item pf-m-dynamic").
						Body(
							app.Span().
								Class("pf-c-helper-text__item-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-minus").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-helper-text__item-text").
								Body(
									app.Text("This is default helper text"),
								),
						),
				),
			app.Div().
				Class("pf-c-helper-text").
				Body(
					app.Div().
						Class("pf-c-helper-text__item pf-m-dynamic pf-m-indeterminate").
						Body(
							app.Span().
								Class("pf-c-helper-text__item-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-minus").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-helper-text__item-text").
								Body(
									app.Text("This is indeterminate helper text"),
								),
						),
				),
			app.Div().
				Class("pf-c-helper-text").
				Body(
					app.Div().
						Class("pf-c-helper-text__item pf-m-dynamic pf-m-warning").
						Body(
							app.Span().
								Class("pf-c-helper-text__item-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-exclamation-triangle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-helper-text__item-text").
								Body(
									app.Text("This is warning helper text"),
								),
						),
				),
			app.Div().
				Class("pf-c-helper-text").
				Body(
					app.Div().
						Class("pf-c-helper-text__item pf-m-dynamic pf-m-success").
						Body(
							app.Span().
								Class("pf-c-helper-text__item-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-check-circle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-helper-text__item-text").
								Body(
									app.Text("This is success helper text"),
								),
						),
				),
			app.Div().
				Class("pf-c-helper-text").
				Body(
					app.Div().
						Class("pf-c-helper-text__item pf-m-dynamic pf-m-error").
						Body(
							app.Span().
								Class("pf-c-helper-text__item-icon").
								Body(
									app.I().
										Class("fas fa-fw fa-exclamation-circle").
										Aria("hidden", true),
								),
							app.Span().
								Class("pf-c-helper-text__item-text").
								Body(
									app.Text("This is error helper text"),
								),
						),
				),
		)
}
