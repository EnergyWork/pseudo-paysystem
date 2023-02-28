package setup

import (
	"testing"

	"github.com/energywork/pseudo-paysystem/lib/config"
)

func TestSetupInit(t *testing.T) {
	cfg := config.New("setup_test")
	if cfg == nil {
		t.Fatal("Config isn't initialized")
	}
	_ = GetSetupForTest(t, cfg)
}
