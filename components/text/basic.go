package text

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-content").
		Body(
			app.H1().
				Body(
					app.Text("Hello world"),
				),
			app.P().
				Body(
					app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius\n    lectus, nec rutrum justo nibh eu lectus. Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum\n    mattis neque. Sub works as well!"),
				),
			app.H2().
				Body(
					app.Text("Second level"),
				),
			app.P().
				Body(
					app.Text("Curabitur accumsan turpis pharetra"), app.Strong().
						Body(
							app.Text("augue tincidunt"),
						),
					app.Text("blandit. Quisque condimentum maximus mi, sit amet commodo arcu rutrum id. Proin pretium urna vel\n    cursus venenatis. Suspendisse potenti. Etiam mattis sem rhoncus lacus dapibus facilisis. Donec at dignissim dui. Ut et\n    neque nisl."),
				),
			app.Ul().
				Body(
					app.Li().
						Body(
							app.Text("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
						),
					app.Li().
						Body(
							app.Text("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
						),
					app.Li().
						Body(
							app.Text("Aliquam nec felis in sapien venenatis viverra fermentum nec lectus."), app.Ul().
								Body(
									app.Li().
										Body(
											app.Text("In fermentum leo eu lectus mollis, quis dictum mi aliquet."),
										),
									app.Li().
										Body(
											app.Text("Morbi eu nulla lobortis, lobortis est in, fringilla felis."),
										),
									app.Li().
										Body(
											app.Text("Ut venenatis, nisl scelerisque."), app.Ol().
												Body(
													app.Li().
														Body(
															app.Text("Donec blandit a lorem id convallis."),
														),
													app.Li().
														Body(
															app.Text("Cras gravida arcu at diam gravida gravida."),
														),
													app.Li().
														Body(
															app.Text("Integer in volutpat libero."),
														),
												),
										),
								),
						),
					app.Li().
						Body(
							app.Text("Ut non enim metus."),
						),
				),
			app.H3().
				Body(
					app.Text("Third level"),
				),
			app.P().
				Body(
					app.Text("Quisque ante lacus, malesuada ac auctor vitae, congue"), app.A().
						Href("#").
						Body(
							app.Text("non ante"),
						),
					app.Text(". Phasellus lacus ex, semper ac tortor nec, fringilla condimentum orci. Fusce eu rutrum tellus."),
				),
			app.Ol().
				Body(
					app.Li().
						Body(
							app.Text("Donec blandit a lorem id convallis."),
						),
					app.Li().
						Body(
							app.Text("Cras gravida arcu at diam gravida gravida."),
						),
					app.Li().
						Body(
							app.Text("Integer in volutpat libero."),
						),
					app.Li().
						Body(
							app.Text("Donec a diam tellus."),
						),
					app.Li().
						Body(
							app.Text("Etiam auctor nisl et."), app.Ul().
								Body(
									app.Li().
										Body(
											app.Text("Donec blandit a lorem id convallis."),
										),
									app.Li().
										Body(
											app.Text("Cras gravida arcu at diam gravida gravida."),
										),
									app.Li().
										Body(
											app.Text("Integer in volutpat libero."), app.Ol().
												Body(
													app.Li().
														Body(
															app.Text("Donec blandit a lorem id convallis."),
														),
													app.Li().
														Body(
															app.Text("Cras gravida arcu at diam gravida gravida."),
														),
												),
										),
								),
						),
					app.Li().
						Body(
							app.Text("Aenean nec tortor orci."),
						),
					app.Li().
						Body(
							app.Text("Quisque aliquam cursus urna, non bibendum massa viverra eget."),
						),
					app.Li().
						Body(
							app.Text("Vivamus maximus ultricies pulvinar."),
						),
				),
			app.Blockquote().
				Body(
					app.Text("Ut venenatis, nisl scelerisque sollicitudin fermentum, quam libero hendrerit ipsum, ut blandit est tellus sit amet turpis."),
				),
			app.P().
				Body(
					app.Text("Quisque at semper enim, eu hendrerit odio. Etiam auctor nisl et"), app.Em().
						Body(
							app.Text("justo sodales"),
						),
					app.Text("elementum. Maecenas ultrices lacus quis neque consectetur, et lobortis nisi molestie."),
				),
			app.Hr(),
			app.H3().
				Body(
					app.Text("Visited link example"),
				),
			app.P().
				Body(
					app.A().
						Class("pf-m-visited").
						Href("").
						Body(
							app.Text("Visited link"),
						),
				),
			app.Hr(),
			app.P().
				Body(
					app.Text("Sed sagittis enim ac tortor maximus rutrum. Nulla facilisi. Donec mattis vulputate risus in luctus. Maecenas vestibulum interdum\n    commodo."),
				),
			app.Dl().
				Body(
					app.Dt().
						Body(
							app.Text("Web"),
						),
					app.Dd().
						Body(
							app.Text("The part of the internet that contains websites and web pages"),
						),
					app.Dt().
						Body(
							app.Text("HTML"),
						),
					app.Dd().
						Body(
							app.Text("A markup language for creating web pages"),
						),
					app.Dt().
						Body(
							app.Text("CSS"),
						),
					app.Dd().
						Body(
							app.Text("A technology to make HTML look better"),
						),
				),
			app.P().
				Body(
					app.Text("Suspendisse egestas sapien non felis placerat elementum. Morbi tortor nisl, suscipit sed mi sit amet, mollis malesuada nulla.\n    Nulla facilisi. Nullam ac erat ante."),
				),
			app.H4().
				Body(
					app.Text("Fourth level"),
				),
			app.P().
				Body(
					app.Text("Nulla efficitur eleifend nisi, sit amet bibendum sapien fringilla ac. Mauris euismod metus a tellus laoreet, at elementum\n    ex efficitur."),
				),
			app.P().
				Body(
					app.Text("Maecenas eleifend sollicitudin dui, faucibus sollicitudin augue cursus non. Ut finibus eleifend arcu ut vehicula. Mauris\n    eu est maximus est porta condimentum in eu justo. Nulla id iaculis sapien."),
				),
			app.Small().
				Body(
					app.Text("Sometimes you need small text to display things like date created"),
				),
			app.P().
				Body(
					app.Text("Phasellus porttitor enim id metus volutpat ultricies. Ut nisi nunc, blandit sed dapibus at, vestibulum in felis. Etiam iaculis\n    lorem ac nibh bibendum rhoncus. Nam interdum efficitur ligula sit amet ullamcorper. Etiam tristique, leo vitae porta faucibus,\n    mi lacus laoreet metus, at cursus leo est vel tellus. Sed ac posuere est. Nunc ultricies nunc neque, vitae ultricies ex\n    sodales quis. Aliquam eu nibh in libero accumsan pulvinar. Nullam nec nisl placerat, pretium metus vel, euismod ipsum.\n    Proin tempor cursus nisl vel condimentum. Nam pharetra varius metus non pellentesque."),
				),
			app.H5().
				Body(
					app.Text("Fifth level"),
				),
			app.P().
				Body(
					app.Text("Aliquam sagittis rhoncus vulputate. Cras non luctus sem, sed tincidunt ligula. Vestibulum at nunc elit. Praesent aliquet\n    ligula mi, in luctus elit volutpat porta. Phasellus molestie diam vel nisi sodales, a eleifend augue laoreet. Sed nec eleifend\n    justo. Nam et sollicitudin odio."),
				),
			app.H6().
				Body(
					app.Text("Sixth level"),
				),
			app.P().
				Body(
					app.Text("Cras in nibh lacinia, venenatis nisi et, auctor urna. Donec pulvinar lacus sed diam dignissim, ut eleifend eros accumsan.\n    Phasellus non tortor eros. Ut sed rutrum lacus. Etiam purus nunc, scelerisque quis enim vitae, malesuada ultrices turpis.\n    Nunc vitae maximus purus, nec consectetur dui. Suspendisse euismod, elit vel rutrum commodo, ipsum tortor maximus dui,\n    sed varius sapien odio vitae est. Etiam at cursus metus."),
				),
		)
}
