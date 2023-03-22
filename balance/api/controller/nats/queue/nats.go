package queue

import (
	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/natsserver"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

const Queue = "balance-queue"

// ---------------------------------------------------------------------------------------------------------------------
// Controller

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

type router struct {
	uc usecase.Balance
}

// ---------------------------------------------------------------------------------------------------------------------
// Router

func newBalanceRoutes(uc usecase.Balance, set *setup.Setup) map[string]func() nats.MsgHandler {
	r := router{uc: uc}
	natsserver.Register(set.Config().Service, "balance/wallet/create", r.create)
	natsserver.Register(set.Config().Service, "balance/wallet/hold", r.hold)
	return natsserver.GetServiceHandlers(set.Config().Service)
}

func (r *router) create() nats.MsgHandler {
	return natsserver.NatsHandler((*usecase.ReqBalanceCreate)(nil), r.uc.CreateBalance)
}

func (r *router) hold() nats.MsgHandler {
	return natsserver.NatsHandler((*usecase.ReqBalanceHold)(nil), r.uc.HoldBalance)
}
