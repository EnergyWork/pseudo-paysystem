package usecase

import (
	"github.com/energywork/pseudo-paysystem/balance/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type Balance interface {
	CreateBalance(api.Request) (api.Reply, *errs.Error)
	HoldBalance(api.Request) (api.Reply, *errs.Error)
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
