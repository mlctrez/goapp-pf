package alertgroup

import (
	"time"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/alert"
	"github.com/mlctrez/goapp-pf/internal/ui"
	"gopkg.in/loremipsum.v1"
)

func Demo() *AlertGroup {
	group := &AlertGroup{
		ID:            "goapp-pf-demo-alert-group",
		Toast:         true,
		OverflowClick: ui.Alert("View X more alerts clicked"),
	}
	return group
}

func DemoAddAlert() app.UI {
	return app.Button().Text("add alert").OnClick(func(ctx app.Context, e app.Event) {
		title := "dragon time " + time.Now().Format(time.RFC3339Nano)

		Demo().Add(ctx, &alert.Alert{
			ID:         uuid.NewString(),
			Variant:    alert.Danger,
			Title:      title,
			Expandable: true,
			ActionLinks: []*alert.ActionLink{
				{Children: "actionLink", OnClick: ui.Alert("clicked on " + title)},
			},
			Children: ui.S(app.Text(loremipsum.New().Sentences(2))),
		})
	})
}
