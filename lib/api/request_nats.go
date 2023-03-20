package api

import (
	"encoding/json"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

func NewNATSRequest(set *setup.Setup, subj string, req Request, rpl Reply) *errs.Error {
	data, err := json.Marshal(req)
	if err != nil {
		return errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}
	msg, err := set.NATS().Request(subj, data, req.Timeout())
	if err != nil {
		if set.NATS().LastError() != nil {
			return errs.New().SetCode(errs.Internal).SetMsg("%v", set.NATS().LastError())
		}
		return errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	var tmp json.RawMessage
	if err = json.Unmarshal(msg.Data, &tmp); err != nil {
		return errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	// var rpl Reply
	if err = json.Unmarshal(tmp, &rpl); err != nil {
		return errs.New().SetCode(errs.Internal).SetMsg("%v", err)
	}

	return nil
}
