package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/demo"
)

func AddRoutes() {
	app.Route("/", &demo.Page{})
}
