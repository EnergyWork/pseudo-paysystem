package app

import (
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type App struct {
	set *setup.Setup
}

func New(set *setup.Setup) (*App, *errs.Error) {
	a := &App{
		set: set, // there are already have connections for NATS, PG and initialized 'echo' object
	}
	return a, nil // and "error" for the future
}

func (a *App) Run() *errs.Error {

	return nil
}

// HTTP Server
/*c := v1.New(a.set, useCase).ConfigureRouter()

go func() {
	if err := c.SaveRoutes(); err != nil {
		a.set.Log().Error(err)
	} else {
		a.set.Log().Info("routes saved (routes.json)")
	}
}()

httpServer := httpserver.NewHTTPServer( // this starts the server
	a.set.Echo(),
	httpserver.Addr(a.set.Config().HttpHost, a.set.Config().HttpPort),
	httpserver.ReadTimeout(time.Duration(a.set.Config().ReadTimeout)),
	httpserver.WriteTimeout(time.Duration(a.set.Config().WriteTimeout)),
	httpserver.ShutdownTimeout(time.Duration(a.set.Config().ShutdownTimeout)),
)

a.set.Log().Info("The http server is listening to an address: %s", httpServer.Addr())*/

/*case err := <-httpServer.Notify():
a.set.Log().Info("Run - httpServer: %s", err)*/

// Shutdown HTTP Server
/*if err := httpServer.Shutdown(); err != nil {
	return nil
}*/
