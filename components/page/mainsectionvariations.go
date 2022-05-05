package page

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type MainSectionVariations struct {
	app.Compo
}

func (c *MainSectionVariations) Render() app.UI {
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
						Class("pf-c-page__main-nav").
						Body(
							app.Text("`.pf-c-page__main-nav` for tertiary navigation"),
						),
					app.Section().
						Class("pf-c-page__main-breadcrumb").
						Body(
							app.Text("`.pf-c-page__main-breadcrumb` for breadcrumbs"),
						),
					app.Section().
						Class("pf-c-page__main-section").
						Body(
							app.Text("`.pf-c-page__main-section` for main sections"),
						),
					app.Section().
						Class("pf-c-page__main-tabs").
						Body(
							app.Text("`.pf-c-page__main-tabs` for tabs"),
						),
					app.Section().
						Class("pf-c-page__main-wizard").
						Body(
							app.Text("`.pf-c-page__main-wizard` for wizards"),
						),
				),
		)
}
