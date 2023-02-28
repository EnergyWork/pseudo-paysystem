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
	api.Header `json:"-"`

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
	api.Header `json:"-"`
}

// HoldBalance : balance holding handler
func (s *UseCase) HoldBalance(req ReqBalanceHold) (api.Reply, *errs.Error) {
	s.log.Info("Request: %+v", req)
	s.log.Info("some actions")

	//

	return nil, errs.ErrInternal.SetMsg("test message")
}
