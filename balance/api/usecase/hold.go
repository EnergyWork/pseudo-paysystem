package usecase

import (
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

// swagger:route POST /balance/hold balance-service RequestBalanceHold
// Hold or Unhold value of wallet balance
//
// responses:
//	default:
//	200: rpl_balance_hold
//	500: Error

// ReqBalanceHold the struct representation request body
//
//swagger:parameters RequestBalanceHold
type ReqBalanceHold struct {
	//swagger:ignore
	api.Header

	// WalletID value of wallet number
	//in:body
	WalletID string `json:"wallet_id"`

	// Value of hold
	//in:body
	Value int `json:"value"`

	// Type is hold/unhold
	//in:body
	Type string `json:"type"`
}

// RplBalanceHold the struct representation response body
//
//swagger:response rpl_balance_hold
type RplBalanceHold struct {
	//swagger:ignore
	api.Header
}

// HoldBalance : balance holding handler
func (s *UseCase) HoldBalance(data api.Request) (api.Reply, *errs.Error) {
	req, ok := data.(*ReqBalanceHold)
	if !ok {
		return nil, errs.ErrInternal.SetMsg("fatal error with type casting")
	}
	rpl := &RplBalanceCreate{}
	s.set.Log().Info("type cast: ok")
	s.set.Log().Info("req: %+v", req)
	s.set.Log().Info("balance holed")
	s.set.Log().Info("rpl: %+v", rpl)
	return rpl, nil
}
