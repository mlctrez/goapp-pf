package ui

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Wrap(ui app.UI) app.UI {
	return &wrapper{ui: ui}
}

type wrapper struct {
	app.Compo
	ui app.UI
}

func (w *wrapper) Render() app.UI {
	return w.ui
}
