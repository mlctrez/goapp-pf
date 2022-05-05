package searchinput

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type AdvancedSearchExpanded struct {
	app.Compo
}

func (c *AdvancedSearchExpanded) Render() app.UI {
	return app.Div().
		Class("pf-c-search-input").
		Body(
			app.Div().
				Class("pf-c-input-group").
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
										Type("text").
										Placeholder("false").
										Aria("label", "username:admin firstname:joe").
										Value("username:root firstname:ned"),
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
					app.Button().
						Class("pf-c-button pf-m-control pf-m-expanded").
						Type("button").
						Aria("expanded", true).
						Aria("label", "Advanced search").
						Body(
							app.I().
								Class("fas fa-caret-down").
								Aria("hidden", true),
						),
					app.Button().
						Class("pf-c-button pf-m-control").
						Type("submit").
						Aria("label", "Search").
						Body(
							app.I().
								Class("fas fa-arrow-right").
								Aria("hidden", true),
						),
				),
			app.Div().
				Class("pf-c-search-input__menu").
				Body(
					app.Div().
						Class("pf-c-search-input__menu-body").
						Body(
							app.Form().
								NoValidate(true).
								Class("pf-c-form").
								Body(
									app.Div().
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-label").
												Body(
													app.Label().
														Class("pf-c-form__label").
														For("advanced-search-input-form-username").
														Body(
															app.Span().
																Class("pf-c-form__label-text").
																Body(
																	app.Text("Username"),
																),
														),
												),
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("root").
														ID("advanced-search-input-form-username").
														Name("advanced-search-input-form-username"),
												),
										),
									app.Div().
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-label").
												Body(
													app.Label().
														Class("pf-c-form__label").
														For("advanced-search-input-form-firstname").
														Body(
															app.Span().
																Class("pf-c-form__label-text").
																Body(
																	app.Text("First name"),
																),
														),
												),
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														Value("ned").
														ID("advanced-search-input-form-firstname").
														Name("advanced-search-input-form-firstname"),
												),
										),
									app.Div().
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-label").
												Body(
													app.Label().
														Class("pf-c-form__label").
														For("advanced-search-input-form-group").
														Body(
															app.Span().
																Class("pf-c-form__label-text").
																Body(
																	app.Text("Group"),
																),
														),
												),
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														ID("advanced-search-input-form-group").
														Name("advanced-search-input-form-group"),
												),
										),
									app.Div().
										Class("pf-c-form__group").
										Body(
											app.Div().
												Class("pf-c-form__group-label").
												Body(
													app.Label().
														Class("pf-c-form__label").
														For("advanced-search-input-form-email").
														Body(
															app.Span().
																Class("pf-c-form__label-text").
																Body(
																	app.Text("Email"),
																),
														),
												),
											app.Div().
												Class("pf-c-form__group-control").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Type("text").
														ID("advanced-search-input-form-email").
														Name("advanced-search-input-form-email"),
												),
										),
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
															app.Text("Submit"),
														),
													app.Button().
														Class("pf-c-button pf-m-link").
														Type("reset").
														Body(
															app.Text("Reset"),
														),
												),
										),
								),
						),
				),
		)
}
