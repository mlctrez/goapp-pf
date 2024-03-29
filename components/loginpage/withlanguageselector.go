package loginpage

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithLanguageSelector struct {
	app.Compo
}

func (c *WithLanguageSelector) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Div().
				Class("pf-c-background-image").
				Body(
					app.Raw("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"pf-c-background-image__filter\" width=\"0\" height=\"0\">\n      <filter id=\"image_overlay\">\n        <feColorMatrix type=\"matrix\" values=\"1 0 0 0 0\n              1 0 0 0 0\n              1 0 0 0 0\n              0 0 0 1 0\"></feColorMatrix>\n        <feComponentTransfer color-interpolation-filters=\"sRGB\" result=\"duotone\">\n          <feFuncR type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncR>\n          <feFuncG type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncG>\n          <feFuncB type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncB>\n          <feFuncA type=\"table\" tableValues=\"0 1\"></feFuncA>\n        </feComponentTransfer>\n      </filter>\n    </svg>"),
				),
			app.Div().
				Class("pf-c-login").
				Body(
					app.Div().
						Class("pf-c-login__container").
						Body(
							app.Header().
								Class("pf-c-login__header").
								Body(
									app.Img().
										Class("pf-c-brand").
										Src("/assets/images/pf_logo_color.svg").
										Alt("PatternFly Logo"),
								),
							app.Main().
								Class("pf-c-login__main").
								Body(
									app.Header().
										Class("pf-c-login__main-header").
										Body(
											app.H1().
												Class("pf-c-title pf-m-3xl").
												Body(
													app.Text("Log in to your account"),
												),
											app.P().
												Class("pf-c-login__main-header-desc").
												Body(
													app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
												),
											app.Div().
												Class("pf-c-login__main-header-utilities").
												Body(
													app.Div().
														Class("pf-c-select").
														Body(
															app.Span().
																ID("login-select-label").
																Hidden(true).
																Body(
																	app.Text("Choose one"),
																),
															app.Button().
																Class("pf-c-select__toggle").
																Type("button").
																ID("login-select-toggle").
																Aria("haspopup", true).
																Aria("expanded", "false").
																Aria("labelledby", "login-select-label login-select-toggle").
																Body(
																	app.Div().
																		Class("pf-c-select__toggle-wrapper").
																		Body(
																			app.Span().
																				Class("pf-c-select__toggle-text").
																				Body(
																					app.Text("English"),
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
																Aria("labelledby", "login-select-label").
																Hidden(true).
																Body(
																	app.Li().
																		Aria("role", "presentation").
																		Body(
																			app.Button().
																				Class("pf-c-select__menu-item pf-m-selected").
																				Aria("role", "option").
																				Aria("selected", true).
																				Body(
																					app.Text("English"), app.Span().
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
																					app.Text("Español"),
																				),
																		),
																	app.Li().
																		Aria("role", "presentation").
																		Body(
																			app.Button().
																				Class("pf-c-select__menu-item").
																				Aria("role", "option").
																				Body(
																					app.Text("Français"),
																				),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-login__main-body").
										Body(
											app.Form().
												NoValidate(true).
												Class("pf-c-form").
												Body(
													app.P().
														Class("pf-c-form__helper-text pf-m-error pf-m-hidden").
														Body(
															app.Span().
																Class("pf-c-form__helper-text-icon").
																Body(
																	app.I().
																		Class("fas fa-exclamation-circle").
																		Aria("hidden", true),
																),
															app.Text("Invalid login credentials."),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("login-demo-form-username").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Username"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Input().
																Class("pf-c-form-control").
																Required(true).
																Type("text").
																ID("login-demo-form-username").
																Name("login-demo-form-username"),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Label().
																Class("pf-c-form__label").
																For("login-demo-form-password").
																Body(
																	app.Span().
																		Class("pf-c-form__label-text").
																		Body(
																			app.Text("Password"),
																		),
																	app.Span().
																		Class("pf-c-form__label-required").
																		Aria("hidden", true).
																		Body(
																			app.Text("*"),
																		),
																),
															app.Input().
																Class("pf-c-form-control").
																Required(true).
																Type("password").
																ID("login-demo-form-password").
																Name("login-demo-form-password"),
														),
													app.Div().
														Class("pf-c-form__group").
														Body(
															app.Div().
																Class("pf-c-check").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("login-demo-checkbox").
																		Name("login-demo-checkbox"),
																	app.Label().
																		Class("pf-c-check__label").
																		For("login-demo-checkbox").
																		Body(
																			app.Text("Keep me logged in for 30 days."),
																		),
																),
														),
													app.Div().
														Class("pf-c-form__group pf-m-action").
														Body(
															app.Button().
																Class("pf-c-button pf-m-primary pf-m-block").
																Type("submit").
																Body(
																	app.Text("Log in"),
																),
														),
												),
										),
									app.Footer().
										Class("pf-c-login__main-footer").
										Body(
											app.Ul().
												Class("pf-c-login__main-footer-links").
												Body(
													app.Li().
														Class("pf-c-login__main-footer-links-item").
														Body(
															app.A().
																Href("#").
																Class("pf-c-login__main-footer-links-item-link").
																Aria("label", "Log in with Google").
																Body(
																	app.Raw("<svg aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 488 512\">\n                  <path d=\"M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z\"></path>\n                </svg>"),
																),
														),
													app.Li().
														Class("pf-c-login__main-footer-links-item").
														Body(
															app.A().
																Href("#").
																Class("pf-c-login__main-footer-links-item-link").
																Aria("label", "Log in with Github").
																Body(
																	app.Raw("<svg aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 496 512\">\n                  <path d=\"M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z\"></path>\n                </svg>"),
																),
														),
													app.Li().
														Class("pf-c-login__main-footer-links-item").
														Body(
															app.A().
																Href("#").
																Class("pf-c-login__main-footer-links-item-link").
																Aria("label", "Log in with Dropbox").
																Body(
																	app.Raw("<svg aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 528 512\">\n                  <path d=\"M264.4 116.3l-132 84.3 132 84.3-132 84.3L0 284.1l132.3-84.3L0 116.3 132.3 32l132.1 84.3zM131.6 395.7l132-84.3 132 84.3-132 84.3-132-84.3zm132.8-111.6l132-84.3-132-83.6L395.7 32 528 116.3l-132.3 84.3L528 284.8l-132.3 84.3-131.3-85z\"></path>\n                </svg>"),
																),
														),
													app.Li().
														Class("pf-c-login__main-footer-links-item").
														Body(
															app.A().
																Href("#").
																Class("pf-c-login__main-footer-links-item-link").
																Aria("label", "Log in with Facebook").
																Body(
																	app.Raw("<svg aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 448 512\">\n                  <path d=\"M448 56.7v398.5c0 13.7-11.1 24.7-24.7 24.7H309.1V306.5h58.2l8.7-67.6h-67v-43.2c0-19.6 5.4-32.9 33.5-32.9h35.8v-60.5c-6.2-.8-27.4-2.7-52.2-2.7-51.6 0-87 31.5-87 89.4v49.9h-58.4v67.6h58.4V480H24.7C11.1 480 0 468.9 0 455.3V56.7C0 43.1 11.1 32 24.7 32h398.5c13.7 0 24.8 11.1 24.8 24.7z\"></path>\n                </svg>"),
																),
														),
													app.Li().
														Class("pf-c-login__main-footer-links-item").
														Body(
															app.A().
																Href("#").
																Class("pf-c-login__main-footer-links-item-link").
																Aria("label", "Log in with Gitlab").
																Body(
																	app.Raw("<svg aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\">\n                  <path d=\"M29.782 199.732L256 493.714 8.074 309.699c-6.856-5.142-9.712-13.996-7.141-21.993l28.849-87.974zm75.405-174.806c-3.142-8.854-15.709-8.854-18.851 0L29.782 199.732h131.961L105.187 24.926zm56.556 174.806L256 493.714l94.257-293.982H161.743zm349.324 87.974l-28.849-87.974L256 493.714l247.926-184.015c6.855-5.142 9.711-13.996 7.141-21.993zm-85.404-262.78c-3.142-8.854-15.709-8.854-18.851 0l-56.555 174.806h131.961L425.663 24.926z\"></path>\n                </svg>"),
																),
														),
												),
											app.Div().
												Class("pf-c-login__main-footer-band").
												Body(
													app.P().
														Class("pf-c-login__main-footer-band-item").
														Body(
															app.Text("Need an account?"), app.A().
																Href("https://www.patternfly.org/").
																Body(
																	app.Text("Sign up."),
																),
														),
													app.P().
														Class("pf-c-login__main-footer-band-item").
														Body(
															app.A().
																Href("#").
																Body(
																	app.Text("Forgot username or password?"),
																),
														),
												),
										),
								),
							app.Footer().
								Class("pf-c-login__footer").
								Body(
									app.P().
										Body(
											app.Text("This is placeholder text only. Use this area to place any information or introductory message about your application that may be relevant to users."),
										),
									app.Ul().
										Class("pf-c-list pf-m-inline").
										Body(
											app.Li().
												Body(
													app.A().
														Href("#").
														Body(
															app.Text("Terms of use"),
														),
												),
											app.Li().
												Body(
													app.A().
														Href("#").
														Body(
															app.Text("Help"),
														),
												),
											app.Li().
												Body(
													app.A().
														Href("#").
														Body(
															app.Text("Privacy policy"),
														),
												),
										),
								),
						),
				),
		)
}
