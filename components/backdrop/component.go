package backdrop

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

var _ ui.UI = (*BackDrop)(nil)

type BackDrop struct {
	Children  []app.UI
	ClassName []string
}

func (b *BackDrop) UI() app.UI {
	return &backDrop{BackDrop: *b}
}

type backDrop struct {
	app.Compo
	BackDrop
}

func (b *backDrop) Render() app.UI {
	div := app.Div().Class("pf-c-backdrop").Class(b.ClassName...)
	return div.Body(b.Children...)
}
