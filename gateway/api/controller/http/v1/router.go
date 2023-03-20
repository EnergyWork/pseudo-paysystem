package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/energywork/pseudo-paysystem/gateway/api/usecase"
)

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
		return c.JSON(http.StatusInternalServerError, rpl)
	}
	return c.JSON(http.StatusOK, rpl)
}

func (r *router) update(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}

func (r *router) get(c echo.Context) error {
	return c.JSON(http.StatusOK, "rpl")
}

/*var req *usecase.ReqBalanceCreate
if errApi := c.Bind(&req); errApi != nil {
	return c.JSON(http.StatusBadRequest, errs.New().SetCode(errs.Syntax).SetMsg("bad request"))
}

rpl, errApi := r.uc.CreateBalance(req)
if errApi != nil {
	return c.JSON(http.StatusInternalServerError, errApi)
}*/
