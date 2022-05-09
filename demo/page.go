package demo

import (
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/aboutmodal"
	"github.com/mlctrez/goapp-pf/components/alertgroup"
	"github.com/mlctrez/goapp-pf/components/avatar"
	"github.com/mlctrez/goapp-pf/components/badge"
	"github.com/mlctrez/goapp-pf/reloader"
)

type Page struct {
	app.Compo
	sync.Once
	version     *Version
	about       *aboutmodal.AboutModal
	alertGroup  *alertgroup.AlertGroup
	aboutButton app.UI
	alertDemo   app.UI
	alertAdd    app.UI
	avatarDemo  app.UI
}

func (r *Page) initComponents() {
	r.about = aboutmodal.Demo()
	r.aboutButton = aboutmodal.DemoOpenButton()
	r.alertGroup = alertgroup.Demo()
	r.alertAdd = alertgroup.DemoAddAlert()
	r.avatarDemo = avatar.Demo()
	//r.alertDemo = alert.Demo()

	r.version = &Version{}
}

func (r *Page) Render() app.UI {
	r.Once.Do(r.initComponents)
	return app.Div().Class("pf-c-page pf-m-resize-observer pf-m-breakpoint-2xl").Body(
		&reloader.ReLoader{},
		r.alertGroup.UI(),
		app.Main().Class("pf-c-page__main").Body(
			app.Section().Class("pf-c-page__main-section pf-m-light").Body(
				badge.Demo(),
				r.alertAdd,
			),

			//(&backgroundimage.BackgroundImage{}).UI(),
		))

}
