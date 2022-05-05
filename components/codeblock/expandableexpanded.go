package codeblock

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type ExpandableExpanded struct {
	app.Compo
}

func (c *ExpandableExpanded) Render() app.UI {
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
									app.Text("apiVersion: helm.openshift.io/v1beta1/\nkind: HelmChartRepository\nmetadata:\nname: azure-sample-repo"), app.Div().
										Class("pf-c-expandable-section pf-m-expanded pf-m-detached").
										Body(
											app.Div().
												Class("pf-c-expandable-section__content").
												ID("code-block-expandable-expanded-content").
												Body(
													app.Text("spec:\nconnectionConfig:\nurl: https://raw.githubusercontent.com/Azure-Samples/helm-charts/master/docs"),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-expandable-section pf-m-expanded pf-m-detached").
						Body(
							app.Button().
								Type("button").
								Class("pf-c-expandable-section__toggle").
								Aria("expanded", true).
								Aria("controls", "code-block-expandable-expanded-content").
								Body(
									app.Span().
										Class("pf-c-expandable-section__toggle-icon pf-m-expand-top").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-expandable-section__toggle-text").
										Body(
											app.Text("Show less"),
										),
								),
						),
				),
		)
}
