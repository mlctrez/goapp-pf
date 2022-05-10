package banner

import (
	"fmt"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type Banner struct {
	ID        string
	Children  []app.UI
	ClassName []string
	Sticky    bool
	Variant   Variant
}

func (b *Banner) UI() app.UI {
	return &banner{Banner: *b}
}

type banner struct {
	app.Compo
	Banner

	sync.Once
}

func (b *banner) OnMount(ctx app.Context) {
}

func (b *banner) Render() app.UI {
	b.Once.Do(func() {
		b.Variant = b.Variant.checkVariant()
	})
	div := app.Div().ID(b.ID)
	b.Variant.class(div)
	if b.Sticky {
		div.Class("pf-m-sticky")
	}
	div.Class(b.ClassName...)
	return div.Body(b.Children...)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Banner{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
