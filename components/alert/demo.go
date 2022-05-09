package alert

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/ui"
	"gopkg.in/loremipsum.v1"
)

func Demo() app.UI {

	ipsum := loremipsum.New()

	var body []app.UI

	for _, variant := range Variants() {
		a := &Alert{Variant: variant, Title: variant.toTitle() + " title",
			Children: ui.S(app.Text(ipsum.Sentence()))}
		body = append(body, a.UI())
	}

	for _, variant := range Variants() {
		a := &Alert{Inline: true, Variant: variant, Title: variant.toTitle() + " inline"}
		body = append(body, a.UI())
	}

	for _, variant := range Variants() {
		a := &Alert{Inline: true, Plain: true, Variant: variant, Title: variant.toTitle() + " plain"}
		body = append(body, a.UI())
	}

	body = append(body, withActionLinks().UI())
	body = append(body, withActionCloseButton().UI())
	body = append(body, withTruncateTitle().UI())
	body = append(body, withCustomIcon().UI())
	body = append(body, withExpandable().UI())

	return app.Div().ID("goapp-pf-demo-alert").Body(body...)
}

func withActionLinks() ui.UI {
	return &Alert{
		Variant:  Success,
		Title:    "with action links",
		Children: ui.S(app.Text(loremipsum.New().Sentence())),
		ActionLinks: []*ActionLink{
			{Children: "View Details", OnClick: ui.Alert("View Details clicked")},
			{Children: "Ignore", OnClick: ui.Alert("Ignore clicked")},
		},
	}
}

func withActionCloseButton() ui.UI {
	return &Alert{
		ActionClose: &ActionCloseButton{OnClose: ui.Alert("You clicked close")},
		Variant:     Success,
		Title:       "With action close button",
		Children:    ui.S(app.Text(loremipsum.New().Sentence())),
		ActionLinks: []*ActionLink{
			{Children: "View Details", OnClick: ui.Alert("View Details clicked")},
			{Children: "Ignore", OnClick: ui.Alert("Ignore clicked")},
		},
	}
}

func withTruncateTitle() ui.UI {
	return &Alert{
		TruncateTitle: 2,
		Variant:       Danger,
		Title:         "With truncate title : " + loremipsum.New().Sentences(5),
	}
}

func withCustomIcon() ui.UI {
	return &Alert{
		CustomIcon: fontawesome.Icon("database").Class("fa-fw"),
		Variant:    Success,
		Title:      "With custom icon",
	}
}

func withExpandable() ui.UI {
	return &Alert{
		Expandable: true,
		Variant:    Success,
		Title:      "With expandable",
		Children:   ui.S(app.Text(loremipsum.New().Sentence())),
		ActionLinks: []*ActionLink{
			{Children: "View Details", OnClick: ui.Alert("View Details clicked")},
			{Children: "Ignore", OnClick: ui.Alert("Ignore clicked")},
		},
	}
}
