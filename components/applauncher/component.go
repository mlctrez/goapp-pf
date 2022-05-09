package applauncher

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// TODO: finish after dropdown

// AppLauncher
// documentation https://www.patternfly.org/v4/components/application-launcher
type AppLauncher struct {
	AriaLabel      string
	ClassName      []string
	Direction      string
	Favorites      []string
	FavoritesLabel string
	Disabled       bool
	Grouped        bool
	Open           bool
	Items          []*Item
}

type Item struct{}

func (a *AppLauncher) UI() app.UI {
	return &appLauncher{AppLauncher: *a}
}

type appLauncher struct {
	app.Compo
	AppLauncher
}
