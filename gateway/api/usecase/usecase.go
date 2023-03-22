package usecase

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

var _ API = (*UseCase)(nil) // make sure the useCase implements the interface

type API interface {
	WalletCreate(*ReqGateWalletCreate) (*RplGateWalletCreate, *errs.Error)
	WalletUpdate(api.Request) (api.Reply, *errs.Error)
	WalletGet(api.Request) (api.Reply, *errs.Error)
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
	api.Header
	Phone string
}

type RplGateWalletCreate struct {
	api.Header
	Test string
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

func (u *UseCase) WalletCreate(req *ReqGateWalletCreate) (*RplGateWalletCreate, *errs.Error) {
	rpl := &RplGateWalletCreate{}
	errApi := api.NewNATSRequest(u.set, "issuing/wallet/create", req, &rpl)
	if errApi != nil {
		return nil, errApi
	}
	rpl.Test = "kek"
	return rpl, nil
}

func (u *UseCase) WalletUpdate(req api.Request) (api.Reply, *errs.Error) {
	return nil, errs.ErrInternal
}

func (u *UseCase) WalletGet(req api.Request) (api.Reply, *errs.Error) {
	return nil, errs.ErrInternal
}
