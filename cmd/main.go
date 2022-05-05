package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/kardianos/service"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/server"
	"github.com/mlctrez/goapp-pf/ui"
	"github.com/mlctrez/servicego"
)

type svc struct {
	servicego.Defaults
	shutdown func(ctx context.Context) error
}

func main() {

	ui.AddRoutes()

	if app.IsClient {
		app.RunWhenOnBrowser()
	} else {
		servicego.Run(&svc{})
	}

}

func (t *svc) Start(_ service.Service) (err error) {
	t.shutdown, err = server.Run()
	return err
}

func (t *svc) Stop(_ service.Service) (err error) {
	if t.shutdown != nil {

		stopContext, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		err = t.shutdown(stopContext)
		if errors.Is(err, context.Canceled) {
			os.Exit(-1)
		}
	}
	return err
}
