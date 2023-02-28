package queue

import (
	"os"

	"github.com/nats-io/nats.go"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type balanceRoutes struct {
	uc  usecase.Balance
	set *setup.Setup
}

func newBalanceRoutes(uc usecase.Balance, set *setup.Setup) map[string]func() nats.MsgHandler {
	r := balanceRoutes{uc: uc, set: set}
	routes := make(map[string]func() nats.MsgHandler, 0)
	routes["balance/wallet/create"] = r.create // balance/wallet/create
	routes["balance/wallet/hold"] = r.hold
	return routes
}

func (r *balanceRoutes) create() nats.MsgHandler {
	return func(msg *nats.Msg) {
		r.set.Log().Info("Received on [%s] Queue[%s] Pid[%d]: '%s'", msg.Subject, msg.Sub.Queue, os.Getpid(), string(msg.Data))
		_ = msg.Respond([]byte(`{"some_reply"":"some_result"}`))
	}
}

func (r *balanceRoutes) hold() nats.MsgHandler {
	return func(msg *nats.Msg) {}
}
