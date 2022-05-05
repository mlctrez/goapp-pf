package codeblock

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-code-block").
		Body(
			app.Div().
				Class("pf-c-code-block__header").
				Body(
					app.Div().
						Class("pf-c-code-block__actions").
						Body(
							app.Div().
								Class("pf-c-code-block__actions-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Copy to clipboard").
										Body(
											app.I().
												Class("fas fa-copy").
												Aria("hidden", true),
										),
								),
							app.Div().
								Class("pf-c-code-block__actions-item").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										Aria("label", "Run in Web Terminal").
										Body(
											app.I().
												Class("fas fa-play").
												Aria("hidden", true),
										),
								),
						),
				),
			app.Div().
				Class("pf-c-code-block__content").
				Body(
					app.Pre().
						Class("pf-c-code-block__pre").
						Body(
							app.Code().
								Class("pf-c-code-block__code").
								Body(
									app.Text("apiVersion: helm.openshift.io/v1beta1/\nkind: HelmChartRepository\nmetadata:\nname: azure-sample-repo\nspec:\nconnectionConfig:\nurl: https://raw.githubusercontent.com/Azure-Samples/helm-charts/master/docs"),
								),
						),
				),
		)
}
