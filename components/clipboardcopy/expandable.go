package clipboardcopy

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Expandable struct {
	app.Compo
}

func (c *Expandable) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.H4().
				Body(
					app.Text("Editable"),
				),
			app.Div().
				Class("pf-c-clipboard-copy").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("expandable-not-expanded-editable-toggle").
								Aria("labelledby", "expandable-not-expanded-editable-toggle expandable-not-expanded-editable-text-input").
								Aria("controls", "expandable-not-expanded-editable-content").
								Body(
									app.Div().
										Class("pf-c-clipboard-copy__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
							app.Input().
								Class("pf-c-form-control").
								Type("text").
								Value("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion.").
								ID("expandable-not-expanded-editable-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("expandable-not-expanded-editable-copy-button").
								Aria("labelledby", "expandable-not-expanded-editable-copy-button expandable-not-expanded-editable-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-clipboard-copy__expandable-content").
						Hidden(true).
						ID("expandable-not-expanded-editable-content").
						Body(
							app.Text("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion."),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-clipboard-copy pf-m-expanded").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control pf-m-expanded").
								Type("button").
								ID("expandable-expanded-editable-toggle").
								Aria("labelledby", "expandable-expanded-editable-toggle expandable-expanded-editable-text-input").
								Aria("controls", "expandable-expanded-editable-content").
								Body(
									app.Div().
										Class("pf-c-clipboard-copy__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
							app.Input().
								Class("pf-c-form-control").
								ReadOnly(true).
								Type("text").
								Value("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion.").
								ID("expandable-expanded-editable-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("expandable-expanded-editable-copy-button").
								Aria("labelledby", "expandable-expanded-editable-copy-button expandable-expanded-editable-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-clipboard-copy__expandable-content").
						ContentEditable(true).
						ID("expandable-expanded-editable-content").
						Body(
							app.Text("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion."),
						),
				),
			app.Br(),
			app.H4().
				Body(
					app.Text("Read-only"),
				),
			app.Div().
				Class("pf-c-clipboard-copy").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								ID("expandable-not-expanded-readonly-toggle").
								Aria("labelledby", "expandable-not-expanded-readonly-toggle expandable-not-expanded-readonly-text-input").
								Aria("controls", "expandable-not-expanded-readonly-content").
								Body(
									app.Div().
										Class("pf-c-clipboard-copy__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
							app.Input().
								Class("pf-c-form-control").
								ReadOnly(true).
								Type("text").
								Value("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion.").
								ID("expandable-not-expanded-readonly-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("expandable-not-expanded-readonly-copy-button").
								Aria("labelledby", "expandable-not-expanded-readonly-copy-button expandable-not-expanded-readonly-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-clipboard-copy__expandable-content").
						Hidden(true).
						ID("expandable-not-expanded-readonly-content").
						Body(
							app.Text("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion."),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-clipboard-copy pf-m-expanded").
				Body(
					app.Div().
						Class("pf-c-clipboard-copy__group").
						Body(
							app.Button().
								Class("pf-c-button pf-m-control pf-m-expanded").
								Type("button").
								ID("expandable-expanded-readonly-toggle").
								Aria("labelledby", "expandable-expanded-readonly-toggle expandable-expanded-readonly-text-input").
								Aria("controls", "expandable-expanded-readonly-content").
								Body(
									app.Div().
										Class("pf-c-clipboard-copy__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
								),
							app.Input().
								Class("pf-c-form-control").
								ReadOnly(true).
								Type("text").
								Value("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion.").
								ID("expandable-expanded-readonly-text-input").
								Aria("label", "Copyable input"),
							app.Button().
								Class("pf-c-button pf-m-control").
								Type("button").
								Aria("label", "Copy to clipboard").
								ID("expandable-expanded-readonly-copy-button").
								Aria("labelledby", "expandable-expanded-readonly-copy-button expandable-expanded-readonly-text-input").
								Body(
									app.I().
										Class("fas fa-copy").
										Aria("hidden", true),
								),
						),
					app.Div().
						Class("pf-c-clipboard-copy__expandable-content").
						ID("expandable-expanded-readonly-content").
						Body(
							app.Text("This is an editable version of the copy to clipboard component that has an expandable section. Got a lot of text here, need to see all of it? Click that arrow on the left side and check out the resulting expansion."),
						),
				),
		)
}
