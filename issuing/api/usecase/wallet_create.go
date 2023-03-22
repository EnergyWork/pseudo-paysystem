package usecase

import (
	balance "github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/issuing/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
	"github.com/energywork/pseudo-paysystem/lib/types"
)

type ReqIssuingWalletCreate struct {
	api.Header
	Phone string
}

type RplIssuingWalletCreate struct {
	api.Header
	WalletID string
}

// WalletCreate : wallet creation handler
func (s *UseCase) WalletCreate(data api.Request) (api.Reply, *errs.Error) {
	req, ok := data.(*ReqIssuingWalletCreate)
	if !ok {
		return nil, errs.ErrInternal.SetMsg("fatal error with type casting")
	}
	s.set.Log().Info("type cast: ok")
	s.set.Log().Info("req: %+v", req)

	rpl := &RplIssuingWalletCreate{}

	// Create a new wallet
	wallet := &repository.Wallet{
		Phone:  req.Phone,
		Status: types.WalletActive.String(),
		Block:  true,
	}

	/*if err := s.repo.Create(wallet); err != nil {
		s.set.Log().Error("unable to create wallet: %s", err)
		return rpl, err
	}*/
	s.set.Log().Info("типа создали валет в issuing базе данных: %+v", wallet)

	// Request in balance service
	if errApi := req.CreateBalance(s.set, wallet); errApi != nil {
		s.set.Log().Error(errApi)
		s.repo.Delete(wallet.Number)
		return rpl, errApi
	}

	// Unblock the wallet
	wallet.Block = false
	/*if err := s.repo.Update(wallet); err != nil {
		s.set.Log().Error(err)
		return rpl, err
	}*/
	s.set.Log().Info("типа обновили валет в issuing базе данных: %+v", wallet)

	// Form Response
	rpl.WalletID = "TESTWALLETID" // wallet.Number
	s.set.Log().Info("rpl: %+v", rpl)
	return rpl, nil
}

// CreateBalance send request in balance service and create a new row for wallet
func (obj *ReqIssuingWalletCreate) CreateBalance(set *setup.Setup, wallet *repository.Wallet) *errs.Error {
	req := &balance.ReqBalanceCreate{
		WalletID: wallet.Number,
	}
	rpl := &balance.RplBalanceCreate{}
	if errApi := api.NewNATSRequest(set, "balance/wallet/create", req, rpl); errApi != nil { // TODO subject
		return errApi
	}
	return nil
}
