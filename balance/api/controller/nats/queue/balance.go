package queue

import (
	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/natsserver"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type balanceRoutes struct {
	uc usecase.Balance
}

func newBalanceRoutes(uc usecase.Balance, set *setup.Setup) map[string]func() nats.MsgHandler {
	r := balanceRoutes{uc: uc}
	natsserver.Register(set.Config().Service, "balance/wallet/create", r.create)
	natsserver.Register(set.Config().Service, "balance/wallet/hold", r.hold)
	return natsserver.GetServiceHandlers(set.Config().Service)
}

func (r *balanceRoutes) create() nats.MsgHandler {
	return natsserver.NatsHandler((*usecase.ReqBalanceCreate)(nil), r.uc.CreateBalance)
}

func (r *balanceRoutes) hold() nats.MsgHandler {
	return natsserver.NatsHandler((*usecase.ReqBalanceHold)(nil), r.uc.HoldBalance)
}
