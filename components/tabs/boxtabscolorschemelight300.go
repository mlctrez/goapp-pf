package tabs

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type BoxTabsColorSchemeLight300 struct {
	app.Compo
}

func (c *BoxTabsColorSchemeLight300) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-tabs pf-m-box pf-m-color-scheme--light-300").
				ID("box-tabs-alt-color-scheme").
				Body(
					app.Button().
						Class("pf-c-tabs__scroll-button").
						Disabled(true).
						Aria("hidden", true).
						Aria("label", "Scroll left").
						Body(
							app.I().
								Class("fas fa-angle-left").
								Aria("hidden", true),
						),
					app.Ul().
						Class("pf-c-tabs__list").
						Body(
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-tabs-alt-color-scheme-users-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Users"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item pf-m-current").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-tabs-alt-color-scheme-containers-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Containers"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-tabs-alt-color-scheme-database-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Database"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										Disabled(true).
										ID("box-tabs-alt-color-scheme-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Disabled"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link pf-m-aria-disabled").
										Aria("disabled", true).
										ID("box-tabs-alt-color-scheme-aria-disabled-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("ARIA disabled"),
												),
										),
								),
							app.Li().
								Class("pf-c-tabs__item").
								Body(
									app.Button().
										Class("pf-c-tabs__link").
										ID("box-tabs-alt-color-scheme-network-wired-link").
										Body(
											app.Span().
												Class("pf-c-tabs__item-text").
												Body(
													app.Text("Network"),
												),
										),
								),
						),
					app.Button().
						Class("pf-c-tabs__scroll-button").
						Disabled(true).
						Aria("hidden", true).
						Aria("label", "Scroll right").
						Body(
							app.I().
								Class("fas fa-angle-right").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("tabs-example-block tabs-example-block--m-color-scheme--light-300"),
		)
}
