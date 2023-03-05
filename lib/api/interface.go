package api

import (
	"time"

	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type Request interface {
	Timeout() time.Duration
	Validate() *errs.Error
	Authorize() *errs.Error
}

type Reply interface {
}
