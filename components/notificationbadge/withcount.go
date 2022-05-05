package notificationbadge

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type WithCount struct {
	app.Compo
}

func (c *WithCount) Render() app.UI {
	return app.Div().
		Class("pf-t-dark").
		Body(
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Notifications").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-read").
						Body(
							app.I().
								Class("pf-icon-bell").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("24"),
								),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Unread notifications").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-unread").
						Body(
							app.I().
								Class("pf-icon-bell").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("25"),
								),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Attention notifications").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-attention").
						Body(
							app.I().
								Class("pf-icon-attention-bell").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("26"),
								),
						),
				),
			app.Br(),
			app.Br(),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Tasks").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-read").
						Body(
							app.I().
								Class("pf-icon-task").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("24"),
								),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Unread tasks").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-unread").
						Body(
							app.I().
								Class("pf-icon-task").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("25"),
								),
						),
				),
			app.Button().
				Class("pf-c-button pf-m-plain").
				Type("button").
				Aria("label", "Attention tasks").
				Body(
					app.Span().
						Class("pf-c-notification-badge pf-m-attention").
						Body(
							app.I().
								Class("pf-icon-task").
								Aria("hidden", true),
							app.Span().
								Class("pf-c-notification-badge__count").
								Body(
									app.Text("26"),
								),
						),
				),
		)
}
