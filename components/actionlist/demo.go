package actionlist

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Demo() app.UI {
	div := app.Div().ID("action-list-demo")

	var body []app.UI
	section := func(name string, list *ActionList) {
		body = append(body, app.Br(), app.H3().Text(name), list.UI())
	}

	section("Single", &ActionList{ID: "action-list-single", Children: []*Group{
		{Children: []*Item{
			{Child: app.Button().Class("pf-c-button pf-m-primary").Text("Next")},
			{Child: app.Button().Class("pf-c-button pf-m-secondary").Text("Back")},
		}},
	}})

	section("Icon", &ActionList{ID: "action-list-icon", IconList: true, Children: []*Group{
		{Children: []*Item{
			{Child: app.Button().Class("pf-c-button pf-m-plain").Body(
				app.I().Class("fas fa-times").Aria("hidden", "true"))},
			{Child: app.Button().Class("pf-c-button pf-m-plain").Body(
				app.I().Class("fas fa-check").Aria("hidden", "true"))},
		}},
	}})

	section("Multiple Groups", &ActionList{ID: "action-list-multiple", Children: []*Group{
		{Children: []*Item{
			{Child: app.Button().Class("pf-c-button pf-m-primary").Text("Next")},
			{Child: app.Button().Class("pf-c-button pf-m-secondary").Text("Back")},
		}},
		{Children: []*Item{
			{Child: app.Button().Class("pf-c-button pf-m-primary").Text("Submit")},
			{Child: app.Button().Class("pf-c-button pf-m-link").Text("Cancel")},
		}},
	}})

	div.Body(body...)

	return div
}
