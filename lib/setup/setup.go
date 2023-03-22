package setup

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/domain"
	"github.com/energywork/pseudo-paysystem/lib/errs"
)

//swagger:ignore
type Setup struct {
	cfg  *config.Config
	db   *gorm.DB
	echo *echo.Echo
	nc   *nats.Conn
	log  domain.Logger
	gin  *gin.Engine
}

func New(logger domain.Logger, cfg *config.Config) *Setup {
	return &Setup{
		log: logger,
		cfg: cfg,
	}
}

func (s *Setup) NewEcho() *errs.Error {
	e := echo.New()
	e.HideBanner = true
	if s.cfg.DEV {
		e.Debug = true
	}
	s.echo = e
	return nil
}

func (s *Setup) NewGin() *errs.Error {
	g := gin.Default()
	s.gin = g
	return nil
}

// ConnectNATS returns connection to nats
func (s *Setup) ConnectNATS(opts ...nats.Option) *errs.Error {
	if s.nc == nil {
		if s.cfg == nil {
			return errs.ErrInternal.SetMsg("config not initialized")
		}

		url := fmt.Sprintf("nats://%s:%s", s.cfg.NatsHost, s.cfg.NatsPort)

		nc, err := nats.Connect(url, opts...)
		if err != nil {
			return errs.ErrInternal.SetMsg("unable to connect to NATS: %s", err)
		}
		s.nc = nc
	}
	s.Log().Info("Connected to NATS")
	return nil
}

// ConnectPostgres returns connection to postgres server
func (s *Setup) ConnectPostgres() *errs.Error {
	db, err := gorm.Open(postgres.Open(s.Config().PostgresDSN()), &gorm.Config{})
	if err != nil {
		return errs.ErrInternal.SetMsg("unable to connect to database")
	}
	s.db = db
	s.Log().Info("Connected to postgres")
	return nil
}
