package main

import (
	"os"
	"os/signal"
	"syscall"

	balance "github.com/energywork/pseudo-paysystem/balance/api"
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/logger"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

var (
	Service = "BALANCE"
)

func main() {
	configPath, ok := os.LookupEnv("CFG_PATH")
	if !ok {
		configPath = "./config.yml"
	}

	l := logger.New(logger.LoadLogger(true)).WithPrefix(Service)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		l.Fatal(err)
	}

	set := setup.New(cfg, l)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	if err := api.WorkerNATS(set, balance.Queue); err != nil {
		l.Error(err)
	}
}
