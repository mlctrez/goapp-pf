package progress

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type OutsideStaticWidthMeasure struct {
	app.Compo
}

func (c *OutsideStaticWidthMeasure) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-example").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("1%"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100").
						Aria("valuenow", "1").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "1%;"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-2-example").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-2-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("50%"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100").
						Aria("valuenow", "50").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "50%;"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-3-example").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-3-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("100%"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100").
						Aria("valuenow", "100").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "100%;"),
						),
				),
			app.Br(),
			app.Br(),
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-4-example").
				Style("--pf-c-progress__measure--m-static-width--MinWidth", " 6ch;").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-4-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("1,000"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100000").
						Aria("valuenow", "1000").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "1%;"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-5-example").
				Style("--pf-c-progress__measure--m-static-width--MinWidth", " 6ch;").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-5-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("50,000"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100000").
						Aria("valuenow", "50000").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "50%;"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-progress pf-m-outside pf-m-lg").
				ID("progress-outside-static-width-6-example").
				Style("--pf-c-progress__measure--m-static-width--MinWidth", " 6ch;").
				Body(
					app.Div().
						Class("pf-c-progress__description").
						ID("progress-outside-static-width-6-example-description"),
					app.Div().
						Class("pf-c-progress__status").
						Aria("hidden", true).
						Body(
							app.Span().
								Class("pf-c-progress__measure pf-m-static-width").
								Body(
									app.Text("100,000"),
								),
						),
					app.Div().
						Class("pf-c-progress__bar").
						Aria("role", "progressbar").
						Aria("valuemin", "0").
						Aria("valuemax", "100000").
						Aria("valuenow", "100000").
						Aria("label", "Progress status").
						Body(
							app.Div().
								Class("pf-c-progress__indicator").
								Style("width", "100%;"),
						),
				),
		)
}
