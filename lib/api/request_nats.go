package api

import (
	"encoding/json"
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

func NewNATSRequest(set *setup.Setup, subj string, req Request, rpl Reply, timeout time.Duration) (Reply, *errs.Error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}
	msg, err := set.NATS().Request(subj, data, timeout)
	if err != nil {
		if set.NATS().LastError() != nil {
			return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", set.NATS().LastError())
		}
		return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	var tmp json.RawMessage
	if err = json.Unmarshal(msg.Data, &tmp); err != nil {
		return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	// var rpl Reply
	if err = json.Unmarshal(tmp, &rpl); err != nil {
		return nil, errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	return rpl, nil
}
