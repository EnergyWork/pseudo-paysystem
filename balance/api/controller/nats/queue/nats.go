package queue

import (
	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

const Queue = "balance-queue"

type Controller struct {
	set *setup.Setup
	uc  usecase.Balance
}

func New(set *setup.Setup, uc usecase.Balance) *Controller {
	return &Controller{
		set: set,
		uc:  uc,
	}
}

func (c *Controller) NewRouter() map[string]func() nats.MsgHandler {
	return newBalanceRoutes(c.uc, c.set)
}
