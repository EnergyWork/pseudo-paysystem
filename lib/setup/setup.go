package setup

import (
	"fmt"

	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/domain"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type Setup struct {
	cfg *config.Config
	db  *gorm.DB
	nc  *nats.Conn
	log domain.Logger
}

func New(cfg *config.Config, logger domain.Logger) *Setup {
	return &Setup{
		cfg: cfg,
		log: logger,
	}
}

// ConnectNATS returns connection to nats
func (s *Setup) ConnectNATS(opts ...nats.Option) *errs.Error {
	if s.nc == nil {
		if s.cfg == nil {
			return errs.ErrInternal.SetMsg("config not initialized")
		}

		url := fmt.Sprintf("nats://%s:%s", s.cfg.NATS.Host, s.cfg.NATS.Port)

		nc, err := nats.Connect(url, opts...)
		if err != nil {
			return errs.ErrInternal.SetMsg("unable to connect to NATS: %s", err)
		}
		s.nc = nc
	}
	return nil
}

// ConnectPostgres returns connection to postgres server
func (s *Setup) ConnectPostgres() *errs.Error {
	return nil
}
