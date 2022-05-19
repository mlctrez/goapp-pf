package button

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

func Demo() app.UI {
	var body []app.UI
	body = append(body,
		app.Style().Text(`.ws-core-c-button .pf-c-button { margin-right: 6px; margin-bottom: 6px; }`),
	)
	section := func(header string, elem ...ui.UI) {
		body = append(body, app.H3().Text(header))
		for _, u := range elem {
			body = append(body, u.UI())
		}
	}

	text := func(t string) []app.UI { return ui.S(app.Text(t)) }

	section("Variations",
		&Button{Variant: Primary, Children: text("Primary")},
		&Button{Variant: Secondary, Children: text("Secondary")},
		&Button{Variant: Secondary, Danger: true, Children: text("Danger secondary")},
		&Button{Variant: Tertiary, Children: text("Tertiary")},
		&Button{Variant: Danger, Children: text("Danger")},
		&Button{Variant: Warning, Children: text("Warning")},
		ui.W(app.Br()), ui.W(app.Br()),
		&Button{Variant: Link, Icon: fa.Icon("plus-circle"), Children: text("Link")},
		&Button{Variant: Link, Icon: fa.Icon("plus-circle"), Children: text("Link"), IconPosition: "right"},
		&Button{Variant: Link, Icon: fa.Icon("plus-circle"), Children: text("Link danger"), Danger: true},
		&Button{Variant: Link, Children: text("Inline link"), Inline: true},
		ui.W(app.Br()), ui.W(app.Br()),
		&Button{Variant: PLain, AriaLabel: "Remove", Children: ui.S(fa.Icon("times"))},
		ui.W(app.Br()), ui.W(app.Br()),
		&Button{Variant: Control, Children: text("Control")},
		&Button{Variant: Control, Children: ui.S(fa.Icon("copy"))},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Disabled",
		&Button{Variant: Primary, Children: text("Primary disabled"), Disabled: true},
		&Button{Variant: Secondary, Children: text("Secondary disabled"), Disabled: true},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)
	section("Aria-disabled",
		&Button{Variant: Primary, Children: text("Primary disabled"), AriaDisabled: true},
		&Button{Variant: Secondary, Children: text("Secondary disabled"), AriaDisabled: true},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Links as buttons",
		&Button{Variant: Primary, Component: "a",
			Children: text("Link to core docs"), Href: "https://pf4.patternfly.org/", Target: "_blank"},
		&Button{Variant: Secondary, Component: "a",
			Children: text("Secondary link to core docs"), Href: "https://pf4.patternfly.org/", Target: "_blank"},
		&Button{Variant: Tertiary, Component: "a",
			Children: text("Tertiary link to core docs"), Href: "https://pf4.patternfly.org/", Target: "_blank"},
		ui.W(app.Br()),
		&Button{Variant: Link, Component: "a",
			Children: text("Jump to modifiers in contribution guidelines"),
			Href:     "https://pf4.patternfly.org/contribution/#modifiers", Target: "_blank"},

		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	sb := &Button{Variant: Link, Inline: true, Component: "span",
		Children: text(" This is long button text that needs to be a span so " +
			"that it will wrap inline with the text around it. ")}

	section("Inline link as span",
		ui.W(app.P().Body(
			app.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), sb.UI(),
			app.Text("Sed hendrerit nisi in cursus maximus. Ut malesuada nisi "+
				"turpis, in condimentum velit elementum non."))),
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Block level",
		&Button{Variant: Primary, Block: true, Children: text("Block level button")},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Types",
		&Button{Variant: Primary, Children: text("Submit"), Type: TypeSubmit},
		&Button{Variant: Primary, Children: text("Reset"), Type: TypeReset},
		&Button{Variant: Primary, Children: text("Default"), Type: TypeButton},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Small",
		&Button{Variant: Primary, Children: text("Primary"), Small: true},
		&Button{Variant: Secondary, Children: text("Secondary"), Small: true},
		&Button{Variant: Tertiary, Children: text("Tertiary"), Small: true},
		&Button{Variant: Danger, Children: text("Danger"), Small: true},
		&Button{Variant: Warning, Children: text("Warning"), Small: true},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Call to action",
		&Button{Variant: Primary, Children: text("Call to action"), Large: true},
		&Button{Variant: Secondary, Children: text("Call to action"), Large: true},
		&Button{Variant: Tertiary, Children: text("Call to action"), Large: true},
		&Button{Variant: Link, Children: text("Call to action"),
			Icon: fa.Icon("arrow-right"), IconPosition: "right", Large: true},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	section("Progress",
		//ui.W(app.If(loadingLogs, pauseLogsButton.UI()).Else(resumeLogsButton.UI())),
		loadingLogsButtonUI(),
		//&Button{Variant: Secondary, Children: text("Click to stop loading"), Loading: true},
		//&Button{Variant: PLain, Children: ui.S(fa.Icon("upload")), Loading: false},
		//&Button{Variant: PLain, Children: ui.S(fa.Icon("upload")), Loading: true},
		ui.W(app.Br()), ui.W(app.Br()), ui.W(app.Br()),
	)

	return app.Div().Class("ws-core-c-button").Body(body...)
}

var loadingLogs *Button

// TODO : clean this up a bit to make it a bit more usable
func loadingLogsButtonUI() ui.UI {
	if loadingLogs == nil {
		loadingLogs = &Button{
			ID: "logsButton", Variant: Primary, Loading: true,
			Children: ui.S(app.Text("Pause loading logs")),
		}
	}
	loadingLogs.OnClick = func(ctx app.Context, e app.Event) {
		loadingLogs.Loading = !loadingLogs.Loading
		txt := "Resume loading logs"
		if loadingLogs.Loading {
			txt = "Pause loading logs"
		}
		loadingLogs.Children = ui.S(app.Text(txt))
		loadingLogs.UpdateState(ctx)
	}
	return loadingLogs
}
