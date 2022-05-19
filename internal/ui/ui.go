package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type UI interface {
	UI() app.UI
}

func ToAppUI(elem ...UI) (appUI []app.UI) {
	for _, ui := range elem {
		appUI = append(appUI, ui.UI())
	}
	return appUI
}

func S(elem app.UI) []app.UI {
	return []app.UI{elem}
}

func W(elem app.UI) UI { return &w{ui: elem} }

type w struct{ ui app.UI }

func (w *w) UI() app.UI { return w.ui }
