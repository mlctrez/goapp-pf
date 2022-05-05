package menutoggle

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithAvatarAndText struct {
	app.Compo
}

func (c *WithAvatarAndText) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Button().
				Class("pf-c-menu-toggle").
				Type("button").
				Aria("expanded", "false").
				Body(
					app.Span().
						Class("pf-c-menu-toggle__icon").
						Body(
							app.Img().
								Class("pf-c-avatar pf-m-light").
								Src("/assets/images/img_avatar-light.svg").
								Alt("Avatar image light"),
						),
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Ned Username"),
						),
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
			app.Button().
				Class("pf-c-menu-toggle pf-m-expanded").
				Type("button").
				Aria("expanded", true).
				Body(
					app.Span().
						Class("pf-c-menu-toggle__icon").
						Body(
							app.Img().
								Class("pf-c-avatar pf-m-light").
								Src("/assets/images/img_avatar-light.svg").
								Alt("Avatar image light"),
						),
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Ned Username"),
						),
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
			app.Button().
				Class("pf-c-menu-toggle").
				Type("button").
				Aria("expanded", "false").
				Disabled(true).
				Body(
					app.Span().
						Class("pf-c-menu-toggle__icon").
						Body(
							app.Img().
								Class("pf-c-avatar pf-m-light").
								Src("/assets/images/img_avatar-light.svg").
								Alt("Avatar image light"),
						),
					app.Span().
						Class("pf-c-menu-toggle__text").
						Body(
							app.Text("Ned Username"),
						),
					app.Span().
						Class("pf-c-menu-toggle__controls").
						Body(
							app.Span().
								Class("pf-c-menu-toggle__toggle-icon").
								Body(
									app.I().
										Class("fas fa-caret-down").
										Aria("hidden", true),
								),
						),
				),
		)
}
