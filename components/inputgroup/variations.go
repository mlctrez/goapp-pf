package inputgroup

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Variations struct {
	app.Compo
}

func (c *Variations) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						ID("textAreaButton1").
						Body(
							app.Text("Button"),
						),
					app.Textarea().
						Class("pf-c-form-control").
						Name("textarea1").
						ID("textarea1").
						Aria("label", "Textarea with buttons").
						Aria("describedby", "textAreaButton1"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Body(
							app.Text("Button"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Textarea().
						Class("pf-c-form-control").
						Name("textarea2").
						ID("textarea2").
						Aria("label", "Textarea with button").
						Aria("describedby", "textAreaButton2"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						ID("textAreaButton2").
						Body(
							app.Text("Button"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						ID("textAreaButton3").
						Body(
							app.Text("Button"),
						),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Body(
							app.Text("Button"),
						),
					app.Textarea().
						Class("pf-c-form-control").
						Name("textarea3").
						ID("textarea3").
						Aria("label", "Textarea with buttons").
						Aria("describedby", "textAreaButton3"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Body(
							app.Text("Button"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Div().
						Class("pf-c-select").
						Style("width", " 100px;").
						Body(
							app.Span().
								ID("select-example-collapsed1-label").
								Hidden(true).
								Body(
									app.Text("Choose one"),
								),
							app.Button().
								Class("pf-c-select__toggle").
								Type("button").
								ID("select-example-collapsed1-toggle").
								Aria("haspopup", true).
								Aria("expanded", "false").
								Aria("labelledby", "select-example-collapsed1-label select-example-collapsed1-toggle").
								Body(
									app.Div().
										Class("pf-c-select__toggle-wrapper").
										Body(
											app.Span().
												Class("pf-c-select__toggle-text").
												Body(
													app.Text("Select"),
												),
										),
									app.Span().
										Class("pf-c-select__toggle-arrow").
										Body(
											app.I().
												Class("fas fa-caret-down").
												Aria("hidden", true),
										),
								),
							app.Ul().
								Class("pf-c-select__menu").
								Aria("role", "listbox").
								Aria("labelledby", "select-example-collapsed1-label").
								Hidden(true).
								Style("width", " 100px;").
								Body(
									app.Li().
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item").
												Aria("role", "option").
												Body(
													app.Text("Running"),
												),
										),
									app.Li().
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item pf-m-selected").
												Aria("role", "option").
												Aria("selected", true).
												Body(
													app.Text("Stopped"), app.Span().
														Class("pf-c-select__menu-item-icon").
														Body(
															app.I().
																Class("fas fa-check").
																Aria("hidden", true),
														),
												),
										),
									app.Li().
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item").
												Aria("role", "option").
												Body(
													app.Text("Down"),
												),
										),
									app.Li().
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item").
												Aria("role", "option").
												Body(
													app.Text("Degraded"),
												),
										),
									app.Li().
										Aria("role", "presentation").
										Body(
											app.Button().
												Class("pf-c-select__menu-item").
												Aria("role", "option").
												Body(
													app.Text("Needs maintenance"),
												),
										),
								),
						),
					app.Input().
						Class("pf-c-form-control").
						Type("text").
						ID("textInput4").
						Name("textInput4").
						Aria("label", "Input with select and button").
						Aria("describedby", "inputSelectButton1"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						ID("inputSelectButton1").
						Body(
							app.Text("Button"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Span().
						Class("pf-c-input-group__text").
						Body(
							app.I().
								Class("fas fa-dollar-sign").
								Aria("hidden", true),
						),
					app.Input().
						Class("pf-c-form-control").
						Type("number").
						ID("textInput5").
						Name("textInput5").
						Aria("label", " Dollar amount input example"),
					app.Span().
						Class("pf-c-input-group__text").
						Body(
							app.Text(".00"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Input().
						Class("pf-c-form-control").
						Type("email").
						ID("textInput6").
						Name("textInput6").
						Aria("label", "Email input field").
						Aria("describedby", "email-example"),
					app.Span().
						Class("pf-c-input-group__text").
						ID("email-example").
						Body(
							app.Text("@example.com"),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Span().
						Class("pf-c-input-group__text").
						Body(
							app.I().
								Class("fas fa-at").
								Aria("hidden", true),
						),
					app.Input().
						Class("pf-c-form-control").
						Required(true).
						Type("email").
						ID("textInput7").
						Name("textInput7").
						Aria("invalid", true).
						Aria("label", "Error state username example"),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Input().
						Class("pf-c-form-control").
						Type("text").
						ID("textInput13").
						Name("textInput13").
						Aria("label", "Input example with popover"),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("button").
						Aria("label", "Popover for input").
						Body(
							app.I().
								Class("fas fa-question-circle").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Input().
						Class("pf-c-form-control").
						Type("text").
						ID("textInput12").
						Name("textInput12").
						Aria("label", "Input example with popover"),
					app.Button().
						Class("pf-c-button pf-m-plain").
						Type("button").
						Aria("label", "Popover for input").
						Body(
							app.I().
								Class("fas fa-question-circle").
								Aria("hidden", true),
						),
				),
			app.Br(),
			app.Div().
				Class("pf-c-input-group").
				Body(
					app.Input().
						Class("pf-c-form-control").
						Type("number").
						ID("textInput14").
						Name("textInput14").
						Aria("label", "Input example with plain unit"),
					app.Span().
						Class("pf-c-input-group__text pf-m-plain").
						Body(
							app.Text("%"),
						),
				),
		)
}
