package helpertext

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type DynamicList struct {
	app.Compo
}

func (c *DynamicList) Render() app.UI {
	return app.Ul().
		Class("pf-c-helper-text").
		Body(
			app.Li().
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
							app.Text("Must be at least 14 characters"),
						),
				),
			app.Li().
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
							app.Text("Cannot contain any variation of the word \"redhat\""),
						),
				),
			app.Li().
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
							app.Text("Must include at least 3 of the following: lowercase letter, uppercase letters, numbers, symbols"),
						),
				),
		)
}
