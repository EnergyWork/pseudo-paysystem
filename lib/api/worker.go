package api

import (
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

// WorkerNATS subscribes to NATS queue
func WorkerNATS(setup *setup.Setup, queue string) *errs.Error {
	// Подписаться на очередь и обрабатывать
	return nil
}
