package repository

import (
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type Repository interface {
	Create(*Wallet) *errs.Error
	Delete(string) *errs.Error
	Update(*Wallet) *errs.Error
}
