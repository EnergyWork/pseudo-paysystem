package api

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type Header struct {
	Error *errs.Error
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
