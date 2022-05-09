package demo

import (
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/aboutmodal"
	"github.com/mlctrez/goapp-pf/components/alertgroup"
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
}

func (r *Page) initComponents() {
	r.about = aboutmodal.Demo()
	r.aboutButton = aboutmodal.DemoOpenButton()
	r.alertGroup = alertgroup.Demo()
	r.alertAdd = alertgroup.DemoAddAlert()
	//r.alertDemo = alert.Demo()

	r.version = &Version{}
}

func (r *Page) Render() app.UI {
	r.Once.Do(r.initComponents)
	return app.Div().Body(
		&reloader.ReLoader{},
		r.about.UI(),
		//r.aboutButton,
		r.alertDemo,
		r.alertGroup.UI(),
		r.alertAdd,
	)
}
