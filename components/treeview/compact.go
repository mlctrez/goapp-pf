package treeview

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Compact struct {
	app.Compo
}

func (c *Compact) Render() app.UI {
	return app.Div().
		Class("pf-c-tree-view pf-m-compact").
		Body(
			app.Ul().
				Class("pf-c-tree-view__list").
				Aria("role", "tree").
				Body(
					app.Li().
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						TabIndex(-1).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-content").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-title").
																Body(
																	app.Text("apiVersion"),
																),
															app.Div().
																Class("pf-c-tree-view__node-text").
																Body(
																	app.Text("APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value and may reject unrecognized values."),
																),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						TabIndex(-1).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-content").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-title").
																Body(
																	app.Text("kind"),
																),
															app.Div().
																Class("pf-c-tree-view__node-text").
																Body(
																	app.Text("Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated is CamelCase. More info:"),
																),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-tree-view__list-item").
						Aria("role", "treeitem").
						TabIndex(-1).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Div().
										Class("pf-c-tree-view__node").
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-content").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-title").
																Body(
																	app.Text("metadata"),
																),
															app.Div().
																Class("pf-c-tree-view__node-text").
																Body(
																	app.Text("Standard object metadata"),
																),
														),
												),
										),
								),
						),
					app.Li().
						Class("pf-c-tree-view__list-item pf-m-expanded").
						Aria("role", "treeitem").
						Aria("expanded", true).
						TabIndex(0).
						Body(
							app.Div().
								Class("pf-c-tree-view__content").
								Body(
									app.Button().
										Class("pf-c-tree-view__node").
										Body(
											app.Div().
												Class("pf-c-tree-view__node-container").
												Body(
													app.Div().
														Class("pf-c-tree-view__node-toggle").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-angle-right").
																		Aria("hidden", true),
																),
														),
													app.Div().
														Class("pf-c-tree-view__node-content").
														Body(
															app.Span().
																Class("pf-c-tree-view__node-title").
																Body(
																	app.Text("spec"),
																),
															app.Div().
																Class("pf-c-tree-view__node-text").
																Body(
																	app.Text("Specification of the desired behavior of deployment."),
																),
														),
												),
										),
								),
							app.Ul().
								Class("pf-c-tree-view__list").
								Aria("role", "group").
								Body(
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("minReadySeconds"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Default to 0 (pod will be considered available as soon as it is ready)."),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("paused"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("Indicates that the deployment is paused"),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("progressDeadlineSeconds"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that the progress will not de estimated during the time a deployment is paused. Defaults to 600s."),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("replicas"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1."),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("revisionHistoryLimit"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10."),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item pf-m-expanded").
										Aria("role", "treeitem").
										Aria("expanded", true).
										TabIndex(0).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Button().
														Class("pf-c-tree-view__node").
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-toggle").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-toggle-icon").
																				Body(
																					app.I().
																						Class("fas fa-angle-right").
																						Aria("hidden", true),
																				),
																		),
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("Selector"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment"),
																				),
																		),
																),
														),
												),
											app.Ul().
												Class("pf-c-tree-view__list").
												Aria("role", "group").
												Body(
													app.Li().
														Class("pf-c-tree-view__list-item pf-m-expanded").
														Aria("role", "treeitem").
														Aria("expanded", true).
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Button().
																		Class("pf-c-tree-view__node").
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node-toggle").
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-toggle-icon").
																								Body(
																									app.I().
																										Class("fas fa-angle-right").
																										Aria("hidden", true),
																								),
																						),
																					app.Div().
																						Class("pf-c-tree-view__node-content").
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-title").
																								Body(
																									app.Text("matchExpressions"),
																								),
																							app.Span().
																								Class("pf-c-tree-view__node-text").
																								Body(
																									app.Text("matchExpressions is a list of the label selector requirements. The requirements and ANDed."),
																								),
																						),
																				),
																		),
																),
															app.Ul().
																Class("pf-c-tree-view__list").
																Aria("role", "group").
																Body(
																	app.Li().
																		Class("pf-c-tree-view__list-item").
																		Aria("role", "treeitem").
																		TabIndex(-1).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__content").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node").
																						TabIndex(0).
																						Body(
																							app.Div().
																								Class("pf-c-tree-view__node-container").
																								Body(
																									app.Div().
																										Class("pf-c-tree-view__node-content").
																										Body(
																											app.Span().
																												Class("pf-c-tree-view__node-title").
																												Body(
																													app.Text("matchLabels"),
																												),
																											app.Span().
																												Class("pf-c-tree-view__node-text").
																												Body(
																													app.Text("matchExpressions is a list of the label selector requirements. The requirements and ANDed."),
																												),
																										),
																								),
																						),
																				),
																		),
																),
														),
													app.Li().
														Class("pf-c-tree-view__list-item").
														Aria("role", "treeitem").
														TabIndex(-1).
														Body(
															app.Div().
																Class("pf-c-tree-view__content").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node").
																		TabIndex(0).
																		Body(
																			app.Div().
																				Class("pf-c-tree-view__node-container").
																				Body(
																					app.Div().
																						Class("pf-c-tree-view__node-content").
																						Body(
																							app.Span().
																								Class("pf-c-tree-view__node-title").
																								Body(
																									app.Text("matchLabels"),
																								),
																							app.Span().
																								Class("pf-c-tree-view__node-text").
																								Body(
																									app.Text("Map of {key.value} pairs. A single {key.value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\" and the values array contains only \"value\". The requirements are ANDed."),
																								),
																						),
																				),
																		),
																),
														),
												),
										),
									app.Li().
										Class("pf-c-tree-view__list-item").
										Aria("role", "treeitem").
										TabIndex(-1).
										Body(
											app.Div().
												Class("pf-c-tree-view__content").
												Body(
													app.Div().
														Class("pf-c-tree-view__node").
														TabIndex(0).
														Body(
															app.Div().
																Class("pf-c-tree-view__node-container").
																Body(
																	app.Div().
																		Class("pf-c-tree-view__node-content").
																		Body(
																			app.Span().
																				Class("pf-c-tree-view__node-title").
																				Body(
																					app.Text("matchLabels"),
																				),
																			app.Span().
																				Class("pf-c-tree-view__node-text").
																				Body(
																					app.Text("Map of {key.value} pairs. A single {key.value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\" and the values array contains only \"value\". The requirements are ANDed."),
																				),
																		),
																),
														),
												),
										),
								),
						),
				),
		)
}
