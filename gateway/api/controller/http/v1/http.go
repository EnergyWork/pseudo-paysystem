// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	"github.com/energywork/pseudo-paysystem/gateway/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/httpserver/middleware"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

// ---------------------------------------------------------------------------------------------------------------------
// Controller

type Controller struct {
	uc  usecase.API
	set *setup.Setup
}

func New(set *setup.Setup, uc usecase.API) *Controller {
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
		AllowMethods: []string{http.MethodPost}},
	))
	e.Use(middleware.RequestLogging) // all routes on the server
	// routers
	g := e.Group("/v1")
	newRoutes(g, c.uc) // route installer
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

// ---------------------------------------------------------------------------------------------------------------------
// Router

type router struct {
	uc usecase.API
}

func newRoutes(g *echo.Group, uc usecase.API) {
	// g-> host:port/v1
	r := &router{uc: uc}

	// define routes
	h := g.Group("/wallet")     // http://host:port/v1/wallet
	h.POST("/create", r.create) // http://host:port/v1/wallet/create
	h.POST("/update", r.update) // http://host:port/v1/wallet/update
	h.POST("/get", r.get)       // http://host:port/v1/wallet/get
}

func (r *router) create(c echo.Context) error {
	req := &usecase.ReqGateWalletCreate{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, `{"error": {"code":"ERROR_REQUEST_SYNTAX", "msg":"bad request"} }`)
	}
	rpl, errApi := r.uc.WalletCreate(req)
	if errApi != nil {
		return c.JSON(http.StatusInternalServerError, errApi)
	}
	log.Printf("Reply: %+v", rpl)
	return c.JSON(http.StatusOK, rpl)
}

func (r *router) update(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}

func (r *router) get(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}

/*
curl -X POST http://localhost:9002/v1/wallet/create -H 'Content-Type: application/json'  -d '{ "phone": "89123731126" }'
*/
