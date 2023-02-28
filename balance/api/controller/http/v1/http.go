// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/httpserver/middleware"
	"github.com/energywork/pseudo-paysystem/lib/setup"

	mw "github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	uc  usecase.Balance
	set *setup.Setup
}

func New(set *setup.Setup, uc usecase.Balance) *Controller {
	return &Controller{
		set: set,
		uc:  uc,
	}
}

// ConfigureRouter ...
func (c *Controller) ConfigureRouter() *Controller {
	e := c.set.Echo()
	// some middleware
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		Skipper:      mw.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete}},
	))
	e.Use(middleware.RequestLogging) // all routes on the server
	// routers
	g := e.Group("/v1")
	newBalanceRoutes(g, c.uc, c.set) // route installer
	return c
}

// SaveRoutes ...
func (c *Controller) SaveRoutes() *errs.Error {
	data, err := json.MarshalIndent(c.set.Echo().Routes(), "", "  ")
	if err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}
	if err := os.WriteFile("routes.json", data, 0644); err != nil {
		return errs.ErrInternal.SetMsg(err.Error())
	}
	return nil
}
