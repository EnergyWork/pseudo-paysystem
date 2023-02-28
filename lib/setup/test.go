package setup

import (
	"testing"

	"github.com/energywork/pseudo-paysystem/lib/config"
	"github.com/energywork/pseudo-paysystem/lib/logger"
)

func GetSetupForTest(t *testing.T, cfg *config.Config, dbneed ...bool) *Setup {
	set := New(logger.New(logger.LoadLogger(true)), cfg)
	if err := set.NewEcho(); err != nil {
		t.Fatal(err)
	}
	if dbneed[0] {
		if err := set.ConnectPostgres(); err != nil {
			t.Fatal(err)
		}
	}
	if err := set.ConnectNATS(); err != nil {
		t.Fatal(err)
	}
	return set
}
