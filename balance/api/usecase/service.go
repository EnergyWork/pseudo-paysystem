package usecase

import (
	"github.com/energywork/pseudo-paysystem/balance/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/api"
	"github.com/energywork/pseudo-paysystem/lib/domain"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type Balance interface {
	CreateBalance(ReqBalanceCreate) (api.Reply, *errs.Error)
	HoldBalance(ReqBalanceHold) (api.Reply, *errs.Error)
}

type UseCase struct { // фактически UseCase
	repo repository.Repository
	log  domain.Logger
	// and another modules, e.g. Cache (Redis, MemCache)
	// or some http client for requests to an external system
}

func New(r repository.Repository, l domain.Logger) *UseCase {
	s := &UseCase{
		repo: r,
		log:  l,
	}
	return s
}
