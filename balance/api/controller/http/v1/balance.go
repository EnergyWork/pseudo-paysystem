package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/energywork/pseudo-paysystem/balance/api/usecase"
	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type balanceRoutes struct {
	uc  usecase.Balance
	set *setup.Setup
}

func newBalanceRoutes(g *echo.Group, uc usecase.Balance, set *setup.Setup) {
	r := &balanceRoutes{uc: uc, set: set}
	h := g.Group("/balance")
	h.POST("/create", r.create)
	h.POST("/hold", r.hold)
}

func (r *balanceRoutes) create(c echo.Context) error {

	var req *usecase.ReqBalanceCreate
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errs.New().SetCode(errs.Syntax).SetMsg("bad request"))
	}

	rpl, err := r.uc.CreateBalance(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rpl)
}

func (r *balanceRoutes) hold(c echo.Context) error {

	var req *usecase.ReqBalanceHold
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errs.New().SetCode(errs.Syntax).SetMsg("bad request"))
	}

	rpl, err := r.uc.HoldBalance(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rpl)
}
