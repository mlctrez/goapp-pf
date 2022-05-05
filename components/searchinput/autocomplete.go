package searchinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Autocomplete struct {
	app.Compo
}

func (c *Autocomplete) Render() app.UI {
	return app.Div().
		Class("pf-c-search-input").
		Body(
			app.Div().
				Class("pf-c-search-input__bar").
				Body(
					app.Span().
						Class("pf-c-search-input__text").
						Body(
							app.Span().
								Class("pf-c-search-input__icon").
								Body(
									app.I().
										Class("fas fa-search fa-fw").
										Aria("hidden", true),
								),
							app.Input().
								Class("pf-c-search-input__text-input").
								ID("search-input-autocomplete-text-input").
								Type("text").
								Placeholder("false").
								Aria("label", "Keyword search").
								Value("app"),
						),
					app.Span().
						Class("pf-c-search-input__utilities").
						Body(
							app.Span().
								Class("pf-c-search-input__clear").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Clear").
										Body(
											app.I().
												Class("fas fa-times fa-fw").
												Aria("hidden", true),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-search-input__menu").
				Body(
					app.Ul().
						Class("pf-c-search-input__menu-list").
						Body(
							app.Li().
								Class("pf-c-search-input__menu-list-item").
								Body(
									app.Button().
										Class("pf-c-search-input__menu-item").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-search-input__menu-item-text").
												Body(
													app.Text("apple"),
												),
										),
								),
							app.Li().
								Class("pf-c-search-input__menu-list-item").
								Body(
									app.Button().
										Class("pf-c-search-input__menu-item").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-search-input__menu-item-text").
												Body(
													app.Text("appleby"),
												),
										),
								),
							app.Li().
								Class("pf-c-search-input__menu-list-item").
								Body(
									app.Button().
										Class("pf-c-search-input__menu-item").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-search-input__menu-item-text").
												Body(
													app.Text("appleseed"),
												),
										),
								),
							app.Li().
								Class("pf-c-search-input__menu-list-item").
								Body(
									app.Button().
										Class("pf-c-search-input__menu-item").
										Type("button").
										Body(
											app.Span().
												Class("pf-c-search-input__menu-item-text").
												Body(
													app.Text("appleton"),
												),
										),
								),
						),
				),
		)
}
