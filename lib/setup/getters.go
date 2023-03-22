package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"

	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/domain"
)

func (s *Setup) Echo() *echo.Echo {
	return s.echo
}

func (s *Setup) Gin() *gin.Engine {
	return s.gin
}

func (s *Setup) NATS() *nats.Conn {
	return s.nc
}

func (s *Setup) GORM() *gorm.DB {
	return s.db
}

func (s *Setup) Log() domain.Logger {
	return s.log
}

func (s *Setup) Config() *config.Config {
	return s.cfg
}
