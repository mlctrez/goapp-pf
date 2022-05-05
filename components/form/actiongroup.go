package form

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ActionGroup struct {
	app.Compo
}

func (c *ActionGroup) Render() app.UI {
	return app.Form().
		NoValidate(true).
		Class("pf-c-form").
		Body(
			app.Div().
				Class("pf-c-form__group pf-m-action").
				Body(
					app.Div().
						Class("pf-c-form__actions").
						Body(
							app.Button().
								Class("pf-c-button pf-m-primary").
								Type("submit").
								Body(
									app.Text("Submit form"),
								),
							app.Button().
								Class("pf-c-button pf-m-link").
								Type("reset").
								Body(
									app.Text("Reset form"),
								),
						),
				),
		)
}
