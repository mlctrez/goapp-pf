package helpertext

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MultipleStatic struct {
	app.Compo
}

func (c *MultipleStatic) Render() app.UI {
	return app.Div().
		Class("pf-c-helper-text").
		Body(
			app.Div().
				Class("pf-c-helper-text__item").
				Body(
					app.Span().
						Class("pf-c-helper-text__item-text").
						Body(
							app.Text("This is default helper text"),
						),
				),
			app.Div().
				Class("pf-c-helper-text__item").
				Body(
					app.Span().
						Class("pf-c-helper-text__item-text").
						Body(
							app.Text("This is another default helper text in the same block"),
						),
				),
			app.Div().
				Class("pf-c-helper-text__item").
				Body(
					app.Span().
						Class("pf-c-helper-text__item-text").
						Body(
							app.Text("And this is more default text in the same block"),
						),
				),
		)
}
