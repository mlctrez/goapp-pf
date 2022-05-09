package bullseye

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

var _ ui.UI = (*BullsEye)(nil)

type BullsEye struct {
	Children  []app.UI
	ClassName []string
}

func (b *BullsEye) UI() app.UI {
	return &bullsEye{BullsEye: *b}
}

type bullsEye struct {
	app.Compo
	BullsEye
}

func (b *bullsEye) Render() app.UI {
	div := app.Div().Class("pf-l-bullseye").Class(b.ClassName...)
	return div.Body(b.Children...)
}
