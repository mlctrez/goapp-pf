package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TextModifiersDataListText struct {
	app.Compo
}

func (c *TextModifiersDataListText) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Data list with modifiers and text").
		ID("data-list-with-text-modifiers-and-text").
		Body(
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-with-text-modifiers-and-text-item-1").
				Body(
					app.Div().
						Class("pf-c-data-list__item-row").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Span().
												ID("data-list-with-text-modifiers-and-text-item-1").
												Body(
													app.Text("This text will wrap to the next line because it has the default behavior of the data list cell."),
												),
											app.Span().
												Class("pf-c-data-list__text pf-m-truncate").
												Body(
													app.Text("This is data list text, you can apply `pf-m-truncate` directly to the text. This is data list text, you can apply `pf-m-truncate` directly to the text."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will wrap to the next line because it has the default behavior of the data list cell."), app.Span().
												Class("pf-c-data-list__text pf-m-break-word").
												Body(
													app.Text("http://thisisaverylongdatalisttextthatneedstobreakusethebreakwordmodifier.org"),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will wrap to the next line because it has the default behavior of the data list cell."), app.Span().
												Class("pf-c-data-list__text pf-m-nowrap").
												Body(
													app.Text("This is data list text, you can apply `pf-m-nowrap` directly to the text."),
												),
										),
								),
						),
				),
		)
}
