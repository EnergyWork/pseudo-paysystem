package usecase

import (
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

// swagger:route POST /balance/create balance-service ReqBalanceCreate
// Create a new balance for wallet
//
// responses:
//	default:
//	200: rpl_balance_create
//	500: Error

//swagger:parameters ReqBalanceCreate
type ReqBalanceCreate2 struct {
	//in:body
	RequestBalanceCreate ReqBalanceCreate
}

// ReqBalanceCreate the struct representation request body
//
//swagger:model RequestBalanceCreate
type ReqBalanceCreate struct {
	api.Header

	// WalletID is value of wallet number
	WalletID string `json:"wallet_id"`

	// Value of the balance when creating
	Value int `json:"value"`
}

// RplBalanceCreate the struct representation response body
//
//swagger:response rpl_balance_create
type RplBalanceCreate struct {
	api.Header
}

// CreateBalance : balance creation handler
func (s *UseCase) CreateBalance(data api.Request) (api.Reply, *errs.Error) {
	req, ok := data.(*ReqBalanceCreate)
	if !ok {
		return nil, errs.ErrInternal.SetMsg("fatal error with type casting")
	}
	rpl := &RplBalanceCreate{}
	s.log.Info("type cast: ok")
	s.log.Info("req: %+v", req)
	s.log.Info("balance created")
	s.log.Info("rpl: %+v", rpl)
	return rpl, nil
}
