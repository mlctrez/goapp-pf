package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Navigation struct {
	app.Compo
	ComponentName     string
	ExampleName       string
	PreviousComponent bool
	PreviousExample   bool
	NextExample       bool
	NextComponent     bool
	Current           string
	Max               string
}

type NavClick func(ctx app.Context)

func (c *Navigation) action(name string) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		ctx.NewAction("navigation", app.T("name", name))
	}
}

func (c *Navigation) Render() app.UI {
	return app.Div().Class("pf-c-pagination pf-m-bottom").Body(

		app.Nav().Class("pf-c-pagination__nav").Aria("label", "Pagination").Body(

			app.Div().Class("pf-c-pagination__nav-control pf-m-first").Body(
				app.Button().Class("pf-c-button pf-m-plain").Type("button").
					OnClick(c.action("PreviousComponent")).
					Disabled(!c.PreviousComponent).
					Aria("label", "Go to previousComponent").
					Body(app.I().Class("fas fa-angle-double-left").Aria("hidden", true)),
			),

			app.Div().Class("pf-c-pagination__nav-control pf-m-prev").Body(
				app.Button().Class("pf-c-button pf-m-plain").Type("button").
					OnClick(c.action("PreviousExample")).
					Disabled(!c.PreviousExample).
					Aria("label", "Go to previous example").
					Body(app.I().Class("fas fa-angle-left").Aria("hidden", true)),
			),

			app.Div().Class("pf-c-pagination__nav-page-select").Body(
				app.Input().Class("pf-c-form-control").Style("width", "4em").Aria("label", "Current example").Size(5).
					Type("number").Min("1").Max(c.Max).Value(c.Current),
				app.Span().Aria("hidden", true).Body(app.Text(" of "+c.Max)),
			),

			app.Div().Class("pf-c-pagination__nav-control pf-m-next").Body(
				app.Button().Class("pf-c-button pf-m-plain").Type("button").
					OnClick(c.action("NextExample")).
					Disabled(!c.NextExample).
					Aria("label", "Go to next example").
					Body(app.I().Class("fas fa-angle-right").Aria("hidden", true)),
			),

			app.Div().Class("pf-c-pagination__nav-control pf-m-last").Body(
				app.Button().Class("pf-c-button pf-m-plain").Type("button").
					OnClick(c.action("NextComponent")).
					Disabled(!c.NextComponent).
					Aria("label", "Go to next component").
					Body(app.I().Class("fas fa-angle-double-right").Aria("hidden", true)),
			),

			app.Div().Styles(map[string]string{
				"display":        "flex",
				"flex-direction": "row",
				"align-items":    "stretch",
				"padding-left":   "1em",
			}).Body(
				app.Div().Style("padding", ".3em").Text(c.ComponentName),
				app.I().Style("padding", ".5em").Class("fas fa-angle-right").Aria("hidden", true),
				app.Div().Style("padding", ".3em").Text(c.ExampleName),
			),
		),
	)
}
