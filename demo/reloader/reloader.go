package reloader

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var _ app.AppUpdater = (*ReLoader)(nil)

type ReLoader struct {
	app.Compo
}

func (r *ReLoader) Render() app.UI {
	return app.Div().ID("dev-reloader").Hidden(true)
}

func (r *ReLoader) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
