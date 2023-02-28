package pg

import (
	"gorm.io/gorm"

	"github.com/energywork/pseudo-paysystem/balance/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type BalanceRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BalanceRepository {
	return &BalanceRepository{
		db: db,
	}
}

// Create creates a new wallet balance
func (obj *BalanceRepository) Create(balance *repository.Balance) *errs.Error {
	err := obj.db.Model(&balance).Create(&balance).Error
	if err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}
	return nil
}
