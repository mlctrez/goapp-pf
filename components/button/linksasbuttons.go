package button

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type LinksAsButtons struct {
	app.Compo
}

func (c *LinksAsButtons) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.A().
				Class("pf-c-button pf-m-primary").
				Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
				Body(
					app.Text("Primary link to W3.org"),
				),
			app.A().
				Class("pf-c-button pf-m-secondary").
				Href("#overview").
				Aria("label", "Read more about button documentation").
				Body(
					app.Text("Secondary link to anchor"),
				),
			app.A().
				Class("pf-c-button pf-m-secondary pf-m-danger").
				Href("#overview").
				Aria("label", "Read more about button documentation").
				Body(
					app.Text("Secondary danger link to anchor"),
				),
			app.A().
				Class("pf-c-button pf-m-tertiary pf-m-disabled").
				Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
				Aria("disabled", true).
				TabIndex(-1).
				Body(
					app.Text("Tertiary link to W3.org"),
				),
			app.A().
				Class("pf-c-button pf-m-link").
				Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
				Body(
					app.Text("Link to W3.org"),
				),
			app.A().
				Class("pf-c-button pf-m-link pf-m-danger").
				Href("https://www.w3.org/TR/WCAG20-TECHS/ARIA8.html#ARIA8-examples").
				Body(
					app.Text("Link danger to W3.org"),
				),
		)
}
