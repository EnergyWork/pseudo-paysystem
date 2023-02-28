package api

import (
	"github.com/energywork/pseudo-paysystem/lib/setup"
)

type Service interface {
	Setup() *setup.Setup
}

type Request interface {
	Setup() *setup.Setup
}

type Reply interface {
}
