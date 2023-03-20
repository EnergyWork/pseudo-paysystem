package usecase

import (
	"github.com/energywork/pseudo-paysystem/issuing/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

var _ Issuing = (*UseCase)(nil) // implementation checker

type Issuing interface {
	WalletCreate(api.Request) (api.Reply, *errs.Error)
	// todo two more methods
	// Update(api.Request) (api.Reply, *errs.Error)
	// Get(api.Request) (api.Reply, *errs.Error)
}

type UseCase struct {
	repo repository.Repository
	set  *setup.Setup
	// and another modules, e.g. Cache (Redis, MemCache)
	// or some http client for requests to an external system
}

func New(r repository.Repository, s *setup.Setup) *UseCase {
	uc := &UseCase{
		repo: r,
		set:  s,
	}
	return uc
}
