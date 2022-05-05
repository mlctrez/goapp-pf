package fileuploadmultiple

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type FileUploadStatus struct {
	app.Compo
}

func (c *FileUploadStatus) Render() app.UI {
	return app.Div().
		Class("pf-c-multiple-file-upload").
		Body(
			app.Div().
				Class("pf-c-multiple-file-upload__main").
				Body(
					app.Div().
						Class("pf-c-multiple-file-upload__title").
						Body(
							app.Div().
								Class("pf-c-multiple-file-upload__title-icon").
								Body(
									app.I().
										Class("fas fa-upload").
										Aria("hidden", true),
								),
							app.Div().
								Class("pf-c-multiple-file-upload__title-text").
								Body(
									app.Text("Drag and drop files here"), app.Div().
										Class("pf-c-multiple-file-upload__title-text-separator").
										Body(
											app.Text("or"),
										),
								),
						),
					app.Div().
						Class("pf-c-multiple-file-upload__upload").
						Body(
							app.Button().
								Class("pf-c-button pf-m-secondary").
								Type("button").
								Body(
									app.Text("Upload"),
								),
						),
					app.Div().
						Class("pf-c-multiple-file-upload__info").
						Body(
							app.Text("Accepted file types: JPEG, Doc, PDF, PNG"),
						),
				),
			app.Div().
				Class("pf-c-multiple-file-upload__status").
				Body(
					app.Div().
						Class("pf-c-expandable-section").
						Body(
							app.Button().
								Type("button").
								Class("pf-c-expandable-section__toggle").
								Aria("expanded", "false").
								Body(
									app.Span().
										Class("pf-c-expandable-section__toggle-icon").
										Body(
											app.I().
												Class("fas fa-angle-right").
												Aria("hidden", true),
										),
									app.Span().
										Class("pf-c-expandable-section__toggle-text").
										Body(
											app.Div().
												Class("pf-c-multiple-file-upload__status-progress").
												Body(
													app.Div().
														Class("pf-c-multiple-file-upload__status-progress-icon").
														Body(
															app.I().
																Class("pf-icon-in-progress"),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-progress-text").
														Body(
															app.Text("0 of 3 files uploaded"),
														),
												),
										),
								),
							app.Div().
								Class("pf-c-expandable-section__content").
								Hidden(true).
								Body(
									app.Ul().
										Class("pf-c-multiple-file-upload__status-list").
										Body(
											app.Li().
												Class("pf-c-multiple-file-upload__status-item").
												Body(
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-icon").
														Body(
															app.I().
																Class("fas fa-file").
																Aria("hidden", true),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-main").
														Body(
															app.Div().
																Class("pf-c-progress").
																ID("multiple-file-upload-progress-png").
																Body(
																	app.Div().
																		Class("pf-c-progress__description").
																		ID("multiple-file-upload-progress-png-description").
																		Body(
																			app.Span().
																				Class("pf-c-multiple-file-upload__status-item-progress").
																				Body(
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-text").
																						Body(
																							app.Text("filename.png"),
																						),
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-size").
																						Body(
																							app.Text("40 mb"),
																						),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__status").
																		Aria("hidden", true).
																		Body(
																			app.Span().
																				Class("pf-c-progress__measure").
																				Body(
																					app.Text("35%"),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__bar").
																		Aria("role", "progressbar").
																		Aria("valuemin", "0").
																		Aria("valuemax", "100").
																		Aria("valuenow", "35").
																		Aria("labelledby", "multiple-file-upload-progress-png-description").
																		Body(
																			app.Div().
																				Class("pf-c-progress__indicator").
																				Style("width", "35%;"),
																		),
																),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-close").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Remove from list").
																Body(
																	app.I().
																		Class("fas fa-times-circle").
																		Aria("hidden", true),
																),
														),
												),
											app.Li().
												Class("pf-c-multiple-file-upload__status-item").
												Body(
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-icon").
														Body(
															app.I().
																Class("fas fa-file").
																Aria("hidden", true),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-main").
														Body(
															app.Div().
																Class("pf-c-progress").
																ID("multiple-file-upload-progress-doc").
																Body(
																	app.Div().
																		Class("pf-c-progress__description").
																		ID("multiple-file-upload-progress-doc-description").
																		Body(
																			app.Span().
																				Class("pf-c-multiple-file-upload__status-item-progress").
																				Body(
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-text").
																						Body(
																							app.Text("filename.doc"),
																						),
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-size").
																						Body(
																							app.Text("30 mb"),
																						),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__status").
																		Aria("hidden", true).
																		Body(
																			app.Span().
																				Class("pf-c-progress__measure").
																				Body(
																					app.Text("70%"),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__bar").
																		Aria("role", "progressbar").
																		Aria("valuemin", "0").
																		Aria("valuemax", "100").
																		Aria("valuenow", "70").
																		Aria("labelledby", "multiple-file-upload-progress-doc-description").
																		Body(
																			app.Div().
																				Class("pf-c-progress__indicator").
																				Style("width", "70%;"),
																		),
																),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-close").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Remove from list").
																Body(
																	app.I().
																		Class("fas fa-times-circle").
																		Aria("hidden", true),
																),
														),
												),
											app.Li().
												Class("pf-c-multiple-file-upload__status-item").
												Body(
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-icon").
														Body(
															app.I().
																Class("fas fa-file").
																Aria("hidden", true),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-main").
														Body(
															app.Div().
																Class("pf-c-progress").
																ID("multiple-file-upload-progress-pdf").
																Body(
																	app.Div().
																		Class("pf-c-progress__description").
																		ID("multiple-file-upload-progress-pdf-description").
																		Body(
																			app.Span().
																				Class("pf-c-multiple-file-upload__status-item-progress").
																				Body(
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-text").
																						Body(
																							app.Text("filename.pdf"),
																						),
																					app.Span().
																						Class("pf-c-multiple-file-upload__status-item-progress-size").
																						Body(
																							app.Text("25 mb"),
																						),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__status").
																		Aria("hidden", true).
																		Body(
																			app.Span().
																				Class("pf-c-progress__measure").
																				Body(
																					app.Text("50%"),
																				),
																		),
																	app.Div().
																		Class("pf-c-progress__bar").
																		Aria("role", "progressbar").
																		Aria("valuemin", "0").
																		Aria("valuemax", "100").
																		Aria("valuenow", "50").
																		Aria("labelledby", "multiple-file-upload-progress-pdf-description").
																		Body(
																			app.Div().
																				Class("pf-c-progress__indicator").
																				Style("width", "50%;"),
																		),
																),
														),
													app.Div().
														Class("pf-c-multiple-file-upload__status-item-close").
														Body(
															app.Button().
																Class("pf-c-button pf-m-plain").
																Type("button").
																Aria("label", "Remove from list").
																Body(
																	app.I().
																		Class("fas fa-times-circle").
																		Aria("hidden", true),
																),
														),
												),
										),
								),
						),
				),
		)
}
