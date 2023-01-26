package test

import (
	"testing"

	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/logger"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

func GetSetup(t *testing.T) *setup.Setup {
	set := setup.New(&config.Config{
		API: config.API{},
		NATS: config.NATS{
			Host: "localhost",
			Port: "4222",
		},
		SQL: config.SQL{},
		DEV: true,
	}, logger.New(logger.LoadLogger(true)))

	if err := set.ConnectNATS(); err != nil {
		t.Fatal(err)
	}

	return set
}
