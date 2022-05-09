package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/server"
)

type Version struct {
}

func (v *Version) UI() app.UI {

	styles := map[string]string{
		"font-size":        "small",
		"position":         "absolute",
		"bottom":           "0",
		"left":             "1em",
		"height":           "1.5em",
		"background-color": "#222",
	}

	div := app.Div().Styles(styles).
		Text("Version : " + server.GoAppVersion())

	return div
}
