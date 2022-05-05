package datalist

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type TextModifiers struct {
	app.Compo
}

func (c *TextModifiers) Render() app.UI {
	return app.Ul().
		Class("pf-c-data-list").
		Aria("role", "list").
		Aria("label", "Data list with text modifiers").
		ID("data-list-with-text-modifiers").
		Body(
			app.Li().
				Class("pf-c-data-list__item").
				Aria("labelledby", "data-list-with-text-modifiers-item").
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
												ID("data-list-with-text-modifiers-item").
												Body(
													app.Text("This text will wrap to the next line because it has the default behavior of the data list cell."),
												),
										),
									app.Div().
										Class("pf-c-data-list__cell pf-m-truncate").
										Body(
											app.Text("This text will truncate because it is very very long."),
										),
									app.Div().
										Class("pf-c-data-list__cell pf-m-break-word").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
									app.Div().
										Class("pf-c-data-list__cell pf-m-nowrap").
										Body(
											app.Text("This text will not break or wrap."),
										),
								),
						),
					app.Div().
						Class("pf-c-data-list__item-row pf-m-truncate").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will truncate because it is very very long. This text will truncate because it is very very long."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will truncate because it is very very long. This text will truncate because it is very very long."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will truncate because it is very very long. This text will truncate because it is very very long."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will truncate because it is very very long. This text will truncate because it is very very long."),
										),
								),
						),
					app.Div().
						Class("pf-c-data-list__item-row pf-m-break-word").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("http://thisisaverylongurlthatneedstobreakusethebreakwordmodifier.org"),
										),
								),
						),
					app.Div().
						Class("pf-c-data-list__item-row pf-m-nowrap").
						Body(
							app.Div().
								Class("pf-c-data-list__item-content").
								Body(
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will not break or wrap."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will not break or wrap."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will not break or wrap."),
										),
									app.Div().
										Class("pf-c-data-list__cell").
										Body(
											app.Text("This text will not break or wrap."),
										),
								),
						),
				),
		)
}
