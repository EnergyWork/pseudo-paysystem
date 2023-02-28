package repository

import (
	"time"
)

type Balance struct {
	ID             uint64    `gorm:"column:id"`
	LimitID        uint64    `gorm:"column:limit_id"`
	WalletID       string    `gorm:"column:wallet_id"`
	Value          uint64    `gorm:"column:value"`
	Hold           uint64    `gorm:"column:hold"`
	Identification int       `gorm:"column:identification"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type Limit struct {
	ID             uint64 `gorm:"column:id"`
	Identification int    `gorm:"column:identification"`
	Upper          uint64 `gorm:"column:upper"`
	Lower          uint64 `gorm:"column:lower"`
}
