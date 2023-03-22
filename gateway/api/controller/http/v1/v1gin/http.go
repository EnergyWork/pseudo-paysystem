// Package v1gin implements routing paths. Each services in own file.
package v1gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/energywork/pseudo-paysystem/gateway/api/usecase"
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
	g := c.set.Gin()
	gr := g.Group("/v2")
	newRoutes(gr, c.uc) // route installer
	return c
}

// ---------------------------------------------------------------------------------------------------------------------
// Router

type router struct {
	uc usecase.API
}

func newRoutes(g *gin.RouterGroup, uc usecase.API) {
	// g-> host:port/v1
	r := &router{uc: uc}

	// define routes
	h := g.Group("/wallet")     // http://host:port/v1/wallet
	h.POST("/create", r.create) // http://host:port/v1/wallet/create
	// h.POST("/update", r.update) // http://host:port/v1/wallet/update
	// h.POST("/get", r.get)       // http://host:port/v1/wallet/get
}

func (r *router) create(c *gin.Context) {
	req := &usecase.ReqGateWalletCreate{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, `{"error": {"code":"ERROR_REQUEST_SYNTAX", "msg":"bad request"} }`)
		return
	}
	rpl, errApi := r.uc.WalletCreate(req)
	if errApi != nil {
		c.JSON(http.StatusInternalServerError, errApi)
		return
	}
	log.Printf("Reply: %+v", rpl)
	c.JSON(http.StatusOK, rpl)
}

/*func (r *router) update(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}

func (r *router) get(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}*/

/*
curl -X POST http://localhost:9002/v1/wallet/create -H 'Content-Type: application/json'  -d '{ "phone": "89123731126" }'
*/
