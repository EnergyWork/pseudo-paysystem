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
	api.Header `json:"-"`

	// WalletID is value of wallet number
	WalletID string `json:"wallet_id"`

	// Value of the balance when creating
	Value int `json:"value"`
}

// RplBalanceCreate the struct representation response body
//
//swagger:response rpl_balance_create
type RplBalanceCreate struct {
	api.Header `json:"-"`
}

// CreateBalance : balance creation handler
func (s *UseCase) CreateBalance(req ReqBalanceCreate) (api.Reply, *errs.Error) {

	//

	return nil, errs.ErrInternal.SetMsg("test message")
}
