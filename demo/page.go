package demo

import (
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/pfswitch"
	"github.com/mlctrez/goapp-pf/demo/reloader"
)

type Page struct {
	app.Compo
	sync.Once
}

func (r *Page) initComponents() {

}

func (r *Page) Render() app.UI {
	r.Once.Do(r.initComponents)
	return app.Div().Class("pf-c-page pf-m-resize-observer pf-m-breakpoint-2xl").Body(
		&reloader.ReLoader{},
		app.Main().Class("pf-c-page__main").Body(
			app.Section().Class("pf-c-page__main-section pf-m-light").Body(
				pfswitch.Demo(),
			),
		))

}
