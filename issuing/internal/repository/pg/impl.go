package pg

import (
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/energywork/pseudo-paysystem/issuing/internal/repository"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

var _ repository.Repository = (*Repository)(nil) // implementation checker

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create creates a new wallet
func (r *Repository) Create(wallet *repository.Wallet) *errs.Error {
	wallet.Number = "80" + strconv.FormatInt(time.Now().Unix(), 10)
	sqlStr := "INSERT INTO wallet (number, phone, status, block) VALUES ($1, $2, $3, $4)"
	err := r.db.Raw(sqlStr, wallet.Number, wallet.Phone, wallet.Status, wallet.Block).Error
	if err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}
	return nil
}

// Update updates a wallet
func (r *Repository) Update(wallet *repository.Wallet) *errs.Error {
	err := r.db.Table("wallet").Updates(&wallet).Error
	if err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}
	return nil
}

// Delete deletes a wallet
func (r *Repository) Delete(walletID string) *errs.Error {
	sqlStr := "DELETE FROM wallet WHERE number=$1"
	tx := r.db.Begin()
	err := tx.Raw(sqlStr, walletID).Error
	if err != nil {
		tx.Rollback()
		return errs.ErrInternal.SetMsg(err.Error())
	}
	tx.Commit()
	return nil
}
