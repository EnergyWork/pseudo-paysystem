package usecase

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

var _ API = (*UseCase)(nil) // make sure the useCase implements the interface

type API interface {
	WalletCreate(req api.Request) (api.Reply, *errs.Error)
	WalletUpdate(req api.Request) (api.Reply, *errs.Error)
	WalletGet(req api.Request) (api.Reply, *errs.Error)
}

type UseCase struct {
	set *setup.Setup
}

func New(set *setup.Setup) *UseCase {
	return &UseCase{
		set: set,
	}
}

type ReqGateWalletCreate struct {
	Phone string
	// todo : more data
}

func (obj *ReqGateWalletCreate) Timeout() time.Duration {
	return 5 * time.Second
}

func (obj *ReqGateWalletCreate) Validate() *errs.Error {
	if obj.Phone == "" {
		return errs.New().SetCode(errs.Syntax).SetMsg("Phone must be not empty")
	}
	return nil
}

func (obj *ReqGateWalletCreate) Authorize() *errs.Error {
	// todo
	return nil
}

func (u *UseCase) WalletCreate(req api.Request) (api.Reply, *errs.Error) {
	var rpl api.Reply
	errApi := api.NewNATSRequest(u.set, "", req, rpl)
	if errApi != nil {
		return nil, errApi
	}
	return rpl, nil
}

func (u *UseCase) WalletUpdate(req api.Request) (api.Reply, *errs.Error) {
	return nil, errs.ErrInternal
}

func (u *UseCase) WalletGet(req api.Request) (api.Reply, *errs.Error) {
	return nil, errs.ErrInternal
}
