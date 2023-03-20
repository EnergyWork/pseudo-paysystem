package repository

import (
	"time"
)

type Wallet struct {
	Number    string    `gorm:"column:number"`                   // wallet number
	Phone     string    `gorm:"column:phone"`                    // wallet phone number
	Status    string    `gorm:"column:status"`                   // wallet status
	Block     bool      `gorm:"column:block"`                    // wallet blocking
	CreatedAt time.Time `gorm:"column:created_at;default:now()"` // wallet creation date
	UpdatedAt time.Time `gorm:"column:updated_at"`               // wallet updating date
}
