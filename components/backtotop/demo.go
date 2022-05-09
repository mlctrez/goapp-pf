package backtotop

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"gopkg.in/loremipsum.v1"
)

func DemoPage() []app.UI {
	ipsum := loremipsum.New()

	var gutterBody []app.UI
	for i := 0; i < 80; i++ {
		gutterBody = append(gutterBody,
			app.Div().Body(
				app.Article().Class("pf-c-card").Body(
					app.Div().Class("pf-c-card__body").Text(ipsum.Sentences(4)))))

	}

	bt := &BackToTop{Title: "Back to top", ScrollableSelector: "#scrollable", AlwaysVisible: false}

	return []app.UI{
		app.Section().Class("pf-c-page__main-section pf-m-light").Body(
			app.Div().Class("pf-c-content").Body(app.H1().Text("Page Title"))),
		app.Section().Class("pf-c-page__main-section pf-m-overflow-scroll").ID("scrollable").Body(
			app.Div().Class("pf-l-gallery pf-m-gutter").Body(gutterBody...)),
		bt.UI()}

}
