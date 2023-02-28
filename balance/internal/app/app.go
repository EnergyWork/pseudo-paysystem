package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/energywork/pseudo-paysystem/balance/api/controller/nats/queue"
	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/balance/internal/repository/pg"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/natsserver"
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
	const BalanceQueue = "balance-queue"

	// Repository
	repository := pg.New(a.set.GORM())

	// Use case
	useCase := usecase.New(repository, a.set.Log())

	// NATS Server
	natsController := queue.New(a.set, useCase)
	natsServer, err := natsserver.New(a.set.NATS(), BalanceQueue, natsController.NewRouter()) // this starts a server
	if err != nil {
		a.set.Log().Error(err)
	}
	a.set.Log().Info("The nats server is listening to an queue: %s", BalanceQueue)

	// Waiting signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-signalChan:
		a.set.Log().Info("Run - signal: %s", s.String())
	case err := <-natsServer.Notify():
		a.set.Log().Info("Run - natsServer: %s", err)
	}

	// Shutdown NATS Server
	if err := natsServer.Shutdown(); err != nil {
		a.set.Log().Error(err)
	}

	return nil
}
