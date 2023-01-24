package setup

import (
	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/domain"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type Setup struct {
	cfg  *config.Config
	db   *gorm.DB
	nats *nats.Conn
	log  domain.Logger
}

func New(cfg *config.Config, logger domain.Logger) *Setup {
	return &Setup{
		cfg: cfg,
		log: logger,
	}
}

// Коннекты к бд, натс
