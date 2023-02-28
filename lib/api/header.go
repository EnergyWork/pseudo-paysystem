package api

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type Header struct {
	//swagger:ignore
	set *setup.Setup
}

func (h *Header) Setup() *setup.Setup {
	return h.set
}

func (h *Header) Timeout() time.Duration {
	return time.Second * 5
}

func (h *Header) Validate() *errs.Error {
	return nil
}

func (h *Header) Authorize() *errs.Error {
	return nil
}

func (h *Header) SetHeader() {

}

func (h *Header) GetHeader() {

}
