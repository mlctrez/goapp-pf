package page

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MainSectionPadding struct {
	app.Compo
}

func (c *MainSectionPadding) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.Div().
						Class("pf-c-page__header-brand").
						Body(
							app.Div().
								Class("pf-c-page__header-brand-toggle").
								Body(
									app.Text("toggle"),
								),
							app.A().
								Href("#").
								Class("pf-c-page__header-brand-link").
								Body(
									app.Text("Logo"),
								),
						),
					app.Div().
						Class("pf-c-page__header-tools").
						Body(
							app.Text("header-tools"),
						),
				),
			app.Div().
				Class("pf-c-page__sidebar").
				Body(
					app.Div().
						Class("pf-c-page__sidebar-body").
						Body(
							app.Text("pf-c-nav"),
						),
				),
			app.Main().
				Class("pf-c-page__main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section").
						Body(
							app.Text("This `.pf-c-page__main-section` has default padding."),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-no-padding pf-m-light").
						Body(
							app.Text("This `.pf-c-page__main-section` uses `.pf-m-no-padding` to remove all padding."),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-no-padding pf-m-padding-on-md").
						Body(
							app.Text("This `.pf-c-page__main-section` uses `.pf-m-no-padding .pf-m-padding-on-md` to remove padding up to the `md` breakpoint."),
						),
				),
		)
}
