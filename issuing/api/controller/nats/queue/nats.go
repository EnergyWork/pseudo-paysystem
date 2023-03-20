package queue

import (
	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/issuing/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/natsserver"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

const Queue = "issuing-queue"

// ---------------------------------------------------------------------------------------------------------------------
// Controller

type Controller struct {
	set *setup.Setup
	uc  usecase.Issuing
}

func New(set *setup.Setup, uc usecase.Issuing) *Controller {
	return &Controller{
		set: set,
		uc:  uc,
	}
}

func (c *Controller) NewRouter() map[string]func() nats.MsgHandler {
	return newRouter(c.uc, c.set)
}

// ---------------------------------------------------------------------------------------------------------------------
// Router

type router struct {
	uc usecase.Issuing
}

func newRouter(uc usecase.Issuing, set *setup.Setup) map[string]func() nats.MsgHandler {
	r := router{uc: uc}
	natsserver.Register(set.Config().Service, "issuing/wallet/create", r.create)
	return natsserver.GetServiceHandlers(set.Config().Service)
}

func (r *router) create() nats.MsgHandler {
	return natsserver.NatsHandler((*usecase.ReqIssuingWalletCreate)(nil), r.uc.WalletCreate)
}
