package banner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BannerWithLinks struct {
	app.Compo
}

func (c *BannerWithLinks) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-banner").
				Body(
					app.Text("Default banner with a"), app.A().
						Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
						Body(
							app.Text("link"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner").
				Body(
					app.Text("Default banner with a"), app.A().
						Class("pf-m-disabled").
						Aria("role", "link").
						Aria("disabled", true).
						Body(
							app.Text("disabled link"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-info").
				Body(
					app.Text("Info banner with an"), app.Button().
						Class("pf-c-button pf-m-inline pf-m-link").
						Type("button").
						Body(
							app.Text("inline link button"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-warning").
				Body(
					app.Text("Warning banner with an"), app.A().
						Class("pf-c-button pf-m-inline pf-m-link").
						Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
						Body(
							app.Text("inline link button (anchor)"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-banner pf-m-danger").
				Body(
					app.Text("Danger banner with a"), app.Button().
						Class("pf-c-button pf-m-link pf-m-inline").
						Type("button").
						Disabled(true).
						Body(
							app.Text("disabled inline link button"),
						),
				),
		)
}
