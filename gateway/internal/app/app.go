package app

import (
	"time"

	"github.com/energywork/pseudo-paysystem/gateway/api/controller/http/v1/v1gin"
	"github.com/energywork/pseudo-paysystem/gateway/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/httpserver"
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

	useCase := usecase.New(a.set)
	_ = v1gin.New(a.set, useCase).ConfigureRouter()

	httpServer := httpserver.New( // this starts the server
		a.set.Gin(),
		httpserver.Addr(a.set.Config().HttpHost, a.set.Config().HttpPort),
		httpserver.ReadTimeout(time.Duration(a.set.Config().ReadTimeout)),
		httpserver.WriteTimeout(time.Duration(a.set.Config().WriteTimeout)),
		httpserver.ShutdownTimeout(time.Duration(a.set.Config().ShutdownTimeout)),
	)

	// a.set.Log().Info("The http server is listening to an address: %s", httpServer.Addr())

	if err := <-httpServer.Notify(); err != nil {
		a.set.Log().Info("Run - httpServer: %s", err)
	}

	// Shutdown HTTP Server
	if err := httpServer.Shutdown(); err != nil {
		return nil
	}

	return nil
}
