package api

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

func NewNATSRequest(set *setup.Setup, subj string, data []byte, timeout time.Duration) (interface{}, *errs.Error) {
	rpl, err := set.NATS().Request(subj, data, timeout)
	if err != nil {
		if set.NATS().LastError() != nil {
			return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", set.NATS().LastError())
		}
		return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}
	return rpl, nil
}
