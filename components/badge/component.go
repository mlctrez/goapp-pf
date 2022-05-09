package badge

import (
	"fmt"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type Badge struct {
	ID        string
	Children  []app.UI
	ClassName []string
	Read      bool
}

func (b *Badge) UI() app.UI {
	return &badge{Badge: *b}
}

type badge struct {
	app.Compo
	Badge

	sync.Once
	read bool
}

func (b *badge) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(b.ID, "read")).Value(&b.read)
}

func (b *badge) Render() app.UI {
	b.Once.Do(func() {
		app.Log("read once", b.ID, b.read, b.Read)
		b.read = b.Read
	})

	app.Log("read     ", b.ID, b.read, b.Read)

	span := app.Span().Class("pf-c-badge")
	if b.read {
		span.Class("pf-m-read")
	} else {
		span.Class("pf-m-unread")
	}

	return span.Class(b.ClassName...).Body(b.Children...)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Badge{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
