package main

import (
	"github.com/caarlos0/env/v7"

	"github.com/energywork/pseudo-paysystem/issuing/internal/app"
	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/logger"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

var (
	Service = "ISSUING"
)

func main() {
	l := logger.New(logger.LoadLogger(true)).WithPrefix(Service)

	// load configuration (env)
	cfg := &config.Config{Service: Service}
	if err := env.Parse(cfg); err != nil {
		l.Fatal(err)
	}

	// forming setup
	set := setup.New(l, cfg)
	if err := set.NewEcho(); err != nil {
		l.Fatal(err)
	}
	/*if err := set.ConnectPostgres(); err != nil {
		l.Fatal(err)
	}*/
	if err := set.ConnectNATS(); err != nil {
		l.Fatal(err)
	}

	// create app
	a, err := app.New(set)
	if err != nil {
		l.Fatal(err)
	}

	// run service
	if err := a.Run(); err != nil {
		l.Fatal(err)
	}
}
