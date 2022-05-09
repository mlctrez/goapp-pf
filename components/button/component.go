package button

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Plain(body app.UI) app.HTMLButton {
	return app.Button().Class("pf-c-button pf-m-plain").Type("button").Body(body)
}
