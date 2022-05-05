package list

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithSmallIcons struct {
	app.Compo
}

func (c *WithSmallIcons) Render() app.UI {
	return app.Ul().
		Class("pf-c-list pf-m-plain").
		Body(
			app.Li().
				Class("pf-c-list__item").
				Body(
					app.Span().
						Class("pf-c-list__item-icon").
						Body(
							app.I().
								Class("fas fa-book-open fa-fw").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-list__item-text").
						Body(
							app.Text("List item one"),
						),
				),
			app.Li().
				Class("pf-c-list__item").
				Body(
					app.Span().
						Class("pf-c-list__item-icon").
						Body(
							app.I().
								Class("fas fa-key fa-fw").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-list__item-text").
						Body(
							app.Text("List item two"),
						),
				),
			app.Li().
				Class("pf-c-list__item").
				Body(
					app.Span().
						Class("pf-c-list__item-icon").
						Body(
							app.I().
								Class("fas fa-desktop fa-fw").
								Aria("hidden", true),
						),
					app.Span().
						Class("pf-c-list__item-text").
						Body(
							app.Text("List item three"),
						),
				),
		)
}
