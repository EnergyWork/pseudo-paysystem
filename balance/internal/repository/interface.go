package repository

import (
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

//go:generate mockery --name Repository
type Repository interface {
	Create(*Balance) *errs.Error
}
