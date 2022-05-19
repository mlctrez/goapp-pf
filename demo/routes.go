package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func AddRoutes() {
	app.Route("/", &Page{})
}
